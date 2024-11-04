# Goblin
(**G**o **B**ut **L**ike **IN**terpreted)

A repository for the Goblin Programming Language.

## Language Features

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
    i = i + 1;
}
```

### Function Decleration & calling
```
fn testPrint(){
    println("Hello, World");
}
testPrint();
```
