package runtime

import (
	"fmt"
)

// Defines the built-in method 'print'.
// Depending on arg value type, print is handled in different ways.
var Print FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	fmt.Fprintf(env.Stdout, "%v\n", builder)

	return MK_NULL()
}

// Defines the built-in method 'println'.
// Same as 'print' but adds a '\n' char at the end of the output.
var Println FunctionCall = func(args []RuntimeValue, env Environment) RuntimeValue {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	fmt.Fprintf(env.Stdout, "%v\n", builder)

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

			val := fmt.Sprintf("%v", arr.Value[i])

			builder += val

			if i < len(arr.Value)-1 {
				builder += ", "
			}
		}
		builder += "]"

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
