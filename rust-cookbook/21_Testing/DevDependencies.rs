#[cfg(test)]
#[macro_use]
extern crate pretty_assertions;

pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
    }
}



/* 
Cargo.tom 
[package]
name = "rust-example"
version = "0.1.1"
authors = ["Alice 957766762@qq.com", "Bob Bob@qq.com"]
[dev-dependencies]
pretty_assertions = "0.4.0" 
$ mkdir src && mv DevDependencies.rs src/lib.rs
$ cargo test
*/
