package runtime

import "fmt"

var Data = Namespace{
	Name: "data",
	Functions: map[string]NativeFunction{
		"push": {
			Type: "NativeFn",
			Call: push,
		},
		"put": {
			Type: "NativeFn",
			Call: put,
		},
		"pop": {
			Type: "NativeFn",
			Call: pop,
		},
		"size": {
			Type: "NativeFn",
			Call: size,
		},
	},
}

// push, pushes a new value into an array (top-down)
// data.push(arr array, val any)
var push FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 2 {
		return nil, fmt.Errorf("unexpected number of args for data.push, expected 2 got %v", numArgs)
	}

	a := args[0]
	arr, isArr := a.(ArrayValue)
	if !isArr {
		return nil, fmt.Errorf("data.push must be used on array type, %v type given", arr.Type)
	}

	// The value we want to push into the array.
	value := args[1]
	*arr.Value = append(*arr.Value, value)

	return MK_NULL(), nil
}

// put, puts a new key/value pair into a map, inserted at the end
// data.put(m map, key any, value any)
var put FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 3 {
		return nil, fmt.Errorf("unexpected number of args for data.put, expected 3 got %v", numArgs)
	}

	m := args[0]
	mapp, isMap := m.(MapValue)
	if !isMap {
		return nil, fmt.Errorf("data.put must be used on map type, %v type given", mapp.Type)
	}

	// The value we want to push into the array.
	key := args[1]
	value := args[2]

	tmp := *mapp.Value
	tmp[key] = value

	return MK_NULL(), nil
}

// pop, returns the last element of the specified array
// data.pop(a array)
var pop FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for data.pop, expected 1 got %v", numArgs)
	}

	a := args[0]
	arr, isArr := a.(ArrayValue)

	if !isArr {
		return nil, fmt.Errorf("unexpected type provided for data.pop, got %v", a)
	}

	if len(*arr.Value) > 0 {

		lastIndex := len(*arr.Value) - 1
		lastItem := (*arr.Value)[lastIndex]
		*arr.Value = (*arr.Value)[:lastIndex]

		return lastItem, nil
	}

	return nil, fmt.Errorf("cannot pop an empty array")
}

// size, returns the size of the array or map specified
// data.size(a array), data.size(m map)
var size FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for data.size, expected 1 got %v", numArgs)
	}

	a := args[0]
	arr, isArr := a.(ArrayValue)
	mapp, isMap := a.(MapValue)

	if !isArr && !isMap {
		return nil, fmt.Errorf("unexpected type provided for data.size, got %v", a)
	}

	size := 0
	if isArr {
		size = len(*arr.Value)
	} else if isMap {
		size = len(*mapp.Value)
	}

	return MK_NUMBER(size), nil
}
