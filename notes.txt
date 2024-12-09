Rough structure on what the lang will look like:

-- defining variables

let x = 10;
const y = 12;
x = 11;

-- functions

-- Simple function.
fn SayHello () {

}

-- Function with example of param and returning.
fn int AddOne (int a) {
    return a + 1;
}

-- Example of defining a function as a variable
-- now 'FnAsVariable' can be passed around.

fn FnAsVariable () => {
    return 10;
}

-- Single line comment
/* ... */ Multi-line comment

-- Conditionals

if ( ... ) {

}
elif ( ... ){

}
else {

}

-- loops

while () {

}

for (int i = 0; i < 8; i++){
    
}

-- data structures

let x = {
    "foo": "bar"
};

let var = x.foo;

-- file io

Inspiration from pyhton?


with open("file-path.txt", "r") as file {

}

Or this? 

// Dedicated file-io iterator
let file = io.open("file-path.txt", "r");

it line in file {

    // Do stuff with line?
}

// -- Bootstrapping Goblin, notes:

- More debug support, things like:
    - ~~Capturing line numbers so that developers can identify which line a particular error is on~~
    - ~~A standardised error message~~
- A better Lexer, find a way to make use of Go Routines to add concurrent threads
- In a similar vein, improve the parser too, there are sections that definitely could do with some revision.
- In order to Bootstrap Goblin with itself, the language needs some core features:
    - A robust way to read from a file line by line (in fact, more robust file io)
    - ~~Better support for built-in data types (i.e. stdlib for arrays, maps)~~
    - ~~Split strings into individual characters~~
    - ~~Ability to compare characters to one another (i.e. for the lexer)~~
    - Ability to define custom Types (at leeast, classes would be nice, but that involved thinking about OOP principles etc)
    - Ability to compare Types (i.e. for the parser, interpreter)
    - Ability to call/execute `file-b.gob` from within `file-a.gob`.

