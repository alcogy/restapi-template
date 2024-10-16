use rocket::serde::json::Json;

use crate::models::sample::*;

#[get("/samples")]
pub fn get_samples() -> Json<Vec<Sample>> {
    let vec = select_all_samples();
    Json(vec)
}

#[get("/sample/<id>")]
pub fn get_sample_by_id(id: &str) -> Json<Sample> {
    let value = select_sample_by_id(id);
    Json(value)
}

#[post("/sample", format = "application/json", data = "<sample>")]
pub fn post_sample(sample: Json<Sample>) -> Json<PostResponse> {
    let res = insert_sample(sample.0);
    Json(res)
}

#[put("/sample", format = "application/json", data = "<sample>")]
pub fn put_sample(sample: Json<Sample>) -> Json<Message> {
    let res = update_sample(sample.0);
    Json(res)
}

#[delete("/sample/<id>")]
pub fn delete_sample(id: &str) -> Json<Message> {
    let res = delete_sample_data(id);
    Json(res)
}