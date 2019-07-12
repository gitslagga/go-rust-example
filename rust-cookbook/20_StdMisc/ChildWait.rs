use std::process::Command;

fn main() {
    let mut child = Command::new("sleep").arg("5").spawn().unwrap();
    let _result = child.wait().unwrap();

    println!("reached end of main");
}


// $ rustc ChildWait.rs && ./ChildWait
// # `wait` keeps running for 5 seconds until the `sleep 5` command finishes
// reached end of main
