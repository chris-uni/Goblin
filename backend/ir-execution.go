/*
Goblin IR Execution Engine v0.1
Author: Chris J.M. Wing
Date: 02/09/2026

Input:
	Optimised GoblinIR program.
Output:
	A fully executed GoblinIR program represented by a Goblin ExecutionState object.

Guarantees:
	- Execution of GoblinIR commands
*/

package backend

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
	a "goblin.org/main/middleware/irtypes/arithmetic"
	ct "goblin.org/main/middleware/irtypes/controlflow"
	m "goblin.org/main/middleware/irtypes/memory"
	p "goblin.org/main/middleware/irtypes/program"
)

func Dispatch(command i.IRCommand, state *i.IRExecutionState) error {

	var err error

	switch com := command.(type) {

	case *a.Add:
		err = ExecAdd(com, state)

	case *a.Sub:
		err = ExecSub(com, state)

	case *a.Mul:
		err = ExecMul(com, state)

	case *a.Div:
		err = ExecDiv(com, state)

	case *a.Mod:
		err = ExecMod(com, state)

	case *m.Store:
		err = ExecStore(com, state)

	case *m.Load:
		err = ExecLoad(com, state)

	case *ct.Jmp:
		err = ExecJmp(com, state)

	case *ct.JmpIf:
		err = ExecJmpIf(com, state)

	case *p.Return:
		err = ExecRtn(com, state)

	default:
		err = fmt.Errorf("execution: unknown command %v", command)
	}

	return err
}

/*
	ARITHMETIC OPERATION SECTION.
*/

func ExecAdd(add *a.Add, state *i.IRExecutionState) error {

	result, err := add.Exec(state)
	if err != nil {
		return err
	}
	state.PushTemporaries(add.Destination.Index, result)

	return nil
}

func ExecSub(sub *a.Sub, state *i.IRExecutionState) error {

	result, err := sub.Exec(state)
	if err != nil {
		return err
	}

	state.PushTemporaries(sub.Destination.Index, result)

	return nil
}

func ExecMul(mul *a.Mul, state *i.IRExecutionState) error {

	result, err := mul.Exec(state)
	if err != nil {
		return err
	}

	state.PushTemporaries(mul.Destination.Index, result)

	return nil
}

func ExecDiv(div *a.Div, state *i.IRExecutionState) error {

	result, err := div.Exec(state)
	if err != nil {
		return err
	}

	state.PushTemporaries(div.Destination.Index, result)

	return nil
}

func ExecMod(mod *a.Mod, state *i.IRExecutionState) error {

	result, err := mod.Exec(state)
	if err != nil {
		return err
	}

	state.PushTemporaries(mod.Destination.Index, result)

	return nil
}

/*
	MEMORY OPERATION SECTION.
*/

func ExecStore(store *m.Store, state *i.IRExecutionState) error {

	_, err := store.Exec(state)
	if err != nil {
		return err
	}

	return nil
}

func ExecLoad(load *m.Load, state *i.IRExecutionState) error {

	_, err := load.Exec(state)
	if err != nil {
		return err
	}

	return nil
}

/*
	CONTROL-FLOW OPERATION SECTION.
*/

func ExecJmp(jmp *ct.Jmp, state *i.IRExecutionState) error {

	_, err := jmp.Exec(state)
	if err != nil {
		return err
	}
	return nil
}

func ExecJmpIf(jmpif *ct.JmpIf, state *i.IRExecutionState) error {

	_, err := jmpif.Exec(state)
	if err != nil {
		return err
	}

	return nil
}

/*
	PROGRAM OPERATION SECTION.
*/

func ExecRtn(rtn *p.Return, state *i.IRExecutionState) error {

	_, err := rtn.Exec(state)
	if err != nil {
		return err
	}

	return nil
}

func Execution(commands []i.IRCommand) (i.IRExecutionResult, error) {

	state := i.IRExecutionState{
		Storage:     make([]i.IRValue, 0),
		Temporaries: make([]i.IRValue, 0),
		Labels:      make([]i.IRLabel, 0),
		Result:      nil,
		PC:          0,
	}

	for state.PC <= len(commands)-1 {

		err := Dispatch(commands[state.PC], &state)
		if err != nil {
			return i.IRExecutionResult{}, err
		}

		if state.Result != nil {

			return *state.Result, nil
		}

	}

	return i.IRExecutionResult{}, nil
}
