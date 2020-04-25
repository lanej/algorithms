use hyper::Body;
use hyper::Method;
use hyper::Request;
use hyper::Response;
use hyper::StatusCode;
use std::convert::Infallible;

use hyper::service::{make_service_fn, service_fn};
use hyper::Server;

extern "C" {
    fn doubler(x: u32) -> u32;
}

fn double(x: u32) -> u32 {
    unsafe { doubler(x) }
}

#[tokio::main]
pub async fn main() -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
    // For every connection, we must make a `Service` to handle all
    // incoming HTTP requests on said connection.
    let make_svc = make_service_fn(|_conn| {
        // This is the `Service` that will handle the connection.
        // `service_fn` is a helper to convert a function that
        // returns a Response into a `Service`.
        async { Ok::<_, Infallible>(service_fn(serve)) }
    });

    let addr = ([127, 0, 0, 1], 3000).into();

    let server = Server::bind(&addr).serve(make_svc);

    println!("Listening on http://{}", addr);

    server.await?;

    Ok(())
}

pub async fn serve(req: Request<Body>) -> Result<Response<Body>, Infallible> {
    let response = match (req.method(), req.uri().path()) {
        (&Method::GET, "/verify") => verify(),
        _ => internal_error(),
    };

    Ok(response)
}

fn verify() -> Response<Body> {
    Response::new(Body::from(format!("{}", double(2))))
}

fn internal_error() -> Response<Body> {
    Response::builder()
        .status(StatusCode::INTERNAL_SERVER_ERROR)
        .body(Body::empty())
        .unwrap()
}

mod tests {
    use crate::double;
    use hyper::{body, Body, Method, Request, Response};

    #[test]
    fn test_double() {
        assert_eq!(double(2), 4)
    }

    #[tokio::test(core_threads = 1)]
    async fn test_fallthrough() {
        let request = Request::builder()
            .method(&Method::GET)
            .body(Body::from(""))
            .unwrap();

        let response: Response<Body> = super::serve(request).await.unwrap();
        match body::to_bytes(response.into_body()).await {
            Ok(bytes) => assert_eq!(bytes, ""),
            Err(err) => panic!(err),
        }
    }

    #[tokio::test(core_threads = 1)]
    async fn test_verify() {
        let request = Request::builder()
            .method(&Method::GET)
            .uri("/verify")
            .body(Body::from(""))
            .unwrap();

        let response: Response<Body> = super::serve(request).await.unwrap();
        match body::to_bytes(response.into_body()).await {
            Ok(bytes) => assert_eq!(bytes, "4"),
            Err(err) => panic!(err),
        }
    }
}
