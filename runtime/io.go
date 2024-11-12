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
		"printf": {
			Type: "NativeFn",
			Call: printf,
		},
		"sprintf": {
			Type: "NativeFn",
			Call: sprintf,
		},
	},
}

// Defines the built-in method 'print'.
// Depending on arg value type, print is handled in different ways.
var print FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	utils.Stdout(builder, env.Stdout)

	return MK_NULL(), nil
}

// Defines the built-in method 'println'.
// Same as 'print' but adds a '\n' char at the end of the output.
var println FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	builder += "\n"

	utils.Stdout(builder, env.Stdout)

	return MK_NULL(), nil
}

// Defines build-in method 'printf'.
// Printf allows for formatted statements to be printed.
var printf FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	s, err := printer(args)
	if err != nil {
		return nil, err
	}

	// Write to std-out.
	utils.Stdout(s, env.Stdout)

	return MK_NULL(), nil
}

// Defines build-in method 'printf'.
// Printf allows for formatted statements to be printed.
var sprintf FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	s, err := printer(args)
	if err != nil {
		return nil, err
	}

	return MK_STRING(s), nil
}

// Helper function for printf, sprintf
func printer(args []RuntimeValue) (string, error) {

	formattedString, isStr := args[0].(StringValue)
	arguments := args[1:]
	builder := ""

	if !isStr {
		return "", fmt.Errorf("string type required for formatted string, got %v", formattedString)
	}
	for i := 0; i < len(formattedString.Value); i++ {
		// If we encounter a '%' character
		if formattedString.Value[i] == '%' {
			// Check if we have enough arguments
			if i+1 < len(formattedString.Value) && len(arguments) > 0 {
				// Switch on the format specifier
				switch formattedString.Value[i+1] {
				case 'd': // Integer
					iVal, err := utils.ToNumber(printHelper(arguments[0]))
					if err != nil {
						return "", err
					}

					builder += fmt.Sprintf("%d", iVal)
					arguments = arguments[1:]
					i++
				case 's': // String
					builder += fmt.Sprint(printHelper(arguments[0]))
					arguments = arguments[1:]
					i++
				case 'v': // Default of the type specified.
					builder += fmt.Sprintf("%v", printHelper(arguments[0]))
					arguments = arguments[1:]
					i++
				// Add more cases for other format specifiers as needed
				default: // If the format specifier is not recognized, print it literally
					builder += fmt.Sprintf("%c", formattedString.Value[i])
				}
			} else { // If there are not enough arguments, print '%' literally
				builder += fmt.Sprintf("%%")
			}
		} else { // If the current character is not '%', print it literally
			builder += fmt.Sprintf("%c", formattedString.Value[i])
		}
	}

	return builder, nil
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
