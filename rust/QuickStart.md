# Getting Set Up and Running

## Installation

`$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh`

## Checking version

`$ rustc --version`

## Updating
`rustup update`

## Uninstalling
`$ rustup self uninstall`


# Getting up and Running
## Basic Rust programme
### New Directory and file
```
<folder>
├── main                // Compiled programme
└── main.rs             // Source file
```

### Code
```
fn main() {
    println!("Hello, world!");
}
```

### Compile and run
```
$ rustc main.rs
$ ./main
Hello, world!
```


## Projects using Cargo

## Create new project (using Cargo)
` cargo new <project name>`

## Cargo directory






# Using Cargo
## New project (using Cargo)
` cargo new <new project>`

## Directory
```
<folder>
├── Cargo.lock          // Lock file
├── Cargo.toml          // Dependencies
├── src                 //Source folder
│   └── main.rs
└── target              // Target folder(build folder)
    ├── CACHEDIR.TAG
    └── debug           // Debug build (when using cargo)
        ├── build
        ├── deps
```

# Syntax
```
// Standard library imports
use std::collections::HashMap;  // For hash maps
use std::io;                   // For input/output
use std::vec::Vec;             // Vec is actually prelude, shown for demonstration

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

You can also define a main function that returns a Result:
```
fn main() -> Result<(), Box<dyn Error>> {
    // Handle potential errors
    Ok(())
}
```


#### Function pointers
Function pointers allow you to store functions in variables and pass them as arguments:

```
fn multiply(x: i32, y: i32) -> i32 {
    x * y
}

fn main() {
    // Explicit type annotation
    let f: fn(i32, i32) -> i32 = multiply;
    
    // Type inference
    let f = multiply;
    
    let result = f(12, 12);
    
    // Functions as arguments
    fn apply(f: fn(i32, i32) -> i32, x: i32, y: i32) -> i32 {
        f(x, y)
    }
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
Enum with different types of variants and usage
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

# Console and formatting


# Arithmatics / Numeric operations
## Arithmatics

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


# Loops & Conditionals

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

# Casting, literals, operators



# References, mutability, ownership, copy, shadowing and borrowing

# Crates & Modules

