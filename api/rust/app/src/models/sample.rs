use serde::{Serialize, Deserialize};
use mysql::*;
use mysql::prelude::*;
use super::database::connection;

#[derive(Serialize, Deserialize, FromRow)]
pub struct Sample {
    pub id: String,
    pub name: String,
    pub value: String,
    pub latest_update: String,
}

#[derive(Serialize, Deserialize)]
pub struct PostResponse {
    pub id: String,
}

#[derive(Serialize, Deserialize)]
pub struct Message {
    pub message: String,
}

// SELECT All
pub fn select_all_samples() -> Vec<Sample> {
    let mut conn = connection().unwrap();
    let selected = conn.query_map(
        "SELECT * FROM my_table;",
        |(id, name, value, latest_update)| {
            Sample { id, name, value, latest_update }
        },
    );

    selected.unwrap()
}

// SELECT ONE
pub fn select_sample_by_id(id: &str) -> Sample {
    let mut conn = connection().unwrap();
    let stmt = conn.prep("SELECT CAST(id AS char) AS id, name, value, DATE_FORMAT(latest_update, '%Y-%m-%dT%H:%i:%s') AS latest_update FROM my_table WHERE id = :id").unwrap();
    let selected = conn.exec_first::<Sample,_,_>(&stmt, params! { "id" => id }).unwrap();
    
    selected.unwrap()
}

// INSERT
pub fn insert_sample(sample: Sample) -> PostResponse {
    let mut conn = connection().unwrap();
    let mut tx = conn.start_transaction(TxOpts::default()).unwrap();
    let _ = tx.exec_drop("INSERT INTO my_table VALUES (0, ?, ?, now())", (sample.name, sample.value,));
    let id: Option<String> = tx.query_first("select last_insert_id();").unwrap();
    let _ = tx.commit();    

    PostResponse { id: id.unwrap() }
}

// UPDATE
pub fn update_sample(sample: Sample) -> Message {
    let mut conn = connection().unwrap();
    let stmt = conn.prep("UPDATE my_table SET latest_update = now(), name = :name, value = :value WHERE id = :id").unwrap();
    conn.exec_drop(&stmt, params! { "name" => sample.name, "value" => sample.value, "id" => sample.id }).unwrap();
    
    Message { message: "ok".to_string() }
}

// DELETE
pub fn delete_sample_data(id: &str) -> Message {
    let mut conn = connection().unwrap();
    let stmt = conn.prep("DELETE FROM my_table WHERE id = :id").unwrap();
    conn.exec_drop(&stmt, params! { "id" => id }).unwrap();
    
    Message { message: "ok".to_string() }
}
