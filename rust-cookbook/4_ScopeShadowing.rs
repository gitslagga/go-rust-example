fn main(){
    let mut long_lived_binding = 1;

    {
        let short_lived_binding = 2;
        println!("inner short {}", short_lived_binding);
        let long_lived_binding = 5_f32;
        println!("inner long {}", long_lived_binding);
    }

    // println!("outer short {}", short_lived_binding);
    println!("outer long {}", long_lived_binding);
    long_lived_binding = 3;
    println!("outer long {}", long_lived_binding);


    
    // Declare a variable binding
    let a_binding;

    {
        let x = 2;

        // Initialize the binding
        a_binding = x * x;
    }

    println!("a binding: {}", a_binding);

    let another_binding;

    // Error! Use of uninitialized binding
    // println!("another binding: {}", another_binding);
    // FIXME ^ Comment out this line

    another_binding = 1;

    println!("another binding: {}", another_binding);
}