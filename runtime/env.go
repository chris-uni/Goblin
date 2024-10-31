package runtime

import (
	"fmt"
	"io"
)

type Environment struct {
	Parent    *Environment
	Stdout    io.Writer
	Variables map[string]RuntimeValue
	Constants map[string]bool
}

// Used to declare a new variable. Includes checking for variable already existing.
func (e Environment) Declare(var_ string, value RuntimeValue, isConst bool) (RuntimeValue, error) {

	_, exists := e.Variables[var_]

	// If this variable is already set.
	if exists {
		return nil, fmt.Errorf("'%v' already defined", var_)
	}

	// If not, set it.
	e.Variables[var_] = value

	// Add our constant variable.
	if isConst {
		e.Constants[var_] = true
	}

	return value, nil
}

// Used to declare a new array. Includes checking for arrays that might already existing.
func (e Environment) DeclareArray(var_ string, values []RuntimeValue, isConst bool) (RuntimeValue, error) {

	_, exists := e.Variables[var_]

	// If this variable is already set.
	if exists {
		return nil, fmt.Errorf("'%v' already defined", var_)
	}

	// Make the array object here, which contains all the RuntimeValues the user specified.
	arr := MK_ARRAY(values)

	// If not, set it.
	e.Variables[var_] = arr

	// Add our constant variable.
	if isConst {
		e.Constants[var_] = true
	}

	return arr, nil
}

// Used to assign values to a variable.
func (e Environment) Assign(var_ string, value RuntimeValue) (RuntimeValue, error) {

	env, err := e.Resolve(var_)
	if err != nil {
		return nil, err
	}

	hasConstant := e.Constants[var_]

	if hasConstant {
		// Cannot assign to a constant.
		return nil, fmt.Errorf("cannot reassign const value '%v'", var_)
	}

	env.Variables[var_] = value

	return value, nil
}

// Used to find the specific Environment a variable is located in (scope resolution).
func (e Environment) Resolve(var_ string) (Environment, error) {

	// Variable in this scope?
	_, exists := e.Variables[var_]
	if exists {
		return e, nil
	}

	// No parent exists in scope.
	if e.Parent == nil {
		return Environment{}, fmt.Errorf("reference to undefined variable '%v'", var_)
	}

	res, err := e.Parent.Resolve(var_)
	if err != nil {
		return Environment{}, err
	}
	// Check the parent scope.
	return res, nil
}

// Returns the value of the variable.
func (e Environment) Lookup(var_ string) (RuntimeValue, error) {

	env, err := e.Resolve(var_)
	if err != nil {
		return nil, err
	}

	return env.Variables[var_], nil
}

// Used in the specific case of looking up individual elements in an array.
func (e Environment) ArrayLookup(var_ string, index int) (RuntimeValue, error) {

	arr_, err := e.Lookup(var_)
	if err != nil {
		return nil, err
	}

	var value RuntimeValue
	if arr, ok := arr_.(ArrayValue); ok {

		// The specified index value is out of bounds!
		if index >= len(arr.Value) {

			e := fmt.Sprintf("index out of bounds for index %v \n", index)
			return MK_STRING(e), fmt.Errorf(e)
		}
		value = arr.Value[index]
	}

	return value, nil
}

func (e Environment) Setup() {

	e.Declare("null", MK_NULL(), true)
	e.Declare("true", MK_BOOL(true), true)
	e.Declare("false", MK_BOOL(false), true)

	e.Declare("print", MK_NATIVE_FN(Print), true)
	e.Declare("println", MK_NATIVE_FN(Println), true)
}
