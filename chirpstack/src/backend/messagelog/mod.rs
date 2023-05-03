use serde::{Serialize};
use chrono::{DateTime, Utc};
use backend::{JoinReqPayload, BasePayload, ULMetaData, JoinAnsPayload,PRStartReqPayload,PRStartAnsPayload,HomeNSReqPayload,HomeNSAnsPayload,XmitDataReqPayload,XmitDataAnsPayload};
// use backend::Answer;
use chirpstack_api::gw;
//{common, gw};
use std::collections::HashMap;
use uuid::Uuid;
use lrwn::EUI64;
use std::error::Error;

#[derive(Debug, Clone, Serialize) ]
struct RxPacket {  // Seea crate::uplink::UplinkFrameSet and api::UplinkFrameLog  for similar but
                   // not exactly the same data-structures
    #[serde(rename="DR")]
    dr: i64,

    #[serde(rename="TXInfo")]
    #[serde(skip_serializing_if = "Option::is_none")]
    tx_info: Option<gw::UplinkTxInfo>,

    #[serde(rename="RXInfoSet")]
    #[serde(skip_serializing_if = "Vec::is_empty")]
    rx_info_set: Vec<gw::UplinkRxInfo>,

    #[serde(rename="GatewayIsPrivate")]  // Hmm. Should this be remaining as-is or split the way
                                         // uplink data does it in v4?
    #[serde(skip_serializing_if = "HashMap::is_empty")]
    gateway_is_private: HashMap<EUI64, bool>,

    #[serde(rename="GatewayServiceProfile")]
    #[serde(skip_serializing_if = "HashMap::is_empty")]
    gateway_service_profile:  HashMap<EUI64, Uuid>,  // This one is probably renamed
                                                           // gateway_tenant_id_map   rather than
                                                           // gateway_service_profile
    #[serde(rename="RoamingMetaData")]
    #[serde(skip_serializing_if = "Option::is_none")]
    roaming_meta_data: Option<RoamingMetaData>,
}


#[derive(Debug, Clone, Serialize)]
pub struct RoamingMetaData {
    #[serde(rename="BasePayload")]
    pub base_payload: BasePayload,
    #[serde(rename="ULMetaData")]
    pub ul_meta_data: ULMetaData,
}



#[derive(Debug, Clone, Serialize)]
struct TxPacket {
    #[serde(rename="PHYPayload")]
    phy_payload: lrwn::PhyPayload,
    #[serde(rename="DownlinkTXInfo")]
    downlink_tx_info:  Option<gw::DownlinkTxInfo>, // TODO: This maybe should be
                                                   // DownlinkTxInfoLegacy
    #[serde(rename= "TimeOnAir")]
    time_on_air: f64,
}

#[derive(Debug, Clone, Serialize)]
#[serde(rename_all = "PascalCase")]
struct FrameStatus {
    result: FrameStatusResult,
    error_desc: String,
}

#[derive(Debug, Clone, Serialize)]
#[serde(rename_all = "UPPERCASE")]
enum Endpoint {
    Gateway,
    Local,
    Roaming,
    JoinServer,
}

#[derive(Debug, Clone, Serialize)]
enum  FrameStatusResult {
    OK,
    NOK,
    WARN,
}


