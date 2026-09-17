/*
Goblin IR Execution Engine v0.1
Author: Chris J.M. Wing
Date: 02/09/2026

Input:
	Optimised GoblinIR program.
Output:
	A fully executed GoblinIR program represented by a GoblinRuntime object.

Guarantees:
	- Execution of GoblinIR commands
*/

package backend

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
	a "goblin.org/main/middleware/irtypes/arithmetic"
)

func Dispatch(command i.IRCommand, state *i.IRExecutionState) error {

	switch com := command.(type) {

	case *a.Add:
		return ExecAdd(com, state)

	default:
		return fmt.Errorf("execution: unknown command %v", command)
	}
}

func ExecAdd(add *a.Add, state *i.IRExecutionState) error {

	result := add.Exec(state)
	fmt.Printf("execution: `add` result %v\n", result)

	return nil
}

func Execution(commands []i.IRCommand) error {

	state := i.IRExecutionState{
		Storage:     make([]i.IRValue, 0),
		Temporaries: make([]i.IRValue, 0),
		Labels:      make([]i.IRLabel, 0),
		PC:          0,
	}

	for _, com := range commands {

		err := Dispatch(com, &state)
		if err != nil {
			return err
		}
	}

	return nil
}
