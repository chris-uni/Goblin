# Goblin
(**G**o **B**ut **L**ike **IN**terpreted)

A repository for the Goblin Programming Language.

## Standard Libary

### `data`
```
using "data";

// push, pushes a new value into an array (top-down)
// data.push(arr array, val any)
data.push(arr, "Hello");

// put, puts a new key/value pair into a map, inserted at the end
// data.put(m map, key any, value any)
data.put(map, "key", "World!");

// pop, returns the last element of the specified array
// data.pop(a array)
data.pop(arr);

// size, returns the size of the array or map specified
// data.size(a array), data.size(m map)
let count = data.size(arr);
```

### `io`
```
using "io";

// print, a standard printing function.
// io.print(msg string)
io.print("Hello, World");

// println - acts the same as print, but appends a new line to the end.
// io.println(msg string)
io.println("Hello, World");

// printf - allows for formatted statements to be printed.
// io.printf(formattedString string, args ...any)
io.printf("Hello, %v", "World");

// sprintf - allows for formatted statements to be printed.
// io.sprintf(formattedString string, args ...any) string
let x = io.sprintf("Hello, %v", "World");

// input - reads a single line from std::in.
// io.input(message string) string
let userInput = os.input("Input: ")

// open - returns a new file object using the specified mode, i.e. r, w, a.
// io.open(fileName string, mode string) fileObject
let f = io.open("path/to/file", "r")

// close - closes the specified file object.
// io.close(fileObject *fileObj)
io.close(f)

// readline - reads a single line from the specified file.
// io.readline(fileObject *fileObj, lineNumber int) string
let line = io.readline(f, 1)

// readlines - reads a file line by line
// io.readlines(fileObject *fileObj) []string
let lines = io.readlines(f)

// writen - writes the contents of the buffer to the specified fileObject.
// io.write(fileObject *fileObj, buffer []byte)
io.write(f, b"information")
```
## Language Design

### Variable decleration
```
let x = 10;
const y = 100;
```
### Array decleration & indexing
```
let arr = [1, 2, 3, 4, 5];
let var = arr[2];
println(var);
```

### Map decleration & indexing
```
let x = {
    "foo": 10,
    20: 30,
};
let var = x["foo"];
println(var);
```

### Conditionals
#### if
```
if (2 > 1){
    println("2 is bigger than 1");
}
```
#### if/else
```
if (1 > 2){
    println("1 isnt bigger than 2");
}
else {
    println("1 is smaller than 2");
}
```

### Loops
#### while
```
let i = 0;

while (i < 10) {
    println(i);
    i++;
}
```

#### for
```
let arr = ["foo", "bar", "foobar"];
let map = {
    "foo": "10",
    "bar": 20,
    "foobar": true,
};

for(let i = 0; i < 3; i++;){
    let key = arr[i];
    let val = map[key];
    println(val);
}
```

### Function Decleration & calling
```
fn testPrint(){
    println("Hello, World");
}
testPrint();
```

### Supported Operators
```
x += 1;
x -= 1;
x /= 1;
x *= 1;
x++;
x--;
```
