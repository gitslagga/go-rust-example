// Assume that crate is called adder, will have to extern it in integration test.
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

// extern crate we're testing, same as any other code would do.
extern crate adder;

#[test]
fn test_add() {
    assert_eq!(adder::add(3, 2), 5);
}



// pub fn setup() {
//     // some setup code, like creating required files/directories, starting
//     // servers, etc.
// }

// // extern crate we're testing, same as any other code will do.
// extern crate adder;

// // importing common module.
// mod common;

// #[test]
// fn test_add() {
//     // using common code.
//     common::setup();
//     assert_eq!(adder::add(3, 2), 5);
// }

// $ cargo test