/*
Goblin IR Orchestrator v0.1
Author: Chris J.M. Wing
Date: 21/08/2026

Input:
	Validated GoblinIR program.
Output:
	A set of IRCommands that represent an optimised version of the original GoblinIR program.

Guarantees:
	- Unaltered functionality of source program
*/

package middleware

import (
	"goblin.org/main/frontend/ast"
)

func OrchestrateIRLayer(program ast.Program) ([]IRCommand, error) {

	context := IRContext{
		Commands:    make([]IRCommand, 0),
		Storage:     make([]IRValue, 0),
		Temporaries: make([]IRValue, 0),
		Labels:      make([]IRLabel, 0),
		Symbols:     make(map[string]IRAddress),
		PC:          0,
	}

	// 1. Reduce the AST down into GoblinIR.
	rawIR, err := Reduce(program, &context)
	if err != nil {
		return []IRCommand{}, err
	}

	PrintIR("raw ir:", rawIR)

	// 2. Validate the Raw GoblinIR into Validated GoblinIR.
	validatedIR, err := Validate(rawIR, &context)
	if err != nil {
		return []IRCommand{}, err
	}

	PrintIR("validated ir:", rawIR)

	return validatedIR, nil
}
