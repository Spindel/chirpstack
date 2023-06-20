use tracing::{error, info, trace};

use anyhow::Result;

mod backend;
mod datatypes;

use crate::config;
pub use datatypes::{
    Endpoint, FrameStatus, FrameStatusResult, LogEntry, LogEntryBuilder, RoamingMetaData, RxPacket,
    TxPacket,
};

use self::backend::mqtt::MqttBackend;
use chirpstack_api::gw;
use tokio::sync::RwLock;

lazy_static! {
    static ref BACKEND: RwLock<Option<MqttBackend>> = RwLock::new(None);
}

pub async fn setup() -> Result<()> {
    let conf = config::get();
    if conf.message_logger.mqtt.servers.is_empty() {
        info!("Message logger disabled.");
    } else {
        let mqtt_backend = MqttBackend::new(&conf.message_logger.mqtt).await?;
        {
            let mut backend = BACKEND.write().await;
            *backend = Some(mqtt_backend);
        }
    }
    Ok(())
}

// Send the log entry, always succeeds
pub async fn send(msg: LogEntry) {
    let guard = BACKEND.read().await;
    if let Some(backend) = &*guard {
        if let Err(e) = backend.log_message(msg).await {
            error!(error = %e, "Messagelog failed to publish");
        }
    } else {
        trace!(msg = ?msg, "Messagelog not configured");
    }
}

fn packet_to_time_on_air(phy_payload: lrwn::PhyPayload, info: gw::LoraModulationInfo) -> f64 {
    match phy_payload.to_vec() {
        Ok(phy) => time_on_air(&phy, info.bandwidth, info.spreading_factor),
        Err(e) => {
            error!("Phy payload marshal error: {}", e);
            0.0
        }
    }
}

fn time_on_air(phy: &[u8], bw: u32, sf: u32) -> f64 {
    let msg_len = phy.len() as f64;
    let de = if (bw == 125) && sf >= 11 { 1.0 } else { 0.0 };

    let t_sym = 2.0_f64.powf(sf as f64) / 1000.0 * bw as f64;

    let t_preamble = 12.5 * t_sym;
    let sf = sf as f64;
    let payload_sym_nb: f64 = (8.0 * msg_len - 4.0 * sf + 44.0) / (4.0 * sf - 2.0 * de);
    let payload_sym_nb = 8.0 + (payload_sym_nb.ceil() * 5.0).max(0.0);
    let t_payload = payload_sym_nb * t_sym;
    t_preamble + t_payload
}
