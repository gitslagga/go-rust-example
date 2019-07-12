use std::env;

fn increase(number: i32) {
    println!("{}", number + 1);
}

fn decrease(number: i32) {
    println!("{}", number - 1);
}

fn help() {
    println!("usage:
ArgumentParsing <string>
    Check whether given string is the answer.
ArgumentParsing {{increase|decrease}} <integer>
    Increase or decrease given integer by one.");
}

fn main() {
    let args: Vec<String> = env::args().collect();

    match args.len() {
        // no arguments passed
        1 => {
            println!("My name is 'ArgumentParsing'. Try passing some arguments!");
        },
        // one argument passed
        2 => {
            match args[1].parse() {
                Ok(42) => println!("This is the answer!"),
                _ => println!("This is not the answer."),
            }
        },
        // one command and one argument passed
        3 => {
            let cmd = &args[1];
            let num = &args[2];
            // parse the number
            let number: i32 = match num.parse() {
                Ok(n) => {
                    n
                },
                Err(_) => {
                    eprintln!("error: second argument not an integer");
                    help();
                    return;
                },
            };
            // parse the command
            match &cmd[..] {
                "increase" => increase(number),
                "decrease" => decrease(number),
                _ => {
                    eprintln!("error: invalid command");
                    help();
                },
            }
        },
        // all the other cases
        _ => {
            // show a help message
            help();
        }
    }
}


// $ rustc ArgumentParsing.rs && ./ArgumentParsing Rust
// This is not the answer.
// $ ./ArgumentParsing 42
// This is the answer!
// $ ./ArgumentParsing do something
// error: second argument not an integer
// usage:
// ArgumentParsing <string>
//     Check whether given string is the answer.
// ArgumentParsing {increase|decrease} <integer>
//     Increase or decrease given integer by one.
// $ ./ArgumentParsing do 42
// error: invalid command
// usage:
// ArgumentParsing <string>
//     Check whether given string is the answer.
// ArgumentParsing {increase|decrease} <integer>
//     Increase or decrease given integer by one.
// $ ./ArgumentParsing increase 42
// 43
