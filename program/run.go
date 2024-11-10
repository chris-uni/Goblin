package program

import (
	"fmt"

	"goblin.org/main/frontend/parser"
	runtime "goblin.org/main/runtime"
	"goblin.org/main/utils"
)

// Where the source goes to be lexed, parsed, interpreted, and returned.
func Run(input string, env runtime.Environment) (runtime.RuntimeValue, error) {

	program, err := parser.ProduceAST(input)
	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err.Error())
	}

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