#[derive(Debug, Clone, Serialize)]
#[serde(rename_all = "PascalCase")]
struct LogEntry {
        // CtxID          interface{}
        #[serde(rename="CtxID")]
        ctx_id: Uuid,
        // .... unknown if needed
        // PublishAt      time.Time
        publish_at: DateTime<Utc>,
        // CreatedAt      time.Time
        created_at: DateTime<Utc>,
        // LogSource      Endpoint
        log_source: Endpoint,
        // SourceID       string
        #[serde(rename="SourceID")]
        source_id: String,
        // LogDestination Endpoint
        log_destination: Endpoint,
        //DestinationID  string
        #[serde(rename="DestinationID")]
        destination_id: String,
        // FrameStatus    FrameStatus
        frame_status: FrameStatus,
//        TimeOnAir      float64
        time_on_air: f64,        
//        DevAddr        lorawan.DevAddr
        dev_addr: lrwn::DevAddr,
//        DevEUI         lorawan.EUI64
        #[serde(rename = "DevEUI")]
        dev_eui: EUI64,   // backend uses Vec<u8> here with a hex_encode encoder.
//        KnownDevice    bool
        known_device: bool,
//        RXPacket       *RXPacket                     `json:"RXPacket,omitempty"`
        #[serde(rename="RXPacket")]
        #[serde(skip_serializing_if = "Option::is_none")]
        rx_packet: Option<RxPacket>,
//        TXPacket       []*TXPacket                   `json:"TXPacket,omitempty"`
        #[serde(rename="TXPacket")]
        #[serde(skip_serializing_if = "Vec::is_empty")]
        tx_packet: Vec<TxPacket>,
//        TXAck          *TXPacket                     `json:"TXAck,omitempty"`
        #[serde(rename="TXAck")]
        #[serde(skip_serializing_if = "Option::is_none")]
        tx_ack: Option<TxPacket>,
//        JoinReq        *lwbackend.JoinReqPayload     `json:"JoinReq,omitempty"`
        #[serde(skip_serializing_if = "Option::is_none")]
        join_req: Option<JoinReqPayload>,
//        JoinAns        *lwbackend.JoinAnsPayload     `json:"JoinAns,omitempty"`
        #[serde(skip_serializing_if = "Option::is_none")]
        join_ans: Option<JoinAnsPayload>,
//        PRStartReq     *lwbackend.PRStartReqPayload  `json:"PRStartReq,omitempty"`
        #[serde(rename="PRStartReq")]
        #[serde(skip_serializing_if = "Option::is_none")]
        pr_start_req: Option<PRStartReqPayload>,
//        PRStartAns     *lwbackend.PRStartAnsPayload  `json:"PRStartAns,omitempty"`
        #[serde(rename="PRStartAns")]
        #[serde(skip_serializing_if = "Option::is_none")]
        pr_start_ans: Option<PRStartAnsPayload>,
//        HomeNSReq      *lwbackend.HomeNSReqPayload   `json:"HomeNSReq,omitempty"`
        #[serde(rename="HomeNSReq")]
        #[serde(skip_serializing_if = "Option::is_none")]
        home_ns_req: Option<HomeNSReqPayload>,
//        HomeNSAns      *lwbackend.HomeNSAnsPayload   `json:"HomeNSAns,omitempty"`
        #[serde(rename="HomeNSAns")]
        #[serde(skip_serializing_if = "Option::is_none")]
        home_ns_ans: Option<HomeNSAnsPayload>,
//        XmitDataReq    *lwbackend.XmitDataReqPayload `json:"XmitDataReq,omitempty"`
        #[serde(skip_serializing_if = "Option::is_none")]
        xmit_data_req: Option<XmitDataReqPayload>,
//        XmitDataAns    *lwbackend.XmitDataAnsPayload `json:"XmitDataAns,omitempty"`
        #[serde(skip_serializing_if = "Option::is_none")]
        xmit_data_ans: Option<XmitDataAnsPayload>,
//        Answer         *lwbackend.Answer             `json:"Answer,omitempty"`
//TODO: This one I cant find as a struct.
//      Investigate it, and see what it is expected to contain
//        #[serde(skip_serializing_if = "Option::is_none")]
//        answer: Option<Answer>,
}


#[cfg(test)]
mod test {
    use super::*;
    use lrwn::DevAddr;
    use uuid::uuid;
    use std::str::FromStr;
 
