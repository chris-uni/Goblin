package middleware

import (
	"fmt"
)

/*
Stages of validator:

1. Invalid address (storage)
2. Invalid temporary
3. Type errors
4. Invalid instruction operands
5. Invalid control flow
*/

func handler(command IRCommand, context *IRContext) (IRCommand, error) {

	com, err := validateAddress(command, context)
	if err != nil {
		return nil, err
	}

	com, err = validateTemporary(com, context)
	if err != nil {
		return nil, err
	}

	com, err = validateTypes(com, context)
	if err != nil {
		return nil, err
	}

	com, err = storeAndLoad(command, context)
	if err != nil {
		return nil, err
	}

	com, err = validateInstrOps(com, context)
	if err != nil {
		return nil, err
	}

	fmt.Printf("successfully validated %v\n", com)

	return com, nil
}

func storeAndLoad(command IRCommand, context *IRContext) (IRCommand, error) {

	switch com := command.(type) {

	/*case *Add:

	index := context.allocateTemporary()
	context.Temporaries[index] = com.Lhs + com.Rhs*/

	case *Store:
		/*
			Store puts a value into an address.
		*/
		value, err := context.resolve(com.Value)
		if err != nil {
			return nil, err
		}

		context.Storage = append(context.Storage, value)

		return com, nil

	case *Load:
		/*
			Load takes an addressed value and puts it into a temporary.
		*/
		address := com.Source
		val := context.Storage[address.Index]
		context.Temporaries = append(context.Temporaries, val)

		return com, nil
	default:
		return com, nil
	}
}

func validateAddress(command IRCommand, _ *IRContext) (IRCommand, error) {

	return command, nil
}

func validateTemporary(command IRCommand, _ *IRContext) (IRCommand, error) {

	return command, nil
}

/*
Performs command dependant type verfication against the GoblinIR opcode spec.
*/
func validateTypes(command IRCommand, context *IRContext) (IRCommand, error) {

	switch c := command.(type) {

	case *Add, *Sub, *Mul, *Div, *Mod:

		bc := c.(IRBinaryCommand)
		lhs := bc.Left()
		rhs := bc.Right()

		if !(isNumber(bc.Left()) || !isAddress(bc.Left()) || !isTemporary(bc.Left())) {
			return nil, fmt.Errorf("invalid type provided for lhs of command add: %v\n", lhs)
		}

		if !(isNumber(bc.Right()) || !isAddress(bc.Right()) || !isTemporary(bc.Right())) {
			return nil, fmt.Errorf("invalid type provided for rhs of command add: %v\n", rhs)
		}

		// Type checking.

		lhsType, err := resolveIRType(lhs, context)
		if err != nil {
			return nil, err
		}

		rhsType, err := resolveIRType(rhs, context)
		if err != nil {
			return nil, err
		}

		if lhsType != rhsType {
			return nil, fmt.Errorf("type mismatch on command %v\n", command)
		}

		return command, nil

	case *Store:

		if !isAddress(c.Destination) {
			return nil, fmt.Errorf("invalid address specified for store command")
		}

		return command, nil

	case *Load:

		if !isTemporary(c.Destination) {
			return nil, fmt.Errorf("invalid temporary for command load")
		}

		if !isAddress(c.Source) {
			return nil, fmt.Errorf("invalid address specified for load command")
		}

		return command, nil

	default:
		return nil, fmt.Errorf("invalidy type for operation %v\n", command)
	}
}

func resolveIRType(value IRValue, context *IRContext) (IRType, error) {

	switch value := value.(type) {

	case IRNumber:
		return IRTypeNumber, nil

	case IRString:
		return IRTypeString, nil

	case IRBoolean:
		return IRTypeBoolean, nil

	case IRAddress:
		if value.Index < 0 || value.Index >= len(context.Storage) {
			return IRTypeUndefined, fmt.Errorf("undefined storage address @%d\n", value.Index)
		}
		return resolveIRType(context.Storage[value.Index], context)

	case IRTemporary:
		if value.Index < 0 || value.Index >= len(context.Temporaries) {
			return IRTypeUndefined, fmt.Errorf("undefined temporary address %%%d\n", value.Index)
		}
		return resolveIRType(context.Temporaries[value.Index], context)

	default:
		return IRTypeUndefined, fmt.Errorf("unrecognised type %v\n", value)
	}
}

func isNumber(value IRValue) bool {
	_, ok := value.(IRNumber)
	return ok
}

func isAddress(value IRValue) bool {
	_, ok := value.(IRAddress)
	return ok
}

func isTemporary(value IRValue) bool {
	_, ok := value.(IRTemporary)
	return ok
}

func validateInstrOps(command IRCommand, _ *IRContext) (IRCommand, error) {

	return command, nil
}

func Validate2(commands []IRCommand) ([]IRCommand, error) {

	context := IRContext{
		Commands:    make([]IRCommand, 0),
		Storage:     make([]IRValue, 0),
		Temporaries: make([]IRValue, 0),
		Symbols:     make(map[string]IRAddress),
		PC:          0,
	}

	for _, command := range commands {

		com, err := handler(command, &context)
		if err != nil {
			return []IRCommand{}, fmt.Errorf("validation error: %v\n", err)
		}

		context.Commands = append(context.Commands, com)
	}

	return context.Commands, nil
}
