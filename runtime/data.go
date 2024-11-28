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
			Call: println,
		},
		"pop": {
			Type: "NativeFn",
			Call: printf,
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
		size = len(mapp.Value)
	}

	return MK_NUMBER(size), nil
}
