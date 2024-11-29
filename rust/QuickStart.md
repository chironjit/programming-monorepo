# Getting Set Up and Running

## Installation
Install Rust via the terminal using the below or via the script on `https://www.rust-lang.org/tools/install`
`$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh`

## Checking version
Check for version installed via `rustc --version`

## Updating
Update Rust via `rustup update`

## Uninstalling
Uninstall Rust using `rustup self uninstall`


## Getting up and Running
Get started with a simple rust programme directly or use Cargo, Rust's built-in package manager

### Basic Rust programme
A basic Rust programme can be created by creating a file with the `.rs` ending, compiling it and running it. The following are the basic steps.

#### New Directory and file
Create a new directory with the file called `main.rs`.
```
<folder>
└── main.rs             // Source file
```

#### Code
Input code in the file.
```
fn main() {
    println!("Hello, world!");
}
```

#### Compile and run
Compile the code using `rustc main.rs` and run the code by calling `./main`.

```
$ rustc main.rs
$ ./main
Hello, world!
```
The folder, after compiling, will have the `main` programme in it.

```
<folder>
├── main                // Compiled programme 
└── main.rs             // Source file
```


### Projects using Cargo
Cargo is Rust's official package manager and build system that handles dependencies, building your code, running tests, and generating documentation. Use Cargo for projects that may be complex and/or have dependecies

#### Create new project (using Cargo)
Create a new project via ` cargo new <project name>`. We'll use `rand-number` - `cargo new rand-number`. You should now have a new folder with the following files.

```
rand-number
├── Cargo.toml
└── src
    └── main.rs
```

#### Add new package and code
We'll add a simple package to illustrate the use of an external package/module. Packages hosted on Cargo are called `crates`

Install the crate called `rand` using `cargo add rand`. Your folder structre will now look like this:

```
├── Cargo.lock
├── Cargo.toml
└── src
    └── main.rs
```

Your original cargo.toml
```
[package]
name = "rand-number"
version = "0.1.0"
edition = "2021"
```
will now have a line for dependencies
```
[package]
name = "rand-number"
version = "0.1.0"
edition = "2021"

[dependencies]
rand = "0.8.5"
```

#### Add code
Add some code to the `main.rs` file:
```
use rand::Rng;

fn main() {
    let random_number = rand::thread_rng().gen_range(1..=100);
    println!("Random number: {}", random_number);
}
```


#### Compile and run
Cargo has 2 build modes, `debug` and `release` modes. The debug mode has additional checks during build that checks for issues such as overflow vulnerabilities, while the build mode has additinal steps that work to optimise the package for performance and binary size.

##### Debug mode 
1. Build in debug mode using `cargo build`
2. Built executable is in `./target/debug` folder
3. Run executable via `./target/debug/rand-number`


##### Release mode:
1. Build in release mode using `cargo build --release`
2. Built executable is in `./target/debug` folder
3. Run executable via `./target/release/rand-number`

#### Cargo run
You can build and run at the same time using `cargo run`

##### Debug mode 
Build and run in debug mode using `cargo run`

##### Release mode
Build and run in release mode using `cargo run --release`

#### Cargo directory
Cargo directories once built can be complex. Here is a simplified version
```
rand-number
├── Cargo.lock
├── Cargo.toml
├── src
│   └── main.rs
└── target
    ├── CACHEDIR.TAG
    ├── debug
    │   ├── build
    │   ├── deps
    │   ├── examples
    │   ├── incremental
    │   ├── rand-number
    │   └── rand-number.d
    └── release
        ├── build
        ├── deps
        ├── examples
        ├── incremental
        ├── rand-number
        └── rand-number.d
```

# Syntax
The code snippet below demonstrates the fundamental syntax and core concepts in Rust. It showcases common language features.

```
/// Standard library imports
// For hash maps
use std::collections::HashMap;
// For input/output
use std::io;
// Vec is actually prelude, shown for demonstration                  
use std::vec::Vec;             

struct Person {
    name: String,
    age: u32,
}

enum Status {
    Active,
    Inactive,
}

fn main() {
    // Variables and types
    let x: i32 = 42;
    let mut y = 10;
    y += 1;

    // Using HashMap from stdlib
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Red"), 50);

    // Basic input example using stdlib
    let mut input = String::new();
    println!("Enter something:");
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line");
    
    // Vector
    let mut vec = vec![1, 2, 3];
    vec.push(4);

    // Control flow
    if x > 40 {
        println!("Greater than 40");
    }

    // Loop
    for i in vec {
        println!("{}", i);
    }

    // Struct usage
    let person = Person {
        name: String::from("Alice"),
        age: 30,
    };

    // Pattern matching
    let status = Status::Active;
    match status {
        Status::Active => println!("Active"),
        Status::Inactive => println!("Inactive"),
    }

    // Using our HashMap
    println!("Score: {:?}", scores.get("Blue"));
}
```

# Comments

```
// This is a single line comment

/* 
This is a multiline comment.
Also usable to comment out specfic parts
as in the example below
*/

fn main() {
    print/*ln*/!("Hello, ");
    println!("world!");
}

```

