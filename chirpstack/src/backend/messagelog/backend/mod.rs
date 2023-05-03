mod mqtt;

lazy_static! {
    static ref BACKENDS: RwLock<Option<mqtt::MqttBackend>> = RwLock::new(None);
}
