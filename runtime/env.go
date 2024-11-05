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
	e.Constants[var_] = isConst

	return value, nil
}

// Used to update an existing variable value.
func (e Environment) Update(var_ string, value RuntimeValue) (RuntimeValue, error) {

	_, exists := e.Variables[var_]

	// If this variable does not exit.
	if !exists {
		return nil, fmt.Errorf("unidentified variable: '%v'", var_)
	}

	// Update it.
	e.Variables[var_] = value

	return value, nil
}

// Used to declare a new array. Includes checking for arrays that might already existing.
func (e Environment) DeclareArray(var_ string, values []RuntimeValue, isConst bool) (RuntimeValue, error) {

	_, exists := e.Variables[var_]

	// If this array is already set.
	if exists {
		return nil, fmt.Errorf("'%v' already defined", var_)
	}

	// Make the array object here, which contains all the RuntimeValues the user specified.
	arr := MK_ARRAY(values)

	// If not, set it.
	e.Variables[var_] = arr

	// Add our constant variable.
	e.Constants[var_] = isConst

	return arr, nil
}

// Used to declare a new map. Includes checking for map that might already existing.
func (e Environment) DeclareMap(var_ string, values map[RuntimeValue]RuntimeValue, isConst bool) (RuntimeValue, error) {

	_, exists := e.Variables[var_]

	// If this map is already set.
	if exists {
		return nil, fmt.Errorf("'%v' already defined", var_)
	}

	// Make the array object here, which contains all the RuntimeValues the user specified.
	map_ := MK_MAP(values)

	// If not, set it.
	e.Variables[var_] = map_

	// Add our constant variable.
	e.Constants[var_] = isConst

	return map_, nil
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

func (e Environment) ArrayOrMapLookup(var_ string, i RuntimeValue) (RuntimeValue, error) {

	datastructure, err := e.Lookup(var_)
	if err != nil {
		return nil, err
	}

	// Dealing with an array.
	if e.IsArray(datastructure) {

		// Arrays can only use ints as their indexer.
		if index, ok := i.(NumberValue); ok {

			val, err := e.ArrayLookup(var_, index.Value)
			if err != nil {
				return nil, err
			}

			return val, nil

		} else {
			return nil, fmt.Errorf("array index must be of type int")
		}

	} else if e.IsMap(datastructure) {
		// Dealing with a map.

		val, err := e.MapLookup(var_, i)
		if err != nil {
			return nil, err
		}

		return val, nil

	} else {
		return nil, fmt.Errorf("unrecognised datastructure provided: %v", datastructure)
	}
}

// Used in the specific case of looking up individual elements in an array.
func (e Environment) ArrayLookup(var_ string, index int) (RuntimeValue, error) {

	arr_, err := e.Lookup(var_)
	if err != nil {
		return nil, err
	}

	if arr, ok := arr_.(ArrayValue); ok {

		// The specified index value is out of bounds!
		if index >= len(arr.Value) {

			e := fmt.Sprintf("index out of bounds for index %v \n", index)
			return MK_STRING(e), fmt.Errorf(e)
		}
		return arr.Value[index], nil
	}

	return nil, fmt.Errorf("invalid array: %v", arr_)
}

// Used in the specific case of looking up individual elements in a map
func (e Environment) MapLookup(var_ string, index RuntimeValue) (RuntimeValue, error) {

	mapp_, err := e.Lookup(var_)
	if err != nil {
		return nil, err
	}

	if mapp, ok := mapp_.(MapValue); ok {

		// The specified index value is out of bounds!
		val, ok := mapp.Value[index]
		if ok {
			return val, nil
		} else {
			return nil, fmt.Errorf("key `%v` does not exist for map: %v", index, mapp)
		}
	}

	return nil, fmt.Errorf("invalid map: %v", mapp_)
}

// Returns true if the Runtime value is an Array type.
func (e Environment) IsArray(r RuntimeValue) bool {

	if _, ok := r.(ArrayValue); ok {
		return true
	}

	return false
}

// Returns true if the Runtime value is a Map type.
func (e Environment) IsMap(r RuntimeValue) bool {

	if _, ok := r.(MapValue); ok {
		return true
	}

	return false
}

func (e Environment) Setup() {

	e.Declare("null", MK_NULL(), true)
	e.Declare("true", MK_BOOL(true), true)
	e.Declare("false", MK_BOOL(false), true)

	e.Declare("print", MK_NATIVE_FN(Print), true)
	e.Declare("println", MK_NATIVE_FN(Println), true)
}
