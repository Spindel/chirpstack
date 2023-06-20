use std::time::Duration;

use anyhow::{Context, Result};
use paho_mqtt as mqtt;
use rand::Rng;
use tracing::{error, info, trace};

use crate::config::MessageLoggerBackendMqtt;

use crate::messagelog;

pub struct MqttBackend {
    client: mqtt::AsyncClient,
    topic: String,
    qos: usize,
}

impl MqttBackend {
    pub async fn new(conf: &MessageLoggerBackendMqtt) -> Result<MqttBackend> {
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
        let client = mqtt::AsyncClient::new(create_opts).context("Create MQTT client")?;

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

    pub async fn log_message(&self, mut log_entry: messagelog::LogEntry) -> Result<()> {
        log_entry.calculate_and_set_time_on_air();
        let payload = serde_json::to_vec(&log_entry)?;
        info!(topic = %self.topic, "Sending log message");
        let msg = mqtt::Message::new(&self.topic, payload, self.qos as i32);
        self.client.publish(msg).await?;
        trace!("Log message sent");
        Ok(())
    }
}

#[cfg(test)]
pub mod test {
    use super::*;
    use crate::test;

    use futures::stream::StreamExt;

    use paho_mqtt as mqtt;

    // Helper to get a listen-able client for the test-case
    async fn make_mqtt_client() -> mqtt::AsyncClient {
        // Set up a mqtt client
        let create_opts = mqtt::CreateOptionsBuilder::new()
            .server_uri("tcp://mosquitto:1883/")
            .finalize();
        let client = mqtt::AsyncClient::new(create_opts).unwrap();
        let conn_opts = mqtt::ConnectOptionsBuilder::new()
            .clean_session(true)
            .finalize();
        client.connect(conn_opts).await.unwrap();

        client
            .subscribe("messagelog/testcase", mqtt::QOS_0)
            .await
            .unwrap();
        client
    }

    #[tokio::test]
    async fn test_messagelog() {
        // We don't use the backend here, as the test-case is only really testing the MQTT
        // connection currently.
        //
        // It should be increased to test/verify our expected events as well.
        let _guard = test::prepare().await;
        let conf = MessageLoggerBackendMqtt {
            log_topic: "messagelog/testcase".into(),
            servers: vec!["tcp://mosquitto:1883/".into()],
            ..Default::default()
        };
        let mqtt_backend = MqttBackend::new(&conf).await.unwrap();
        let mut client = make_mqtt_client().await;
        let mut stream = client.get_stream(10);

        // TODO: Spindel,  add test-case for Roaming
        //      backend.PRStartReq
        //      backend.XmitDataReq
        //      handlePRStartAns
        //      handlePRStartReq
        //      handlePRStartReqData

        // TODO: Spindel, add test-case for HandleDownlinkTXAck (downlink /ack )
        //  forJoinAcceptPayload
        //
        //
        //
        // uplink event

        // TODO: Spindel, add test-case for HandleDownlinkTXAck (downlink /ack )
        //

        let log_entry = messagelog::LogEntry::default();
        let expected = serde_json::to_string(&log_entry).unwrap();
        mqtt_backend.log_message(log_entry).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(msg.payload_str(), expected);

        /*
        let pl = integration::UplinkEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.uplink_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/up",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // join event
        let pl = integration::JoinEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.join_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/join",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // ack event
        let pl = integration::AckEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.ack_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/ack",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // txack event
        let pl = integration::TxAckEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.txack_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/txack",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // log event
        let pl = integration::LogEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.log_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/log",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // status event
        let pl = integration::StatusEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.status_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/status",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // location event
        let pl = integration::LocationEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.location_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/location",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // integration event
        let pl = integration::IntegrationEvent {
            device_info: Some(integration::DeviceInfo {
                application_id: Uuid::nil().to_string(),
                dev_eui: "0102030405060708".to_string(),
                ..Default::default()
            }),
            ..Default::default()
        };
        i.integration_event(&HashMap::new(), &pl).await.unwrap();
        let msg = stream.next().await.unwrap().unwrap();
        assert_eq!(
            "application/00000000-0000-0000-0000-000000000000/device/0102030405060708/event/integration",
            msg.topic()
        );
        assert_eq!(serde_json::to_string(&pl).unwrap(), msg.payload_str());

        // downlink command
        let down_cmd = integration::DownlinkCommand {
            id: Uuid::new_v4().to_string(),
            dev_eui: dev.dev_eui.to_string(),
            confirmed: false,
            f_port: 10,
            data: vec![1, 2, 3],
            object: None,
        };
        let down_cmd_json = serde_json::to_string(&down_cmd).unwrap();
        client
            .publish(mqtt::Message::new(
                format!("application/{}/device/{}/command/down", app.id, dev.dev_eui),
                down_cmd_json,
                mqtt::QOS_0,
            ))
            .await
            .unwrap();

        // give the async consumer some time to process
        sleep(Duration::from_millis(200)).await;

        let queue_items = device_queue::get_for_dev_eui(&dev.dev_eui).await.unwrap();
        assert_eq!(1, queue_items.len());
        assert_eq!(down_cmd.id, queue_items[0].id.to_string());
        assert_eq!(dev.dev_eui, queue_items[0].dev_eui);
        assert_eq!(10, queue_items[0].f_port);
        assert_eq!(vec![1, 2, 3], queue_items[0].data);

        */
    }
}
