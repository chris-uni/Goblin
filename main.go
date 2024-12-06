package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	program "goblin.org/main/program"
	"goblin.org/main/runtime"
	"goblin.org/main/utils"
)

func main() {

	env := runtime.Environment{
		Stdout:     os.Stdout,
		Stdin:      os.Stdin,
		Variables:  map[string]runtime.RuntimeValue{},
		Constants:  map[string]bool{},
		Namespaces: map[string]runtime.Namespace{},
	}

	env.Setup()

	reader := bufio.NewReader(os.Stdin)
	args := os.Args

	// Execute file mode.
	if len(args) == 2 {

		// Get the source file.
		executable := args[1]

		if !strings.HasSuffix(executable, ".gob") {
			utils.Stdout("Error: File must have a .gob extension!", env.Stdout)
			return
		}

		// Open the file
		file, err := os.Open(executable)
		if err != nil {
			e := fmt.Sprintf("Error: %v", err)
			utils.Stdout(e, env.Stdout)
			return
		}
		defer file.Close()

		// Since loading from source, point the entry point to containing folder.
		env.EntryLocation = filepath.Dir(file.Name())

		// Read the contents of the file.
		content, err := os.ReadFile(executable)
		if err != nil {
			e := fmt.Sprintf("Error reading file: %v", err)
			utils.Stdout(e, env.Stdout)
			return
		}

		// Run the program.
		result, err := program.Run(string(content), env)
		if err != nil {
			utils.Stdout(err.Error(), env.Stdout)
		} else {
			// Only really want to print to console if its a statement that needs returning.
			if result != nil {
				r := fmt.Sprintf("%v", result)
				fmt.Printf("Debug: %v\n", result)
				utils.Stdout(r, env.Stdout)
			}
		}

	} else {

		// REPL Mode.
		fmt.Println("Goblin v0.1")

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current working directory:", err)
			// Handle the error appropriately (e.g., exit)
		}

		// Set the Entry location to wherever the users terminal currently is.
		env.EntryLocation = cwd

		for {

			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				e := fmt.Sprintf("Error reading input: %v", err)
				utils.Stdout(e, env.Stdout)
				return
			}

			// Remove newline character
			input = strings.TrimSpace(input)

			if input == "exit" {
				utils.Stdout("Goodbye!", env.Stdout)
				return
			}

			// Run the program.
			result, err := program.Run(input, env)
			if err != nil {
				utils.Stdout(err.Error(), env.Stdout)
			} else {
				// Only really want to print to console if its a statement that needs returning.
				if result != nil {
					r := fmt.Sprintf("%v", result)
					utils.Stdout(r, env.Stdout)
				}
			}
		}
	}
}
