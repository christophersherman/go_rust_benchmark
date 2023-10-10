#[macro_use]
extern crate rocket;

#[get("/")]
fn mem_test() -> &'static str {
    let mut vec = Vec::with_capacity(10000);
    for i in 0..10000 {
        vec.push(i);
    }
    "Done"
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![mem_test])
}

