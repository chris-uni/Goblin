package irtypes

import "fmt"

type IRExecutionState struct {
	Storage     []IRValue
	Temporaries []IRValue
	Labels      []IRLabel
	PC          int
}

func (state *IRExecutionState) PushStorage(index int, val IRValue) {
	for len(state.Storage) <= index {
		state.Storage = append(state.Storage, nil)
	}

	state.Storage[index] = val
}

func (state *IRExecutionState) PushTemporaries(index int, val IRValue) {
	for len(state.Temporaries) <= index {
		state.Temporaries = append(state.Temporaries, nil)
	}

	state.Temporaries[index] = val
}

func (state *IRExecutionState) Resolve(i IRValue) (IRValue, error) {

	switch value := i.(type) {

	case IRNumber:
		return value, nil

	case IRString:
		return value, nil

	case IRBoolean:
		return value, nil

	case IRAddress:

		if value.Index < 0 || value.Index >= len(state.Storage) {
			return nil, fmt.Errorf("invalid IR address: @%v\n", value.Index)
		}

		if state.Storage[value.Index] == nil {
			return nil, fmt.Errorf("null pointer at IR address: @%v\n", value.Index)
		}

		return state.Storage[value.Index], nil

	case IRTemporary:

		if value.Index < 0 || value.Index >= len(state.Temporaries) {
			return nil, fmt.Errorf("invalid IR temporary: %%%v\n", value.Index)
		}

		if state.Temporaries[value.Index] == nil {
			return nil, fmt.Errorf("null pointer at IR temporary: @%v\n", value.Index)
		}

		return state.Temporaries[value.Index], nil
	}

	return nil, fmt.Errorf("no IRValue type found for %v\n", i)
}

func (es *IRExecutionState) String() string {

	builder := ""

	builder += fmt.Sprintf("Storage: %v\n", es.Storage)
	builder += fmt.Sprintf("Temporaries: %v\n", es.Temporaries)
	builder += fmt.Sprintf("Labels: %v\n", es.Labels)
	builder += fmt.Sprintf("Final PC: %v\n", es.PC)

	return builder
}

type IRExecutionResult struct {
	Value IRValue
}
