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
	m "goblin.org/main/middleware/irtypes/memory"
)

func Dispatch(command i.IRCommand, state *i.IRExecutionState) error {

	switch com := command.(type) {

	case *a.Add:
		return ExecAdd(com, state)

	case *m.Store:
		return ExecStore(com, state)

	case *m.Load:
		return ExecLoad(com, state)

	default:
		return fmt.Errorf("execution: unknown command %v", command)
	}
}

func ExecAdd(add *a.Add, state *i.IRExecutionState) error {

	result := add.Exec(state)
	state.PushTemporaries(add.Destination.Index, result)

	return nil
}

func ExecStore(store *m.Store, state *i.IRExecutionState) error {

	val, err := state.Resolve(store.Value)
	if err != nil {
		return err
	}

	state.PushStorage(store.Destination.Index, val)
	return nil
}

func ExecLoad(load *m.Load, state *i.IRExecutionState) error {

	index := load.Destination.Index

	val, err := state.Resolve(load.Source)
	if err != nil {
		return err
	}

	fmt.Printf("pushing value %v into temporary index %v\n", val, index)

	state.PushTemporaries(index, val)
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

	fmt.Printf("%v\n", state.String())

	return nil
}
