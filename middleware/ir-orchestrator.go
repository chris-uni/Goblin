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

	PrintIR("validated ir:", rawIR)

	return validatedIR, nil
}
