use anyhow::Result;

pub mod joinserver;
pub mod keywrap;
pub mod roaming;

pub mod messagelog;

pub fn setup() -> Result<()> {
    joinserver::setup()?;
    roaming::setup()?;

    Ok(())
}
