package program

import (
	"fmt"

	"goblin.org/main/frontend/lexer"
	"goblin.org/main/frontend/parser"
	runtime "goblin.org/main/runtime"
	"goblin.org/main/utils"
)

// Where the source goes to be lexed, parsed, interpreted, and returned.
func Run(input string, env runtime.Environment) (runtime.RuntimeValue, error) {

	// Stage 1. Lex the input.
	tokens, audit := lexer.Tokenize(input)

	fmt.Printf("Audit: %v\n\n", audit)

	// Stage 2. Produce the Abstract Syntax Tree.
	program, err := parser.ProduceAST(tokens, audit)
	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err.Error())
	}

	fmt.Printf("Program: %v\n\n", program.Body)

	// Stage 3. Interprete the AST.
	evaluation, err := runtime.Evaluate(program, env)
	if err != nil {
		return nil, fmt.Errorf("interpreter error: %v", err.Error())
	}

	if f, isNum := evaluation.(runtime.NativeFunction); isNum {

		r := fmt.Sprintf("%v\n", f.Call)
		utils.Stdout(r, env.Stdout)
		return nil, nil
	}

	return nil, nil
}
