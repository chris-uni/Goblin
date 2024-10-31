package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	program "goblin.org/main/program"
	"goblin.org/main/runtime"
)

func main() {

	env := runtime.Environment{
		Stdout:    os.Stdout, // Set the stdout as os.stdout
		Variables: map[string]runtime.RuntimeValue{},
		Constants: map[string]bool{},
	}

	env.Setup()

	reader := bufio.NewReader(os.Stdin)
	args := os.Args

	// Execute file mode.
	if len(args) == 2 {

		// Get the source file.
		executable := args[1]

		if !strings.HasSuffix(executable, ".gob") {
			fmt.Println("Error: File must have a .gob extension!")
			return
		}

		// Open the file
		file, err := os.Open(executable)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		// Read the contents of the file.
		content, err := os.ReadFile(executable)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Run the program.
		result, err := program.Run(string(content), env)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			// Only really want to print to console if its a statement that needs returning.
			if result != nil {
				fmt.Printf("%v\n", result)
			}
		}

	} else {

		// REPL Mode.
		fmt.Println("Goblin v0.1")

		for {

			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			// Remove newline character
			input = strings.TrimSpace(input)

			if input == "exit" {
				fmt.Println("Goodbye!")
				return
			}

			// Run the program.
			result, err := program.Run(input, env)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				// Only really want to print to console if its a statement that needs returning.
				if result != nil {
					fmt.Printf("%v\n", result)
				}
			}
		}
	}
}
