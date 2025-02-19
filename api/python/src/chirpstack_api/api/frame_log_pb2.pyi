from google.protobuf import timestamp_pb2 as _timestamp_pb2
from chirpstack_api.common import common_pb2 as _common_pb2
from chirpstack_api.gw import gw_pb2 as _gw_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DownlinkFrameLog(_message.Message):
    __slots__ = ["dev_addr", "dev_eui", "downlink_id", "gateway_id", "m_type", "phy_payload", "plaintext_f_opts", "plaintext_frm_payload", "time", "tx_info"]
    DEV_ADDR_FIELD_NUMBER: _ClassVar[int]
    DEV_EUI_FIELD_NUMBER: _ClassVar[int]
    DOWNLINK_ID_FIELD_NUMBER: _ClassVar[int]
    GATEWAY_ID_FIELD_NUMBER: _ClassVar[int]
    M_TYPE_FIELD_NUMBER: _ClassVar[int]
    PHY_PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PLAINTEXT_FRM_PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PLAINTEXT_F_OPTS_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    TX_INFO_FIELD_NUMBER: _ClassVar[int]
    dev_addr: str
    dev_eui: str
    downlink_id: int
    gateway_id: str
    m_type: _common_pb2.MType
    phy_payload: bytes
    plaintext_f_opts: bool
    plaintext_frm_payload: bool
    time: _timestamp_pb2.Timestamp
    tx_info: _gw_pb2.DownlinkTxInfo
    def __init__(self, time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., phy_payload: _Optional[bytes] = ..., tx_info: _Optional[_Union[_gw_pb2.DownlinkTxInfo, _Mapping]] = ..., downlink_id: _Optional[int] = ..., gateway_id: _Optional[str] = ..., m_type: _Optional[_Union[_common_pb2.MType, str]] = ..., dev_addr: _Optional[str] = ..., dev_eui: _Optional[str] = ..., plaintext_f_opts: bool = ..., plaintext_frm_payload: bool = ...) -> None: ...

class UplinkFrameLog(_message.Message):
    __slots__ = ["dev_addr", "dev_eui", "m_type", "phy_payload", "plaintext_f_opts", "plaintext_frm_payload", "rx_info", "time", "tx_info"]
    DEV_ADDR_FIELD_NUMBER: _ClassVar[int]
    DEV_EUI_FIELD_NUMBER: _ClassVar[int]
    M_TYPE_FIELD_NUMBER: _ClassVar[int]
    PHY_PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PLAINTEXT_FRM_PAYLOAD_FIELD_NUMBER: _ClassVar[int]
    PLAINTEXT_F_OPTS_FIELD_NUMBER: _ClassVar[int]
    RX_INFO_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    TX_INFO_FIELD_NUMBER: _ClassVar[int]
    dev_addr: str
    dev_eui: str
    m_type: _common_pb2.MType
    phy_payload: bytes
    plaintext_f_opts: bool
    plaintext_frm_payload: bool
    rx_info: _containers.RepeatedCompositeFieldContainer[_gw_pb2.UplinkRxInfo]
    time: _timestamp_pb2.Timestamp
    tx_info: _gw_pb2.UplinkTxInfo
    def __init__(self, phy_payload: _Optional[bytes] = ..., tx_info: _Optional[_Union[_gw_pb2.UplinkTxInfo, _Mapping]] = ..., rx_info: _Optional[_Iterable[_Union[_gw_pb2.UplinkRxInfo, _Mapping]]] = ..., m_type: _Optional[_Union[_common_pb2.MType, str]] = ..., dev_addr: _Optional[str] = ..., dev_eui: _Optional[str] = ..., time: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., plaintext_f_opts: bool = ..., plaintext_frm_payload: bool = ...) -> None: ...
