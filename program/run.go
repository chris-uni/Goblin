package program

import (
	"fmt"

	"goblin.org/main/frontend/parser"
	"goblin.org/main/runtime"
)

// Where the source goes to be lexed, parsed, interpreted, and returned.
func Run(input string, env runtime.Environment) (runtime.RuntimeValue, error) {

	program, err := parser.ProduceAST(input)
	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err.Error())
	}

	fmt.Printf("%v \n", program)

	evaluation, err := runtime.Evaluate(program, env)
	if err != nil {
		return nil, fmt.Errorf("interpreter error: %v", err.Error())
	}

	if f, isNum := evaluation.(runtime.NativeFunction); isNum {

		fmt.Printf("%v\n", f.Call)
		return nil, nil
	}

	return nil, nil
}
