# Crates & Modules

# Types
## Generics

# Functions
## Function pointers
Function pointers allow you to store functions in variables and pass them as arguments. Examples of usage:

```
// Strategy pattern - choosing different functions at runtime
fn add_one(x: i32) -> i32 { x + 1 }
fn multiply_by_2(x: i32) -> i32 { x * 2 }

fn apply_operation(operation: fn(i32) -> i32, value: i32) -> i32 {
    operation(value)
}

let result1 = apply_operation(add_one, 5);      // 6
let result2 = apply_operation(multiply_by_2, 5); // 10

// Callback functions in C-style APIs
fn callback(x: i32) -> i32 { x * x }
let function_pointer: fn(i32) -> i32 = callback;

// Creating tables of functions
let operations: [fn(i32) -> i32; 2] = [add_one, multiply_by_2];

```
## Closures
Closures are anonymous functions that can capture and use variables from their surrounding environment, making them useful for callbacks and short operations. Closures

- have variables in `||`
- use `{}` for multiline closures
- can use outside/global variables
- once type is inferred, cannot be changed 

```
let outer_var = 2;
let closure = |x: i32| x + outer_var;  // Can capture environment
```

For multiline closures:
```
let outer_var = 2;
    let closure = |x:i32, y: i32| -> i32 {
        x.pow(outer_var) *  y 
    };

    let z = closure(5, 2);
    println!("closure_annotated: {}", z);

```


# Bitwise operators

# Option & Result

# C-like enums

## From and Into, TryFrom and TryInto 

# HashSet & BTreeSet

# BinaryHeap

# VecDeque

# Traits
## Attributes

## Impl

## Format

## From

# Custom Macros

# Threads 

# Testing

# Sync / ASync

# Checked arithmatics and other built in methods