# Types
## Primitives
Rust primitives are simple built-in types that are implemented directly by the compiler. Rust primitives are:

 - Immutable - immutable by default, unless specified as mutable `mut`
 - Fixed size - size set at compile time, enabling stack allocation
 - Stack allocation - stored directly in stack memory, no heap allocation needed
 - Copy semantics - automatically copied when assigned or passed to functions, rather than being moved like other types
  
### Unsigned integers

The unsigned integers are: `u8`, `u16`, `u32`, `u64`, `u128` and `usize` .

 - Default when not specified is `u32`
 - `usize` sets size to u32 for 32-bit systems and u64 for 64-bit systems
 - you can find min and max values of a size 

```
let a = 1000;

let b: u32 = 1000;

let c = 1000u32; // type at the end

let d: u8 = 1000; // this will error out on compile

// use underscores for readability. Compiler will ignore underscores

let e: u64 = 1_000_000_000_000;

```

You can get the min and max of a given size by usig the `MIN` and `MAX` method, eg:

```
println!("Min u32 is {}", std::u32::MIN);
println!("Max u32 is {}", std::u32::MAX);
```

### Signed integers
The signed integers are: `i8` , `i16` , `i32` , `i64` , `i128` , and `isize`.

 - Default when not specified is `i32`
 - `isize` sets size to i32 for 32-bit systems and i64 for 64-bit systems


```
let a = -1000;

let b: i32 = -1000;

let c = -1000i32;

let d: i8 = -1000; // this will error out on compile

let e: isize = 1__000__000__000__000;
```

You can get the min and max of a given size by usig the `MIN` and `MAX` method, eg:

```
println!("Min i8 is {}", std::i8::MIN);
println!("Max i8 is {}", std::i8::MAX);
```

### Floats
Floats in rust are either `f32` or `f64`. Default assignment is `f64`.

```
let x = 1.618034        // default assignment is f64
let y: f32 = 1.618034   // cast as f32

// decimal required if we want the compiler to infer this as decimal

let z = 10  // inferred as unsingned integer
let z = 10. // inferred as floating point
```


### Characters
#### Chars
Chars are single unicode characters that represent any utf-8 symbol

 - Chars are assigned using single quotes `''` only
 - Can represent basic alphabets, symbols, emojis, etc
 - All chars use 4 bytes of memory but: 
 - When used as part of a string, each char is encoded using the least amount of memory needed for each character.


```
let umlaut = 'ö';
let smiley = '😊';
```

 We use `.len` to get the size of the char in bytes when part of a string:

 ```
println!("Size of a char: {}", std::mem::size_of::<char>());    // 4 bytes
println!("Size of string containing 'a': {}", "a".len());       // 1 byte 
println!("Size of string containing 'ö': {}", "ö".len());       // 2 bytes
println!("Size of string containing '喜': {}", "喜".len());      // 3 bytes
println!("Size of string containing '': {}", "".len());         // 0 byte

```

#### Str (String Slices)
String slices are a fixed-size snippet of a string. In Rust, an `str` is a primitive string type. There is an actual `String` type(see Strings section).

 - Is a reference to a fixed-sized UTF-8 bytes
 - Mutable `str` variables can be reassigned but their contents cannot be modified
 - Can only be accessed via a `&str` reference
 - Is more memory efficient and faster than Strings

```
let x = "Hello"; // defaults to a string slice

let y: & str = "Hello";

// The 'static lifetime means the string literal lives for the entire program duration.
let z: &'static str = "Hello"; 
```



### Boolean
Rust has a built-in boolean type, named bool. It has two values, `true` and `false`

```
let x = true;

let y: bool = false;

```


### Compound types
#### Arrays
Rust arrays are a fixed-size list of elements of the same type. By default, arrays are immutable

```
    let x: [i32; 3] = [1, 2, 3];
    let y = [8, 9, 10];
    let mut z = ['😀', '😄', '😅'];
```

Get the number of elements in an array with the `.len()` method:

```
let x = [1, 2, 3];

println!("x has {} elements", x.len());
```

Access a particular element of an array with subscript notation:

```
let animals = ["Bear", "Panda", "Fox"];

println!("The first animal is: {}", animals[0]);
```

Edit a mutable array:
```
let mut animals = ["Bear", "Panda", "Fox"];
    
animals[0] = "Tiger";
    
println!("The first animal is: {}", animals[0]);
```


#### Tuples
A tuple is an ordered list of fixed size. It can contain element of differing types. You can't change the length or types of a tuple after it's created.

```
let x = ("John", "Doe", 32, true);
let y: (&str, &str, i32, bool) = ("Jane", "Doe", 32, true);
let z: = (0,); // A single-element tuple
```

Access a particular element of a tuple via indexing:

```
let x = ("John", "Doe", 32, true);
println!("The surname is {}", y.1);
```

Edit a mutable tuple:
```
let mut y: (&str, &str, i32, bool) = ("Jane", "Doe", 32, true); 

y.2 = 33;

println!("The new age is {}", y.2);
```

