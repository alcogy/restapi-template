use mysql::*;
use dotenv::dotenv;
use std::env;

pub fn connection() -> Result<PooledConn, Error> {
  dotenv().ok();
  
  let url = env::var("DB_URL").expect("undefined [DB_URL]");
  let pool = Pool::new(url.as_ref()).unwrap();

  pool.get_conn()
}