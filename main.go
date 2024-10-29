package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parser "goblin.org/main/Parser"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Go REPL!")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		// Remove leading/trailing whitespace and check for exit command.
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		// Process the input (replace with your logic).
		program := parser.BuildAst(input)

		fmt.Printf("%v\n", program)
	}
}
