use std::collections::hash_map::DefaultHasher;
use std::collections::HashMap;
use std::env::temp_dir;
use std::hash::Hasher;
use std::io::Cursor;
use std::sync::RwLock;
use std::time::Duration;

use anyhow::{Context, Result};
use async_trait::async_trait;
use futures::stream::StreamExt;
use handlebars::Handlebars;
use paho_mqtt as mqtt;
use prometheus_client::encoding::EncodeLabelSet;
use prometheus_client::metrics::counter::Counter;
use prometheus_client::metrics::family::Family;
use prost::Message;
use rand::Rng;
use serde::Serialize;
use tokio::sync::mpsc;
use tokio::task;
use tracing::{error, info, trace};

use super::GatewayBackend;
use crate::config::GatewayBackendMqtt;
use crate::monitoring::prometheus;
use crate::storage::{get_redis_conn, redis_key};
use crate::{downlink, uplink};
use lrwn::region::CommonName;

pub struct MqttBackend<'a> {
    client: mqtt::AsyncClient,
    topic: String,
    qos: usize,
}


impl<'a> MqttBackend<'a> {
    pub async fn new(conf: &MessageLoggerBackendMqtt) -> Result<MqttBackend<'a>> {
        let topic = conf.log_topic.clone();
        // get client id, this will generate a random client_id when no client_id has been
        // configured.
        let client_id = if conf.client_id.is_empty() {
            let mut rnd = rand::thread_rng();
            let client_id: u64 = rnd.gen();
            format!("{:x}", client_id)
        } else {
            conf.client_id.clone()
        };

        // create client
        let create_opts = mqtt::CreateOptionsBuilder::new()
            .client_id(&client_id)
            .finalize();
        let mut client = mqtt::AsyncClient::new(create_opts).context("Create MQTT client")?;

        client.set_connected_callback(|_client| {
            info!("MQTT connection to messagelog backend.");
        });
        client.set_connection_lost_callback(|_client| {
            error!("MQTT connection to messagelog backend lost");
        });

        // connection options
        let mut conn_opts_b = mqtt::ConnectOptionsBuilder::new();
        conn_opts_b.server_uris(&conf.servers);
        conn_opts_b.automatic_reconnect(Duration::from_secs(1), Duration::from_secs(30));
        conn_opts_b.clean_session(conf.clean_session);
        conn_opts_b.keep_alive_interval(conf.keep_alive_interval);
        conn_opts_b.user_name(&conf.username);
        conn_opts_b.password(&conf.password);
        if !conf.ca_cert.is_empty() || !conf.tls_cert.is_empty() || !conf.tls_key.is_empty() {
            info!(
                ca_cert = conf.ca_cert.as_str(),
                tls_cert = conf.tls_cert.as_str(),
                tls_key = conf.tls_key.as_str(),
                "Configuring connection with TLS certificate"
            );

            let mut ssl_opts_b = mqtt::SslOptionsBuilder::new();

            if !conf.ca_cert.is_empty() {
                ssl_opts_b
                    .trust_store(&conf.ca_cert)
                    .context("Failed to set gateway ca_cert")?;
            }

            if !conf.tls_cert.is_empty() {
                ssl_opts_b
                    .key_store(&conf.tls_cert)
                    .context("Failed to set gateway tls_cert")?;
            }

            if !conf.tls_key.is_empty() {
                ssl_opts_b
                    .private_key(&conf.tls_key)
                    .context("Failed to set gateway tls_key")?;
            }
            conn_opts_b.ssl_options(ssl_opts_b.finalize());
        }
        let conn_opts = conn_opts_b.finalize();

        // get message stream
        let mut stream = client.get_stream(25);

        let b = MqttBackend {
            client,
            topic,
            qos: conf.qos,
        };

        // connect
        info!(clean_session = conf.clean_session, client_id = %client_id, "Connecting to MQTT broker");
        b.client
            .connect(conn_opts)
            .await
            .context("Connect to MQTT broker")?;

        // return backend
        Ok(b)
    }
}

impl MessageLoggerBackend for MqttBackend<'_> {
    async fn log_message(&self, log_entry: messagelog::LogEntry) -> Result<()> {
        todo!();
        let payload = serde_json::to_vec(&log_entry)?;
        info!(topic = %self.topic, "Sending log mesage");
        let msg = mqtt::Message::new(topic, payload, self.qos as i32);
        self.client.publish(msg).await?;
        trace!("Message sent");
        Ok(())
    }
}

