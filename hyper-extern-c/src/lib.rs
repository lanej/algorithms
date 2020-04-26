use core::convert::Infallible;

use hyper::Body;
use hyper::Method;
use hyper::Request;
use hyper::Response;
use hyper::StatusCode;
use serde_json;
use serde_json::json;

extern "C" {
    fn verifier(x: [char; 255]) -> VerificationResult;
}

#[repr(C)]
pub struct VerificationResult {
    message: [char; 255],
}

pub fn verify(x: String) -> VerificationResult {
    let input: [char; 255] = ['\0'; 255];

    unsafe { verifier(input) }
}

pub async fn serve(req: Request<Body>) -> Result<Response<Body>, Infallible> {
    let action: Result<Response<Body>, RequestError> = match (req.method(), req.uri().path()) {
        (&Method::GET, "/verify") => verify_action(req).await,
        _ => Ok(Response::builder()
            .status(StatusCode::NOT_FOUND)
            .body(Body::from(
                json!({"error": {"message":format!("'{}' not found", req.uri().path())}})
                    .to_string(),
            ))
            .unwrap()),
    };

    match action {
        Ok(response) => Ok(response),
        Err(err) => Ok(Response::builder()
            .status(StatusCode::INTERNAL_SERVER_ERROR)
            .body(Body::from(
                json!({"error": {"message": err.message}}).to_string(),
            ))
            .unwrap()),
    }
}

async fn verify_action(request: Request<Body>) -> Result<Response<Body>, RequestError> {
    // Aggregate the body...
    let whole_body = hyper::body::to_bytes(request).await?;
    // Decode as JSON...
    let json_data: serde_json::Value = serde_json::from_slice(&whole_body).unwrap();
    let unverified_input: String;

    match json_data.get("multiple") {
        Some(data) => unverified_input = data.to_string(),
        None => {
            return Ok(Response::builder()
                .status(StatusCode::BAD_REQUEST)
                .body(Body::from(
                    json!({"error": {"message": "'multiple' parameter missing or blank"}})
                        .to_string(),
                ))
                .unwrap());
        }
    };

    let response_body = format!("Not sure yet: {}", unverified_input);

    Ok(Response::builder()
        .status(StatusCode::OK)
        .body(Body::from(
            json!({"result": { "message": response_body}}).to_string(),
        ))
        .unwrap())
}

struct RequestError {
    message: serde_json::Value,
}

impl<T> std::convert::From<T> for RequestError
where
    T: std::fmt::Display,
{
    fn from(error: T) -> Self {
        RequestError {
            message: json!({ "error": { "message": format!("{}", error) }}),
        }
    }
}