### Functions
#### Syntax
Below is the simplest form of a function, one without arguments and return values. It is written in the shortened form.

- starts with an `fn`
- underscore in function names is the convention
- argumets in brackets(unit `()` type if none)
- specify types in args and returns, e.g `fn a(i32) -> i32`
- arrow `->` when a return type is indicated
- function body in curly braces`{}`
- every expression inside function ends in semicolon`;` (see exceptions below)

```
fn hello_world() {
    println!("Hello, world");
}
```
For functions with parameters and return values, we must specify the types:
```
fn multiply(x: i32, y: i32) -> i32 {
    x * y  // Expression without semicolon for implicit return
}
```

There are two ways to return values:

1. Implicit return (expression without semicolon):
   ```
    fn multiply(x: i32, y: i32) -> i32 {
        x * y
    }
    ```
2. Explicit return statement (with semicolon):
   ```
    fn multiply(x: i32, y: i32) -> i32 {
        return x * y;
    }
   ```
Important note: Adding a semicolon to an expression turns it into a statement and returns () (unit type):
```
fn multiply(x: i32, y: i32) -> i32 {
    x * y;  // ERROR: mismatched types, expected `i32`, found `()`
}
```

#### main()
Every executable Rust program must have a main function, which is the entry point:

```
fn main() {
    // Program starts executing here
}
```

The `main` function can have two valid return types:

1. The unit type (as the above example)
2. Result with an error type `Result<(), E>`
3. 
```
enum MyError {
    InvalidInput,
    FileNotFound,
}

fn main() -> Result<(), MyError> {
    if invalid_input {
        return Err(MyError::InvalidInput);
    }
    Ok(())
}
```


#### Macros
Macros are metaprogramming tools that expand into regular code at compile time. They're denoted by a `!` suffix.

Common built-in macros:
```
// Print with newline
println!("Hello, {}!", "world");

// Print without newline
print!("Loading...");

// Create vectors
let numbers = vec![1, 2, 3, 4, 5];

// Format strings
let message = format!("x = {}, y = {}", 10, 20);

// Assertions for testing
assert!(2 + 2 == 4);
assert_eq!(4, 2 + 2);

// Debug printing
dbg!(numbers);
```

Additional macro features:

- Macros can take a variable number of arguments
- They can generate different code based on the input

```
// Example of variable arguments
println!("{}, {} and {}", "x", "y", "z");

// Repetition in macros
let numbers = vec![1; 5];  // Creates [1, 1, 1, 1, 1]
```

#### Ternary operator
Rust does not have a shortform ternary operator. Instead use `let result = if condition { value1 } else { value2 };`

```
let number = 5;
let message = if number > 0 { "positive" } else { "negative" };
```

### Slice
A `slice` lets you reference a portion of a collection rather than the whole thing. Slices are written as [T], where T is the type of the elements. Slices can be used on `Arrays`, `Vectors`, `Strings` and `str`

```
let array: [i32; 5] = [1, 2, 3, 4, 5];
let slice = &array[1..4];  // Type: &[i32]

let vector: Vec<u32> = vec![1, 2, 3, 4, 5];
let slice = &vector[..];  // Type: &[u32]

// Owned String
let string = String::from("Hello, world!");
let slice = &string[0..5];  // Type: &str

// String literal
let literal = "Hello, world!";
let slice = &literal[7..];  // Type: &str
```



### Other primitive types
Other primitives in Rust:
- unit: The () type, also called “unit”.
- pointer: Raw, unsafe pointers, *const T, and *mut T
- reference: References, &T and &mut T.

Ignore for the quick start guide. To be discussed in later sections

## Advanced types
### Strings
Rust has two main string types: `String` and `&str`. `String` is a growable, heap-allocated data structure, while `&str` is an immutable reference to a UTF-8 encoded string slice.

String Operations
```
// Creating strings
let a = String::from("Hello, world!");

let b = "Hello, world!".to_string();

let mut c = String::new();
c = String::from("Hello, world!");
// or
c = "Hello, world!".to_string();
// or
c.push_str("Hello, world!");

let d: String = "Hello, world!".into(); // Using into() - type annotation needed
```
Getting length in bytes
```
let byte_len = a.len();  // 13
// Note: len() returns the number of bytes, not characters
```
Getting character count
```
let char_count = a.chars().count();  // 13

// For strings with non-ASCII characters:
let unicode = "Hello, 世界!";
println!("Bytes: {}", unicode.len());         // 13 bytes
println!("Chars: {}", unicode.chars().count()); // 9 characters
```
Slicing
```
let hello = &string[0..5];    // "Hello"

// Slicing safety
let s = String::from("Hello, 世界!");

// This will panic - strings must be sliced on character boundaries
// let invalid_slice = &s[0..4];

// Safe way to slice strings
if let Some(slice) = s.get(0..4) {
    println!("Slice: {}", slice);
}
```
Common string operations
```
string.push_str(" More text");  // Append string
string.push('!');              // Append character
string.replace("Hello", "Hi"); // Replace substring
string.to_uppercase();         // Convert to uppercase
string.to_lowercase();         // Convert to lowercase
string.trim();                 // Remove whitespace

// String concatenation
let s1 = String::from("Hello, ");
let s2 = String::from("world!");
let s3 = s1 + &s2;  // Note: s1 is moved here
```

