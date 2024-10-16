#[macro_use] extern crate rocket;

mod services;
mod models;
use crate::services::sample::*;

#[launch]
fn rocket() -> _ {
    rocket::build()
    .configure(rocket::Config::figment().merge(("port", 8080)))
    .mount("/", routes![get_samples])
    .mount("/", routes![get_sample_by_id])
    .mount("/", routes![post_sample])
    .mount("/", routes![put_sample])
    .mount("/", routes![delete_sample])
}
