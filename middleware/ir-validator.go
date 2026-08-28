/*
Goblin IR Validator v0.1
Author: Chris J.M. Wing
Date: 21/08/2026

Input:
	Raw GoblinIR from the Reducer stage.
Output:
	A set of valid i.IRCommands that represent a validated GoblinIR program.
*/

package middleware

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
	a "goblin.org/main/middleware/irtypes/arithmetic"
)

func validateCommand(command i.IRCommand, context *i.IRContext) (i.IRCommand, error) {

	switch com := command.(type) {

	case *a.Add:
		_, err := validateAdd(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *a.Sub:
		_, err := validateSub(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *a.Mul:
		_, err := validateMul(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *a.Div:
		_, err := validateDiv(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *a.Mod:
		_, err := validateMod(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Store:
		_, err := validateStore(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Load:
		_, err := validateLoad(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Eq:
		_, err := validateEquality(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Neq:
		_, err := validateNotEquality(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Lt:
		_, err := validateLessThan(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Lte:
		_, err := validateLessThanEqualTo(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Gt:
		_, err := validateGreaterThan(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.Gte:
		_, err := validateGreaterThanEqualTo(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *i.JmpIf:
		_, err := validateJmpIf(com, context)
		if err != nil {
			return nil, err
		}

		return com, nil

	case *i.Jmp:
		_, err := validateJmp(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	default:
		return nil, fmt.Errorf("unrecognised ir command %v\n", command)
	}
}

/*
Validates Store, simply commits value to storage.
*/
func validateStore(s *i.Store, context *i.IRContext) (i.IRCommand, error) {
	return s, nil
}

/*
Validates Load, checks if source is a valid address space, then commits source to temporaries.
*/
func validateLoad(l *i.Load, context *i.IRContext) (i.IRCommand, error) {

	/*
		For load, we need to check the source exists as we are loading a stored variable into temporary memory.
	*/
	_, err := resolveIRType(l.Source, context)
	if err != nil {
		return nil, err
	}

	return l, nil
}

/*
Validates Add, performs type checking on provided operands, resolving where needed.
*/
func validateAdd(a *a.Add, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(a.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(a.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeString) &&
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeString) {
		return nil, fmt.Errorf("type error: add: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return nil, fmt.Errorf("type error: add: incompatible types\n")
	}

	return a, nil
}

/*
Validates Sub, performs type checking on provided operands, resolving where needed.
*/
func validateSub(s *a.Sub, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(s.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(s.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: sub: operands of invalid type\n")
	}

	return s, nil
}

/*
Validates Mul, performs type checking on provided operands, resolving where needed.
*/
func validateMul(m *a.Mul, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: mul: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Div, performs type checking on provided operands, resolving where needed.
*/
func validateDiv(d *a.Div, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(d.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(d.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: div: operands of invalid type\n")
	}

	return d, nil
}

/*
Validates Mod, performs type checking on provided operands, resolving where needed.
*/
func validateMod(m *a.Mod, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: mod: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Eq, performs type checking on provided operands, resolving where needed.
*/
func validateEquality(m *i.Eq, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeBoolean) &&
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeBoolean) {
		return nil, fmt.Errorf("type error: eq: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return nil, fmt.Errorf("type error: eq: incompatible types\n")
	}

	return m, nil
}

/*
Validates Neq, performs type checking on provided operands, resolving where needed.
*/
func validateNotEquality(m *i.Neq, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeBoolean) &&
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeBoolean) {
		return nil, fmt.Errorf("type error: eq: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return nil, fmt.Errorf("type error: eq: incompatible types\n")
	}

	return m, nil
}

/*
Validates Lt, performs type checking on provided operands, resolving where needed.
*/
func validateLessThan(m *i.Lt, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: lt: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Lte, performs type checking on provided operands, resolving where needed.
*/
func validateLessThanEqualTo(m *i.Lte, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: lte: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Gt, performs type checking on provided operands, resolving where needed.
*/
func validateGreaterThan(m *i.Gt, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: gt: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Gte, performs type checking on provided operands, resolving where needed.
*/
func validateGreaterThanEqualTo(m *i.Gte, context *i.IRContext) (i.IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return nil, fmt.Errorf("type error: gte: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates JmpIf, performs type checking on provided operands, resolving where needed.
*/
func validateJmpIf(ji *i.JmpIf, context *i.IRContext) (i.IRCommand, error) {

	conditionType, err := resolveIRType(ji.Condition, context)
	if err != nil {
		return nil, err
	}

	_, err = resolveIRType(ji.Destination, context)
	if err != nil {
		return nil, err
	}

	// Is the condition a truthy type?
	if conditionType != i.IRTypeBoolean {
		return nil, fmt.Errorf("type error: jmpif: condition of invalid type\n")
	}

	return ji, nil
}

/*
Validates JmpIf, performs type checking on provided operands, resolving where needed.
*/
func validateJmp(j *i.Jmp, context *i.IRContext) (i.IRCommand, error) {

	_, err := resolveIRType(j.Destination, context)
	if err != nil {
		return nil, err
	}

	return j, nil
}

/*
Resolves incoming IRValue type to a compariable IRType value. Will recursively resolve IRAddress and IRTemporary values.
*/
func resolveIRType(value i.IRValue, context *i.IRContext) (i.IRType, error) {

	switch val := value.(type) {

	case i.IRNumber:
		return i.IRTypeNumber, nil

	case i.IRString:
		return i.IRTypeString, nil

	case i.IRBoolean:
		return i.IRTypeBoolean, nil

	case i.IRAddress:
		if val.Index < 0 || val.Index >= len(context.Storage) {
			return i.IRTypeUndefined, fmt.Errorf("undefined storage address @%d\n", val.Index)
		}
		return resolveIRType(context.Storage[val.Index], context)

	case i.IRTemporary:
		if val.Index < 0 || val.Index >= len(context.Temporaries) {
			return i.IRTypeUndefined, fmt.Errorf("undefined temporary address %%%d\n", val.Index)
		}
		return resolveIRType(context.Temporaries[val.Index], context)

	case i.IRLabel:

		if val.Value < 0 || val.Value >= len(context.Labels) {
			return i.IRTypeUndefined, fmt.Errorf("label[%v] index out of bounds for value %v\n", val, val.Value)
		}

		if val.PCOffset < 0 || val.PCOffset > len(context.Commands) {
			return i.IRTypeUndefined, fmt.Errorf("label[%v] offset out of bounds for offset %v\n", val, val.PCOffset)
		}

		return i.IRTypeLabel, nil

	default:
		return i.IRTypeUndefined, fmt.Errorf("unrecognised type %v\n", value)
	}
}

func Validate(commands []i.IRCommand, context *i.IRContext) ([]i.IRCommand, error) {

	context.PC = 0

	for _, command := range commands {

		_, err := validateCommand(command, context)
		if err != nil {
			return []i.IRCommand{}, fmt.Errorf("validation error: %v\n", err)
		}
	}

	return context.Commands, nil
}