Parsing a string
```
let x: i32 = "32".parse().unwrap();
let y = "15".parse::<i32>().unwrap();
```

### Vec (Vector)
Vectors are growable arrays stored on the heap.

```
// Creating vectors
let mut vec: Vec<i32> = Vec::new();
let vec = vec![1, 2, 3, 4, 5];  // Using vec! macro
let vec = Vec::from([1, 2, 3]); // From array

// Common operations
vec.push(6);             // Add element
vec.pop();              // Remove and return last element
vec.insert(1, 10);      // Insert at index
vec.remove(1);          // Remove at index
vec.clear();            // Remove all elements

// Accessing elements
let third: &i32 = &vec[2];     // Can panic if index out of bounds
let third: Option<&i32> = vec.get(2);  // Safe access

// Iteration
for element in &vec {
    println!("{}", element);
}

// Vector with capacity
let mut vec = Vec::with_capacity(10);
println!("Capacity: {}", vec.capacity());
println!("Length: {}", vec.len());

// Extend vector
let mut vec1 = vec![1, 2, 3];
let vec2 = vec![4, 5, 6];
vec1.extend(vec2);

// Convert to slice
let slice: &[i32] = &vec[..];
```


### Structs
Structs are custom data types that let you group related data together.

```
struct Person {
    name: String,
    age: u32,
    email: String,
}

fn main() {
    let mut person1 = Person {
        name: String::from("Alice"),
        age: 30,
        email: String::from("alice@example.com"),
    };

    person1.age = 31;

    let person2 = Person {
        name: String::from("Bob"),
        ..person1
    };
}

```


### Enums
Enums allow you to define a type that can be one of several variants.

Examples:
```
// Enum with just unit variants
enum Direction {
    Up,
    Down,
    Left,
    Right
}

// Enum with numeric values
enum Status {
    Active = 1,
    Inactive = 0
}
```
Using enums
```
enum Status {
    Active,
    Dormant,
    Paused,
}

let current_state = Status::Active

fn main() {
    match current_state {
        Status::Active => println!("Active"),
        Status::Dormant => println!("Dormant"),
        Status::Paused => println!("Paused"),
    }
}

```
You can also utilise the `use` declaration to shorten the syntax:
```
enum Status {
    Active,
    Dormant,
    Paused,
}

// Import selectively
use crate::Status::{Active, Dormant, Paused};

or 

// Import all
use crate::Status::*

let current_state = Active

fn main() {
    match current_state {
        Active => println!("Active"),
        Dormant => println!("Dormant"),
        Paused => println!("Paused"),
    }
}
```

Enums can have different types of variants and usage
```
enum Message {
    // Unit variant (no data)
    Quit,
    // Named fields variant
    Move { x: i32, y: i32 },
    // Tuple variant (unnamed fields)
    Write(String),
    // Multiple fields tuple variant
    Color(u8, u8, u8)
}

fn main() {
    // Creating enum values
    let quit_message = Message::Quit;
    let move_message = Message::Move { x: 10, y: 20 };
    let write_message = Message::Write(String::from("hello"));
    let color_message = Message::Color(255, 0, 0);

    // Pattern matching with match
    match write_message {
        Message::Quit => println!("Quitting"),
        Message::Move { x, y } => println!("Moving to x: {}, y: {}", x, y),
        Message::Write(text) => println!("Text message: {}", text),
        Message::Color(r, g, b) => println!("Color: rgb({}, {}, {})", r, g, b)
    }

    // Simple if-let pattern matching
    if let Message::Move { x, y } = move_message {
        println!("Moving to position ({}, {})", x, y);
    }
}

```

### Constants

```
const PI: f32 = 3.14159265358979323846;

fn main() {
    let r = 5f32;
    let a: f32 = PI * r * r;
    let c: f32 = 2. * PI * r;
    println!("The area is {a:.2}, and the circumference is {c:.2}");
}
```

# Console and formatting
## Print to console
Use the following macros to print to console(io::stdout):

- `print!("Hello, world!")`: Prints the text in double quotes to the console
- `println!("Hello, world")`: Prints the text in double quotes to the console with a new line appended (in effect the same as `print!("Hello, world!\n")`).


## Formatting strings
Use the `format!` macro to format strings
```
// Basic formatting
let name = "Alice";
let age = 30;
let formatted = format!("Name: {}, Age: {}", name, age);

// Named arguments
let message = format!("{name} is {age} years old");

// Number formatting
let pi = 3.14159;
let formatted_num = format!("{:.2}", pi);  // "3.14"

// Printing in different bases
let b10 = format!("Base 10: {}",   1024);               // 1024
let b2 = format!("Base 2 (binary): {:b}", 1024);        // 10000000000
let b8 = format!("Base 8 (octal): {:o}", 1024);         // 2000
let b16 = format!("Base 16 (hexadecimal): {:x}", 1024); // 400

// Padding and alignment
let padded = format!("{:>10}", "right");  // Right align, width 10
let zeros = format!("{:0>5}", "42");      // "00042"

// Debug formatting
let vector = vec![1, 2, 3];
let debug_str = format!("{:?}", vector);  // "[1, 2, 3]"
```

