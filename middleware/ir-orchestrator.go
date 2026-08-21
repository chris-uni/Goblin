package middleware

import (
	"goblin.org/main/frontend/ast"
)

func OrchestrateIRLayer(program ast.Program) ([]IRCommand, error) {

	// 1. Reduce the AST down into GoblinIR.
	rawIR, err := Reduce(program)
	if err != nil {
		return []IRCommand{}, err
	}

	PrintIR("raw ir:", rawIR)

	// 2. Validate the Raw GoblinIR into Validated GoblinIR.
	validatedIR, err := Validate(rawIR)
	if err != nil {
		return []IRCommand{}, err
	}

	return validatedIR, nil
}
