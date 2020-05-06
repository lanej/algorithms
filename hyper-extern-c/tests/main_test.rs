use hyper::{body, Body, Method, Request, Response};

#[test]
fn test_verify_fn() {
    assert_eq!(app::verify("foo").message, "C Programming");
}

#[test]
fn test_lib_struct() {
    assert_eq!(
        app::lib_struct(),
        app::LibResult {
            message: "foobar".to_string()
        }
    );
}

#[test]
fn test_lib_struct_with_input() {
    assert_eq!(
        app::lib_struct_with_input("foobar"),
        app::LibResult {
            message: "foobar".to_string()
        }
    );
}

#[tokio::test(core_threads = 1)]
#[ignore]
async fn test_fallthrough() {
    let request = Request::builder()
        .method(&Method::GET)
        .body(Body::from(""))
        .unwrap();

    let response: Response<Body> = app::serve(request).await.unwrap();
    match body::to_bytes(response.into_body()).await {
        Ok(bytes) => assert_eq!(bytes, ""),
        Err(err) => panic!(err),
    }
}

#[tokio::test(core_threads = 1)]
#[ignore]
async fn test_verify() {
    let request = Request::builder()
        .method(&Method::GET)
        .uri("/verify")
        .body(Body::from(""))
        .unwrap();

    let response: Response<Body> = app::serve(request).await.unwrap();
    match body::to_bytes(response.into_body()).await {
        Ok(bytes) => assert_eq!(bytes, "4"),
        Err(err) => panic!(err),
    }
}