You can also use these formatting steps in `print` and `println!`

## Printing variables and formatted strings
You can print variables and formatted strings:

```
let first_name = "John";
let last_name = "Doe";
let age:u32 = 30;

println!("This is {} {} and he is {} years old", first_name, last_name, age);

println!("This is {first_name} {last_name} and he is {age} years old");

println!("This is {} {} and he is {} years old", "John", "Doe", 30);
```

## Print error
Use `eprint!("")` and `eprintln!("")` to print out errors to the standard error(io::stderr)

## Getting inputs from console/command line
Get inputs from the console / command line via:

```
use std::io;

fn main() {
   let mut input = String::new();
   
   println!("Enter your name:");
   
   io::stdin()
       .read_line(&mut input)
       .expect("Failed to read line");

   // Remove trailing newline
   let input = input.trim();
   
   println!("Hello, {}", input);
}
```




# Arithmatics / Numeric operations
## Basic Arithmatics

```
let x: i32 = 100;
let y: i32 = 3;

// Addition
let a = x + y;

// Subtraction
let b = y - x;

// Multiplication
let c = x * y;

// Division
let d = x/y;

// Remainder / Modulo
let e = x%y;

```

Additionally, you can do this:
```
let mut x = 5;

x += 3;  // x = x + 3
x -= 2;  // x = x - 2
x *= 4;  // x = x * 4
x /= 2;  // x = x / 2
x %= 3;  // x = x % 3

```


## Other useful arithmatics
### Integers
#### Min & Max Number
Find any min or max of a primitive type of signed, unsigned integer.

```
println!("Min i8 is {}", i8::MIN);
println!("Max i8 is {}", i8::MAX);
```

#### Power
For integers, you can use `pow` to raise the integer by the power

```
let x: i32 = 5

let y: i32 = x.pow(5);
```

#### Absolute
Get the absolute value of a signed integer or float.

```
let x: i32 = -200;

println!("The absolute of x is {}", x.abs())
```

#### Checked arithmatics
You can use checked arithmatics `checked_add`, `checked_sub`, `checked_mul`, `checked_div`, `checked_pow`, `checked_abs`, `checked_rem` to check for overflows of the result

```
let a = 10i32.checked_add(i32::MAX);// None (overflow)
let b = i32::MIN.checked_sub(1);    // None (underflow)
let c = 1000.checked_mul(1000000);  // None (overflow)
let d = 10.checked_div(0);          // None (div by zero)
let e = 2_i32.checked_pow(31);      // None (overflow)
let f = (-128_i8).checked_abs();    // None (overflow)
let g = 5.checked_rem(0);           // None (div by zero)
```

The result in the case of an overflow is the result `None`


### Floats
#### Min & Max
Find any min or max floating point type using `MIN` and `MAX`.

```
println!("Min f32 is {}", f32::MIN);
println!("Max f32 is {}", f32::MAX);
```

Use `min` and `max` methods to get the min and max of 2 numbers

```
let x: f32 = 5.0;
let y: f32 = 6.5;

println!("Min: {}, Max: {}", x.min(y), x.max(y));

```

#### Power & Root
Floating point numbers have 2 unique power functions, `powi`(raised by an integer) and `powf`(raised by a float)

```
let x: f64 = 2.13;
let y = x.powi(2);

let z = 2.13_f64.powf(2.2);
```

You can get square roots for positive floats using the `sqrt` method
```
let x:f64 = 64.0;

println!("The square root is {}", x.sqrt());
```

#### Absolute, Clamp, Recip
Get an absolute value using `abs`
```
let x = -5.0_f64;
let y = x.abs();
```

Use `clamp` to get a max or min value if the variable exceeds the min/max

```
let x = -5.0_f32; 

let y = x.clamp(0.0, 2.0); // y will be 0.0

```

Use `recip` to get the reciprocal(`1/x`)
```
let x = 5.0_f32; 

let y = x.recip(); // Will be 0.2
```


#### Round, Ceil, Floor, Fract, Trunc
Floats have the built in methods to round, get the floor or ceiling and get either the fraction or the integer section of the number.

```
let x: f32 = 3.1415926536;

let r = x.round();
let c = x.ceil();
let f = x.floor();
let fr = x.fract();
let t = x.trunc();

println!("{r}, {c}, {f}, {fr}, {t}");

```

These functions do not support selecting the decimal points to round to.


