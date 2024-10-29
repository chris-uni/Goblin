package main

import "fmt"

func main() {

	source := "let x = 10;"

	tokens := Tokenize(source)

	fmt.Printf("Tokens: %v \n", tokens)
}
