package runtime

import (
	"fmt"
	"strings"
)

var Strings = Namespace{
	Name: "data",
	Functions: map[string]NativeFunction{
		"split": {
			Type: "NativeFn",
			Call: split,
		},
	},
}

// split, splits string `s` by delimiter `d`, returns an array of sub-string elements.
// strings.split(s str, d str)
var split FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 2 {
		return nil, fmt.Errorf("unexpected number of args for strings.split, expected 2 got %v", numArgs)
	}

	s := args[0]
	str, isStr := s.(StringValue)
	if !isStr {
		return nil, fmt.Errorf("strings.split must be used on string type, %v type given", str.Type)
	}

	d := args[1]
	del, isStr := d.(StringValue)
	if !isStr {
		return nil, fmt.Errorf("strings.split delimiter must be string type, %v type given", del.Type)
	}

	splits := strings.Split(str.Value, del.Value)

	sRuntime := make([]RuntimeValue, 0)
	for _, st := range splits {

		sRuntime = append(sRuntime, StringValue{Value: st, Type: "StringValue"})
	}

	return MK_ARRAY(sRuntime), nil
}