## Converting between types
Convert between types using `as` 
```
// Integer to Integer
let i: i32 = 100;
let u: u32 = i as u32;    // Signed to unsigned
let l: i64 = i as i64;    // 32-bit to 64-bit

// Integer to Float
let i: i32 = 100;
let f32_num: f32 = i as f32;
let f64_num: f64 = i as f64;

// Float to Integer (truncates decimal part)
let f: f64 = 3.99;
let i: i32 = f as i32;    // i = 3
let u: u32 = f as u32;    // u = 3

// Float to Float
let f64_num: f64 = 3.14159265359;
let f32_num: f32 = f64_num as f32;  // Loses precision

// String to Number
let s = "42";
let i: i32 = s.parse().unwrap();
let f: f64 = s.parse().unwrap();
```


## Scientific notation

Rust supports scientific E-notation, e.g. `2e5`, `-8.8e-9`. The associated type is `f64`.

## Binary, Hexadecimal, Octal, Decimal
Unsigned integers can be represented in binary, hexadecimal and octal using `0b`, `0x` and `0o` respectively.

```
// Binary (base 2)
let binary = 0b1010; // 10 in decimal
let large_binary = 0b1111_0000; // Use _ for readability

// Hexadecimal (base 16)
let hex = 0xFF; // 255 in decimal
let large_hex = 0xDEAD_BEEF;

// Octal (base 8) 
let octal = 0o77; // 63 in decimal
let large_octal = 0o755; // Common file permissions

// Decimal (base 10)
let decimal = 42;
let large_decimal = 1_000_000; // Use _ for readability

// Printing in different formats
println!("Binary: {:b}", binary);
println!("Hex: {:x}", hex);
println!("Octal: {:o}", octal);
println!("Decimal: {}", decimal);
```


## Equality
Rust has the usual equality operators, `==`, `!=`, `>`, `>=`, `<`, `<=`. You can only compare elements of the same type.

```
// Comparison operators
let x = 5;
let y = 10;

assert!(x == y-x); // Equal to
assert!(x != y);   // Not equal to
assert!(x < y);   // Less than
assert!(y > x);   // Greater than
assert!(x <= y);  // Less than or equal
assert!(y >= x);  // Greater than or equal
```

## Boolean logic
The boolean logic of *and* (`&&`), *or* (`||`) and *not* (`!`) are can be used in Rust:

```
// Checking age and membership requirements
let age = 25;
let is_member = true;
if age >= 18 && is_member {
   println!("Access granted to members-only area");
}

// Validating input
let input = "hello";
if input.is_empty() || input.len() > 100 {
   println!("Input must be between 1-100 characters");
}

// Toggling flags
let is_visible = true;
let hidden = !is_visible; // hidden will be false
```

## NaN
Sometimes, calculations requsted will may not be valid. Rust returns `NaN` when that happens:

```
let negative = -4.0_f64;
assert!(negative.sqrt().is_nan());

```

# Loops & Conditionals
## if/else

```
let n = -10;

if n < 0 {
    print!("{} is negative", n);
} else if n > 0 {
    print!("{} is positive", n);
} else {
    print!("{} is zero", n);
}
```

The result can be assigned to a variable:

```
let n = -10;

let sign = if n < 0 {
    "negative"
} else if n > 0 {
    "positive"
} else {
    "zero"
};  // needs semicolon here

```


## loop
`loop` creates an infinite loop unless stopped by `break` or continued via `continue`

### Basic loops

```
let mut counter = 0;

loop {
    counter += 1;
    
    // Skip printing even numbers
    if counter % 2 == 0 {
        continue;
    }
    
    println!("Count: {}", counter);
    
    // Exit loop when counter reaches 5
    if counter >= 5 {
        break;
    }
}

println!("Loop finished!");
```

You can also return values from a loop using break:
```
let result = loop {
    counter += 1;
    if counter == 10 {
        break counter * 2;
    }
};
```

### Nestin & labelling loops
You can label loops to access them in nested loops

```
'outer: loop {
    println!("Outer loop");
    
    'inner: loop {
        println!("Inner loop");
        break 'outer;  // Breaks the outer loop
        // break 'inner;  // Would break just the inner loop
    }
    
    println!("This line will never be reached");
}

println!("Both loops finished!");
```

## while
Use the `while` loop to continue until the condition is met

```
fn main() {
    let mut count = 0;
    
    while count < 5 {
        println!("Count: {}", count);
        count += 1;
    }
    
    println!("Done counting!");
}
```

## for loops
You can use the for loopr to iterate over a range or elements of an array

### for..in a range
Use the `for..in` construct to iterate over a range
```
// this will print the numbers 1 - 9
for n in 1..10 {
        println!("{}", n);
}
    
    
```

Use `a..=b` for a range that is inclusive on both ends
```
// This will print the numbers 1 - 10
for n in 1..=10 {
    println!("{}", n);
}
```

### for..in an iterator
You can iterate over an array or vec using the `.iter` method
```
let animals = vec!["Bear", "Tiger", "Fox", "Lion"];
    
for animal in animals.iter() {
    println!("Animal: {}", animal)
}
```

Note that this method borrows each element ant thn returns it back. For other options, see `into_iter()` and `.iter_mut() ` traits


## match
Rust's match allows for branching based on matches, akin to match and switch statement in ither languages.

