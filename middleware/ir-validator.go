/*
Goblin IR Validator v0.1
Author: Chris J.M. Wing
Date: 28/08/2026

Input:
	Raw GoblinIR from the Reducer stage.
Output:
	A set of valid IRCommands that represent a validated GoblinIR program.
*/

package middleware

import "fmt"

func validateCommand(command IRCommand, context *IRContext) (IRCommand, error) {

	switch com := command.(type) {

	case *Add:
		_, err := validateAdd(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Sub:
		_, err := validateSub(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Mul:
		_, err := validateMul(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Div:
		_, err := validateDiv(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Mod:
		_, err := validateMod(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Store:
		_, err := validateStore(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Load:
		_, err := validateLoad(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Eq:
		_, err := validateEquality(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Neq:
		_, err := validateNotEquality(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Lt:
		_, err := validateLessThan(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Lte:
		_, err := validateLessThanEqualTo(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Gt:
		_, err := validateGreaterThan(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *Gte:
		_, err := validateGreaterThanEqualTo(com, context)
		if err != nil {
			return nil, err
		}
		return com, nil

	case *JmpIf:
		_, err := validateJmpIf(com, context)
		if err != nil {
			return nil, err
		}

		return com, nil

	case *Jmp:
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
func validateStore(s *Store, context *IRContext) (IRCommand, error) {
	return s, nil
}

/*
Validates Load, checks if source is a valid address space, then commits source to temporaries.
*/
func validateLoad(l *Load, context *IRContext) (IRCommand, error) {

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
func validateAdd(a *Add, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(a.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(a.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == IRTypeNumber || lhsType == IRTypeString) &&
		!(rhsType == IRTypeNumber || rhsType == IRTypeString) {
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
func validateSub(s *Sub, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(s.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(s.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: sub: operands of invalid type\n")
	}

	return s, nil
}

/*
Validates Mul, performs type checking on provided operands, resolving where needed.
*/
func validateMul(m *Mul, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: mul: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Div, performs type checking on provided operands, resolving where needed.
*/
func validateDiv(d *Div, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(d.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(d.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: div: operands of invalid type\n")
	}

	return d, nil
}

/*
Validates Mod, performs type checking on provided operands, resolving where needed.
*/
func validateMod(m *Mod, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: mod: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Eq, performs type checking on provided operands, resolving where needed.
*/
func validateEquality(m *Eq, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == IRTypeNumber || lhsType == IRTypeBoolean) &&
		!(rhsType == IRTypeNumber || rhsType == IRTypeBoolean) {
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
func validateNotEquality(m *Neq, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if !(lhsType == IRTypeNumber || lhsType == IRTypeBoolean) &&
		!(rhsType == IRTypeNumber || rhsType == IRTypeBoolean) {
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
func validateLessThan(m *Lt, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: lt: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Lte, performs type checking on provided operands, resolving where needed.
*/
func validateLessThanEqualTo(m *Lte, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: lte: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Gt, performs type checking on provided operands, resolving where needed.
*/
func validateGreaterThan(m *Gt, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: gt: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates Gte, performs type checking on provided operands, resolving where needed.
*/
func validateGreaterThanEqualTo(m *Gte, context *IRContext) (IRCommand, error) {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := resolveIRType(m.Lhs, context)
	if err != nil {
		return nil, err
	}

	rhsType, err := resolveIRType(m.Rhs, context)
	if err != nil {
		return nil, err
	}

	if lhsType != IRTypeNumber || rhsType != IRTypeNumber {
		return nil, fmt.Errorf("type error: gte: operands of invalid type\n")
	}

	return m, nil
}

/*
Validates JmpIf, performs type checking on provided operands, resolving where needed.
*/
func validateJmpIf(ji *JmpIf, context *IRContext) (IRCommand, error) {

	conditionType, err := resolveIRType(ji.Condition, context)
	if err != nil {
		return nil, err
	}

	_, err = resolveIRType(ji.Destination, context)
	if err != nil {
		return nil, err
	}

	// Is the condition a truthy type?
	if conditionType != IRTypeBoolean {
		return nil, fmt.Errorf("type error: jmpif: condition of invalid type\n")
	}

	return ji, nil
}

/*
Validates JmpIf, performs type checking on provided operands, resolving where needed.
*/
func validateJmp(j *Jmp, context *IRContext) (IRCommand, error) {

	_, err := resolveIRType(j.Destination, context)
	if err != nil {
		return nil, err
	}

	return j, nil
}

/*
Resolves incoming IRValue type to a compariable IRType value. Will recursively resolve IRAddress and IRTemporary values.
*/
func resolveIRType(value IRValue, context *IRContext) (IRType, error) {

	switch val := value.(type) {

	case IRNumber:
		return IRTypeNumber, nil

	case IRString:
		return IRTypeString, nil

	case IRBoolean:
		return IRTypeBoolean, nil

	case IRAddress:
		if val.Index < 0 || val.Index >= len(context.Storage) {
			return IRTypeUndefined, fmt.Errorf("undefined storage address @%d\n", val.Index)
		}
		return resolveIRType(context.Storage[val.Index], context)

	case IRTemporary:
		if val.Index < 0 || val.Index >= len(context.Temporaries) {
			return IRTypeUndefined, fmt.Errorf("undefined temporary address %%%d\n", val.Index)
		}
		return resolveIRType(context.Temporaries[val.Index], context)

	case IRLabel:

		if val.Value < 0 || val.Value >= len(context.Labels) {
			return IRTypeUndefined, fmt.Errorf("label[%v] index out of bounds for value %v\n", val, val.Value)
		}

		if val.PCOffset < 0 || val.PCOffset > len(context.Commands) {
			return IRTypeUndefined, fmt.Errorf("label[%v] offset out of bounds for offset %v\n", val, val.PCOffset)
		}

		return IRTypeLabel, nil

	default:
		return IRTypeUndefined, fmt.Errorf("unrecognised type %v\n", value)
	}
}

func Validate(commands []IRCommand, context *IRContext) ([]IRCommand, error) {

	context.PC = 0

	for _, command := range commands {

		_, err := validateCommand(command, context)
		if err != nil {
			return []IRCommand{}, fmt.Errorf("validation error: %v\n", err)
		}
	}

	return context.Commands, nil
}
