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

	default:
		return com, nil
	}
}

/*
Validates Store, simply commits value to storage.
*/
func validateStore(s *Store, context *IRContext) (IRCommand, error) {

	/*
		For store, we are effectively creating a new variable, so we dont need to check if the address exists, as it probably doesnt.
	*/
	context.Storage = append(context.Storage, s.Value)
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

	context.Temporaries = append(context.Temporaries, l.Source)
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

	if lhsType != rhsType {
		return nil, fmt.Errorf("type error: add: incompatible types\n")
	}

	/*
		Since for the Add operation we can perform on any type (number, string, boolean), and we can only,
		perform an add operation on two values of the same type, we simply commit the lhs value to temporaries.

		This is because, simply commiting a dummy number value might cause validation errors in more complex
		assemblies (i.e. pulling a temp of dummy IRNumber{} when in reality it resolves to an IRString{}).
	*/
	context.Temporaries = append(context.Temporaries, a.Lhs)

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

	// Since we know sub can only perform calculation on a number, lets just add the lhs value in
	// as a place holder.
	context.Temporaries = append(context.Temporaries, s.Lhs)

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

	// Since we know mul can only perform calculation on a number, lets just add the lhs value in
	// as a place holder.
	context.Temporaries = append(context.Temporaries, m.Lhs)

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

	// Since we know div can only perform calculation on a number, lets just add the lhs value in
	// as a place holder.
	context.Temporaries = append(context.Temporaries, d.Lhs)

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

	// Since we know mod can only perform calculation on a number, lets just add the lhs value in
	// as a place holder.
	context.Temporaries = append(context.Temporaries, m.Lhs)

	return m, nil
}

func Validate(commands []IRCommand) ([]IRCommand, error) {

	context := IRContext{
		Commands:    make([]IRCommand, 0),
		Storage:     make([]IRValue, 0),
		Temporaries: make([]IRValue, 0),
		Symbols:     make(map[string]IRAddress),
		PC:          0,
	}

	for _, command := range commands {

		com, err := validateCommand(command, &context)
		if err != nil {
			return []IRCommand{}, fmt.Errorf("validation error: %v\n", err)
		}

		context.Commands = append(context.Commands, com)
	}

	return context.Commands, nil
}
