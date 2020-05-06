fn main() {
    println!("cargo:rerun-if-changed=lib.c");
    cc::Build::new().file("lib.c").compile("lib");
}
