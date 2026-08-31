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
	i "goblin.org/main/middleware/irtypes"
)

func OrchestrateIRLayer(program ast.Program) ([]i.IRCommand, error) {

	context := i.IRContext{
		Commands:    make([]i.IRCommand, 0),
		Storage:     make([]i.IRValue, 0),
		Temporaries: make([]i.IRValue, 0),
		Labels:      make([]i.IRLabel, 0),
		Symbols:     make(map[string]i.IRAddress),
		PC:          0,
	}

	// 1. Reduce the AST down into GoblinIR.
	rawIR, err := Reduce(program, &context)
	if err != nil {
		return []i.IRCommand{}, err
	}

	PrintIR("raw ir:", rawIR)

	// 2. Validate the Raw GoblinIR into Validated GoblinIR.
	validatedIR, err := Validate(rawIR, &context)
	if err != nil {
		return []i.IRCommand{}, err
	}

	PrintIR("validated ir:", rawIR)

	return validatedIR, nil
}
