use core::convert::Infallible;

use hyper::Body;
use hyper::Method;
use hyper::Request;
use hyper::Response;
use hyper::StatusCode;
use libc::{c_char, c_int};
use serde_json;
use serde_json::json;
use std::ffi::CStr;
use std::ffi::CString;

#[repr(C)]
#[derive(Debug)]
struct CResult {
    message: *mut c_char,
    status: c_int,
    code: [c_char; 3],
    enabled: c_char,
}

#[repr(C)]
#[derive(Debug, Default, PartialEq)]
pub struct RustResult {
    pub message: String,
    pub status: i32,
    pub code: String,
    pub enabled: Toggle,
}

#[repr(i8)]
#[derive(Debug, PartialEq)]
pub enum Toggle {
    Unknown = 0,
    On = 'Y' as i8,
    Off = 'N' as i8,
}

impl Default for Toggle {
    fn default() -> Self {
        Toggle::Unknown
    }
}

pub fn verify(input: &str) -> RustResult {
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

    RustResult {
        message: message.to_owned(),
        ..Default::default()
    }
}

pub fn lib_struct() -> RustResult {
    extern "C" {
        fn just_struct() -> CResult;
    }

    let foo = unsafe { just_struct() };

    let message = unsafe { CStr::from_ptr(foo.message).to_str().unwrap() };

    dbg!(&message);

    RustResult {
        message: message.to_owned(),
        ..Default::default()
    }
}

pub fn lib_struct_with_input(input: &str) -> RustResult {
    extern "C" {
        fn just_struct_with_input(input: *const c_char) -> *mut CResult;
    }

    let cresult = unsafe { just_struct_with_input(CString::new(input).unwrap().as_ptr()) };

    dbg!(&cresult);

    let message = unsafe { CStr::from_ptr((*cresult).message).to_str().unwrap() };

    RustResult {
        message: message.to_owned(),
        ..Default::default()
    }
}

fn to_char_array(bytes: Vec<u8>) -> Option<[i8; 3]> {
    let mut slice = [0i8; 3];
    if bytes.len() > slice.len() {
        return None;
    }

    slice.copy_from_slice(
        bytes
            .iter()
            .map(|c| *c as i8)
            .into_iter()
            .collect::<Vec<i8>>()[0..bytes.len()]
            .as_ref(),
    );
    return Some(slice);
}

pub fn lib_mutate_me() -> RustResult {
    extern "C" {
        fn mutate_struct(result: *mut CResult) -> libc::c_int;
    }

    let mut result = CResult {
        message: CString::new("Hello! ").unwrap().into_raw(),
        status: 1,
        code: to_char_array(CString::new("ab").unwrap().into_bytes_with_nul()).unwrap(),
        enabled: 'N' as libc::c_char,
    };

    unsafe { mutate_struct(&mut result) };

    return result.into();
}

impl std::convert::From<CResult> for RustResult {
    fn from(cresult: CResult) -> Self {
        return RustResult {
            message: unsafe {
                CString::from_raw(cresult.message)
                    .to_string_lossy()
                    .to_string()
            },
            status: cresult.status,
            enabled: unsafe { std::mem::transmute(cresult.enabled) },
            code: unsafe {
                CString::from_vec_unchecked(
                    cresult.code.to_vec().into_iter().map(|c| c as u8).collect(),
                )
                .to_string_lossy()
                .to_string()
            },
        };
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
