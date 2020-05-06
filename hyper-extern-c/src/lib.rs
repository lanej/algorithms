use core::convert::Infallible;

use hyper::Body;
use hyper::Method;
use hyper::Request;
use hyper::Response;
use hyper::StatusCode;
use serde_json;
use serde_json::json;
use std::ffi::CStr;
use std::ffi::CString;
use std::os::raw::c_char;

#[repr(C)]
#[derive(Debug)]
struct libResult {
    message: *const c_char,
}

#[repr(C)]
#[derive(Debug, PartialEq)]
pub struct LibResult {
    pub message: String,
}

pub fn verify(input: &str) -> LibResult {
    extern "C" {
        fn just_message(input: *const c_char) -> *const c_char;
    }

    let foo = unsafe { just_message(CString::new(input).unwrap().as_ptr()) };

    let message = unsafe {
        CStr::from_ptr(foo)
            .to_str()
            .expect("invalid string from verifier")
    };

    dbg!(&message);

    LibResult {
        message: message.to_owned(),
    }
}

pub fn lib_struct() -> LibResult {
    extern "C" {
        fn just_struct() -> libResult;
    }

    let foo = unsafe { just_struct() };

    let message = unsafe { CStr::from_ptr(foo.message).to_str().unwrap() };

    dbg!(&message);

    LibResult {
        message: message.to_owned(),
    }
}

pub fn lib_struct_with_input(input: &str) -> LibResult {
    extern "C" {
        fn just_struct_with_input(input: *const c_char) -> *mut libResult;
    }

    let lib_result = unsafe { just_struct_with_input(CString::new(input).unwrap().as_ptr()) };

    dbg!(&lib_result);

    let message = unsafe { CStr::from_ptr((*lib_result).message).to_str().unwrap() };

    LibResult {
        message: message.to_owned(),
    }
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