Match statements have to be exhaustive(ie, needs to cover every possible option)

```
// Basic match with multiple patterns
let number = 13;
match number {
    1 => println!("One"),
    2 | 3 | 5 | 7 | 11 => println!("Prime"),
    13..=19 => println!("Teen"),
    _ => println!("Other")
}

// Match that returns a value
let boolean = true;
let result = match boolean {
    true => 1,
    false => 0,
};
println!("{}", result);
```

## if let
In cases where matching evevry outcome is not necessary, an `if..let` statement can be used

```
enum Color {
    Red,
    Blue,
    Green,
}

fn main() {
    let color = Color::Red;
    
    if let Color::Red = color {
        println!("It's red!");
    }
}

```

Also can be used to get the matching value:
```
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
}

let msg = Message::Write(String::from("hello"));

if let Message::Write(text) = msg {
    println!("Text message: {}", text);
}
```



## let else
Use `let else` when you need some form of action if the let case fails (the opposite of `if let`)

```
// if-let: handles the success case
if let Some(x) = value {
    // Do something with x when value is Some
}

// let-else: handles the failure case
let Some(x) = value else {
    // Do something when value is None
    return;
};
// Continue with x when value is Some
```

And example in use:
```
fn process_name(input: &str) {
   let name = {
       let Some(val) = input.trim().is_empty().then(|| input.trim()) else {
           println!("Name cannot be empty!");
           return;
       };
       val
   };

   println!("Hello, {}!", name);
}

fn main() {
   process_name("Alice");     // Prints: Hello, Alice!
   process_name("   ");       // Prints: Name cannot be empty!
   process_name("");          // Prints: Name cannot be empty!
}


```

## while let
While let allow you to continue processing values until a specifed pattern stops matching

```
// A vector we'll take items from
let mut stack = vec![1, 2, 3, 4, 5];

// Continue popping items until None is returned
while let Some(value) = stack.pop() {
    println!("Got value: {}", value);
}
```
An example using an iterator:
```
let mut iterator = vec![1, 2, 3].into_iter();

while let Some(num) = iterator.next() {
    println!("Number: {}", num);
}
```


# Destructuring

Destructure elements in a variable:

```
// Structs
struct Point { x: i32, y: i32 }
let p = Point { x: 0, y: 7 };
let Point { x, y } = p;

// Tuples
let tuple = (1, "hello", true);
let (num, text, flag) = tuple;

// Arrays/Slices
let arr = [1, 2, 3];
let [a, b, c] = arr;

// Enums
enum Message {
   Quit,
   Move { x: i32, y: i32 },
   Write(String),
}
let msg = Message::Move { x: 3, y: 4 };
if let Message::Move { x, y } = msg { }

// References
let reference = &(1, 2);
let &(a, b) = reference;

// Pattern matching in function parameters
fn print_coordinates(&(x, y): &(i32, i32)) {
   println!("({}, {})", x, y);
}
```

# Option and result
Rust has built in enums called `Option` and `Result` that allow you to handle specific cases where the expected return value/type is not guaranteed.

## Option
`Option` is an enum that has:
 - `None` - to indicate failure or the lack of a value
 - `Some(value)` - to indicate success/presence of a value

`Some(value)` is a tuple struct that wraps a value with type T(generic type).  

```
fn divide(numerator: u8, denominator: u8) -> Option<u8> {
    // Check for division by zero
    if denominator == 0 {
        return None;
    }
    
    // Check for potential overflow
    if numerator > (u8::MAX / denominator) {
        return None;
    }
    
    Some(numerator / denominator)
}

fn main() {
    println!("250 / 10 = {:?}", divide(250, 10));   // None (would overflow u8)
    println!("100 / 0 = {:?}", divide(100, 0));     // None (division by zero)
    println!("10 / 2 = {:?}", divide(10, 2));       // Some(5)
}
```

## Result
`Result` is Rust's way of handling operations that can either succeed with a value (`Ok`) or fail with an error (`Err`).

```
fn divide(x: i32, y: i32) -> Result<i32, String> {
    if y == 0 {
        Err(String::from("Cannot divide by zero"))
    } else {
        Ok(x / y)
    }
}

fn main() {
    match divide(10, 2) {
        Ok(result) => println!("Result: {}", result),
        Err(e) => println!("Error: {}", e),
    }

    match divide(10, 0) {
        Ok(result) => println!("Result: {}", result),
        Err(e) => println!("Error: {}", e),
    }
}
```



# Mutability, Ownership, References, Borrowing & Shadowing
## Mutability
All variables in Rust is immutable by default. Ti make it mutable, use `mut` 
```
let x = 5;                  // immutable
// x = 6;                   // Error!

let mut y = 5;              // mutable
y = 6;    
```

## Ownership
Every value in Rust can only have one owner. To use that value, you either need to copy it or borrow it.


In primitive types, the Rust compiler automatically copies the value on assignment as primitives implement the `Copy` trait
```
// x and y are now pointing to 2 distinct number in memory
let mut x = 100i32;
let y = x;

println! ("{}, {}", x, y);

x = 200;

println! ("{}, {}", x, y);

```