    #[test]
    fn test_json_format() -> Result<(), Box<dyn Error>>  {

        let mac_payload = lrwn::MACPayload {
            fhdr: lrwn::FHDR {
                f_ctrl: lrwn::FCtrl {
                    adr: true,
                    ..Default::default()
                },
                devaddr: DevAddr::from_str("e0040750")?,
                f_cnt: 7486,
                ..Default::default()
            },
            f_port: Some(2),
            frm_payload: None,
        };

        let phy_payload = lrwn::PhyPayload {
                    mhdr: lrwn::MHDR {
                        m_type: lrwn::MType::ConfirmedDataDown,
                        major:lrwn::Major::LoRaWANR1,
                    },
                    mic: Some([0xc7, 0x2f, 0xeb, 0xdf]),
                    payload: lrwn::Payload::MACPayload(mac_payload),
        };
        let downlink_tx_info = chirpstack_api::gw::DownlinkTxInfo {
            frequency: 869525000,
            power: 27,
            context: vec![0x06, 0xbd, 0x34, 0x40],
            modulation: Some(
                    chirpstack_api::gw::Modulation {
                        parameters: Some(chirpstack_api::gw::modulation::Parameters::Lora(
                                            chirpstack_api::gw::LoraModulationInfo {
                                                bandwidth: 125000,
                                                spreading_factor: 12,
                                                code_rate: gw::CodeRate::Cr45.into(),
                                                polarization_inversion: true,
                                                code_rate_legacy: "".to_string(),
                                    })),
                }),
                ..Default::default()
        };

        let orig = LogEntry {
            publish_at: DateTime::parse_from_rfc3339("2023-05-03T11:58:41.21027935+02:00")?.into(),
            dev_addr: DevAddr::from_str("00000000")?,
            ctx_id: uuid!("c2864e43-174a-42a9-a8a5-71b0bd87b644"),
            known_device: true,
            tx_packet: vec![
                TxPacket {
                    phy_payload: phy_payload,
                    downlink_tx_info: None,
                    time_on_air: 1.4827519999999998,
                },
            ],
            time_on_air:  0.0,
            created_at: DateTime::parse_from_rfc3339("2023-05-03T11:58:41.204119632+02:00")?.into(),
            dev_eui: EUI64::from_str("0080e1150044bb7a")?,
            source_id: "600002".into(),
            log_destination: Endpoint::Gateway,
            frame_status:  FrameStatus { 
                error_desc: "".to_string(),
                result: FrameStatusResult::NOK,
            },
            log_source: Endpoint::Local,
            destination_id: "647fdafffe00c7bb".into(),
            home_ns_ans: None,
            home_ns_req: None,
            join_ans: None,
            join_req: None,
            pr_start_ans: None,
            pr_start_req: None,
            rx_packet: None,
            tx_ack: None,
            xmit_data_ans: None,
            xmit_data_req: None,
        };

        let encoded = serde_json::to_string_pretty(&orig)?;
        println!("{encoded}");

        Ok(())

    }
    const EXPECTED_01: &str = r#"{
  "PublishAt": "2023-05-03T11:58:41.21027935+02:00",
  "DevAddr": "00000000",
  "CtxID": "c2864e43-174a-42a9-a8a5-71b0bd87b644",
  "KnownDevice": true,
  "TXPacket": [
    {
      "PHYPayload": {
        "mhdr": {
          "mType": "ConfirmedDataDown",
          "major": "LoRaWANR1"
        },
        "mic": "c72febdf",
        "macPayload": {
          "fhdr": {
            "fCtrl": {
              "ack": false,
              "fPending": false,
              "classB": false,
              "adrAckReq": false,
              "adr": true
            },
            "fOpts": null,
            "devAddr": "e0040750",
            "fCnt": 7486
          },
          "frmPayload": null,
          "fPort": 2
        }
      },
      "DownlinkTXInfo": {
        "frequency": 869525000,
        "power": 27,
        "context": "vQZANA==",
        "ModulationInfo": {
          "LoraModulationInfo": {
            "bandwidth": 125,
            "polarization_inversion": true,
            "code_rate": "4/5",
            "spreading_factor": 12
          }
        },
        "TimingInfo": {
          "ImmediatelyTimingInfo": {}
        }
      },
      "TimeOnAir": 1.4827519999999998
    }
  ],
  "CreatedAt": "2023-05-03T11:58:41.204119632+02:00",
  "DevEUI": "0080e1150044bb7a",
  "SourceID": "600002",
  "LogDestination": "GATEWAY",
  "FrameStatus": {
    "ErrorDesc": "",
    "Result": "NOK"
  },
  "TimeOnAir": 0,
  "LogSource": "LOCAL",
  "DestinationID": "647fdafffe00c7bb"
}"#;
}



