package runtime

import (
	"fmt"

	"goblin.org/main/utils"
)

var IO = Namespace{
	Name: "io",
	Functions: map[string]NativeFunction{
		"print": {
			Type: "NativeFn",
			Call: print,
		},
		"println": {
			Type: "NativeFn",
			Call: println,
		},
	},
}

// Defines the built-in method 'print'.
// Depending on arg value type, print is handled in different ways.
var print FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	utils.Stdout(builder, env.Stdout)

	return MK_NULL()
}

// Defines the built-in method 'println'.
// Same as 'print' but adds a '\n' char at the end of the output.
var println FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	builder += "\n"

	utils.Stdout(builder, env.Stdout)

	return MK_NULL()
}

// Helper funcition for the 'print' built-in function defined above.
// Recursive function to identifiy what type wants to be printed
// and handles accordingly.
func printHelper(arg RuntimeValue) string {

	builder := ""

	if num, ok := arg.(NumberValue); ok {

		builder = fmt.Sprintf("%v", num.Value)

	} else if boolean, ok := arg.(BooleanValue); ok {

		builder = fmt.Sprintf("%v", boolean.Value)

	} else if str, ok := arg.(StringValue); ok {

		builder = fmt.Sprintf("%v", str.Value)

	} else if arr, ok := arg.(ArrayValue); ok {

		builder = "["
		for i := 0; i < len(arr.Value); i++ {

			val := fmt.Sprintf("%v", printHelper(arr.Value[i]))

			builder += val

			if i < len(arr.Value)-1 {
				builder += ", "
			}
		}
		builder += "]"

	} else if map_, ok := arg.(MapValue); ok {

		builder = "{"
		counter := 0

		for key, value := range map_.Value {

			s := fmt.Sprintf("%v : %v", printHelper(key), printHelper(value))

			builder += s

			if counter < (len(map_.Value) - 1) {
				builder += ", "
			}

			counter++
		}

		builder += "}"

	} else if null, ok := arg.(NullValue); ok {

		builder = fmt.Sprintf("%v", null.Value)

	} else if obj, ok := arg.(ObjectVal); ok {

		numArgs := len(obj.Properties)
		counter := 0

		for name, arg := range obj.Properties {
			counter++

			if counter < numArgs {
				builder += fmt.Sprintf("%v: %v, ", name, printHelper(arg))
			} else {
				builder += fmt.Sprintf("%v: %v", name, printHelper(arg))
			}
		}
	}

	return builder
}