In other types, this copy is not guaranteed
```
let x = String::from("hello");
let mut y = x;

// This will error as the value from x has been moved to y
// x does not exist anymore
// println! ("{}, {}", x, y); 

println! ("{}", y);

y = "Hello, world!".to_string();

println! ("{}", y);

```

Note that Tuples only implement the Copy trait if they contain types with the Copy trait

## References
References are a way to refer to a value without taking ownership of it. Use the ampersand `&` to borrow a value
```
fn calculate_length(s: &String) -> usize {
    s.len()
}

let x = String::from("hello");

let y = calculate_length(&x);    // & creates a reference

println!("Length of '{}' is {}", x, y);

```

## Borrowing
Borrowing is the act of temporarily borrow a variable using references. You can make mutiple immutable references but only one mutable one

```
let mut s = String::from("hello");

let r1 = &s;                // immutable borrow
let r2 = &s;                // multiple immutable borrows ok
println!("{}, {} {}", s, r1, r2);

let r3 = &mut s;            // mutable borrow
r3.push_str(" world");
println!("{}", r3);
// println!("{}, {}", s, r3) // Will not work
```

## Clone
You can manually clone variables of types that do not implement copy. That said, this can be computationally expensive and there may be better options to utilise the same set of data
```
let x = String::from("hello");
let y = x.clone();

println!("{}, {} world!", x, y);
```

## Shadowing
Shadowing is creating a new variable with the same name as the previous variable
```
fn main() {
    let x = 5;
    let x = x + 1;        // shadows previous x
    let x = x * 2;        // shadows again
    println!("{}", x);    // Prints: 12
    
    let spaces = "   ";
    let spaces = spaces.len();  // Can even change type
}
```

## Ownership and functions
Passing a value to a function affects ownership just like passing to a function

With primitive types, the value is copied when passed to a function
```
let x = 63.4;

fn conversion(y: f64) {
    let height = y * 2.54 / 100.;
    println!("You are {} metres tall", height);
}

conversion(x);

println!("You are {} inches tall", x);
```

With more complex types, this doesn't work
```
let x = String::from("Sammy");

fn greeting(name: String) {
    println!("Hello, {}", name);
}

greeting(x);

// println!("Hello, {}", x); // error[E0382]: borrow of moved value: `x`

```

For strings, we can easily solve this by taking a string slice instead of owned String literal
```
let x = String::from("Sammy");

fn greeting(name: &str) {
    println!("Hello, {}", name);
}

greeting(&x);    // Pass a reference to x
println!("Hello, {}", x); 
```

# Useful Methods
## Slices
### String slices
String slices allow you to take specific slices of a string (first digit inclusive, last digit exclusive)
```
let hello = String::from("Hello, Bob");

// Get the first x characters
let slice = &hello[0..2];
let slice = &hello[..2];

let len = hello.len();

// Get the last x characters
let slice = &hello[8..len];
let slice = &hello[8..];

// Get the whole string
let slice = &hello[0..len];
let slice = &hello[..];
```

### Other slices
Other types, including arrays and vecs can also use slices
```
let alphabet = ['a', 'b', 'c', 'd', 'e'];

let slice = &alphabet[1..3];
```

## Iterators
### Iterators
Use `.iter()`, `.iter_mut()` & `into_iter()` methods:
```
let numbers = vec![1, 2, 3, 4, 5];

// Create an iterator
let iter = numbers.iter();           // immutable references
let iter_mut = numbers.iter_mut();   // mutable references
let into_iter = numbers.into_iter(); // takes ownership

// Common iterator methods
let doubled: Vec<i32> = numbers.iter()
    .map(|x| x * 2)                 // transform elements
    .collect();                     // collect into a Vec

let sum: i32 = numbers.iter()
    .sum();                         // sum all elements

let even: Vec<&i32> = numbers.iter()
    .filter(|x| *x % 2 == 0)        // keep only even numbers
    .collect();
```

### Useful iterator methods
Iterators can be chained with other methods to get useful results
```
let nums = vec![1, 2, 3, 4, 5];

// count() - count elements
let count = nums.iter().count();

// any() - check if any element matches predicate
let has_even = nums.iter().any(|x| x % 2 == 0);

// all() - check if all elements match predicate
let all_positive = nums.iter().all(|x| *x > 0);

// find() - find first element matching predicate
let first_even = nums.iter().find(|x| *x % 2 == 0);

// position() - find index of first matching element
let pos = nums.iter().position(|x| *x == 3);

// max() and min() - find maximum/minimum element
let max = nums.iter().max();
let min = nums.iter().min();

// take(n) - take first n elements
let first_three: Vec<&i32> = nums.iter().take(3).collect();

// skip(n) - skip first n elements
let last_two: Vec<&i32> = nums.iter().skip(3).collect();

// enumerate() - pair elements with their index
for (index, value) in nums.iter().enumerate() {
    println!("Index: {}, Value: {}", index, value);
}
```


