use actix_web::{get, App, HttpResponse, HttpServer, Responder};
use tokio::fs;

#[get("/")]
async fn greet() -> impl Responder {
    HttpResponse::Ok()
        .content_type("application/json")
        .body("Hello this is code from Rust!")
}

#[get("/health/readiness")]
async fn readiness() -> impl Responder {
    match fs::metadata("/tmp/ready").await {
        Ok(_) => HttpResponse::Ok()
            .content_type("application/json")
            .body("200 OK, it's ready!"),
        Err(_) => HttpResponse::ServiceUnavailable()
            .content_type("application/json")
            .body("503 Service unavailable, it's not ready!")
    }
}

#[get("/health/liveness")]
async fn liveness() -> impl Responder {
    match fs::metadata("/tmp/ready").await {
        Ok(_) => HttpResponse::Ok()
            .content_type("application/json")
            .body("200 OK, it lives!"),
        Err(_) => HttpResponse::ServiceUnavailable()
            .content_type("application/json")
            .body("503 service unavailable, it doesn't live!")
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(greet)
            .service(readiness)
            .service(liveness)
    })
    .bind(("0.0.0.0", 80))?
    .run()
    .await
}
