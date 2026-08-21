package middleware

import (
	"fmt"
)

type IRType int

const (
	IRTypeNumber IRType = iota
	IRTypeString
	IRTypeBoolean
	IRTypeUndefined
)

func (t IRType) String() string {
	switch t {
	case IRTypeNumber:
		return "number"
	case IRTypeString:
		return "string"
	case IRTypeBoolean:
		return "boolean"
	default:
		return "unknown"
	}
}

type IRContext struct {
	Commands []IRCommand

	Storage     []IRValue
	Temporaries []IRValue

	Symbols map[string]IRAddress

	PC int
}

type IRResult struct {
	Commands []IRCommand
	Value    IRValue
}

func (engine *IRContext) resolve(i IRValue) (IRValue, error) {

	switch value := i.(type) {

	case IRNumber:
		return value, nil

	case IRString:
		return value, nil

	case IRBoolean:
		return value, nil

	case IRAddress:

		if value.Index < 0 || value.Index >= len(engine.Storage) {
			return nil, fmt.Errorf("invalid IR address: @%v\n", value.Index)
		}

		if engine.Storage[value.Index] == nil {
			return nil, fmt.Errorf("null pointer at IR address: @%v\n", value.Index)
		}

		return engine.Storage[value.Index], nil

	case IRTemporary:

		if value.Index < 0 || value.Index >= len(engine.Temporaries) {
			return nil, fmt.Errorf("invalid IR temporary: %%%v\n", value.Index)
		}

		if engine.Temporaries[value.Index] == nil {
			return nil, fmt.Errorf("null pointer at IR temporary: @%v\n", value.Index)
		}

		return engine.Temporaries[value.Index], nil
	}

	return nil, fmt.Errorf("no IRValue type found for %v\n", i)
}
