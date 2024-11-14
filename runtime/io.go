package runtime

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"goblin.org/main/utils"
)

/*
TODO:
- Arg count check, add error handling to check that all functions have the correct number of arguments specified, and that they are of the proper type.
- Fix bug in open that occurs when specifying local file (files are local to interpreter/main.go, not local to the .gob source file)
- Add new type: fileObject
	- contains a reference to the specified file
	- value that specifies the mode of the file (read, write, append)
- Add new type: byte
	- Similar to string, but is a collection of bytes
*/

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
		"input": {
			Type: "NativeFn",
			Call: input,
		},
		"readline": {
			Type: "NativeFn",
			Call: readline,
		},
		"readlines": {
			Type: "NativeFn",
			Call: readlines,
		},
		"open": {
			Type: "NativeFn",
			Call: open,
		},
		"close": {
			Type: "NativeFn",
			Call: close,
		},
		"write": {
			Type: "NativeFn",
			Call: write,
		},
	},
}

// print, a standard printing function.
// io.print(msg string)
var print FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for io.print, expected 1 got %v", numArgs)
	}

	str, err := printer(args)
	if err != nil {
		return nil, err
	}

	utils.Stdout(str, env.Stdout)

	return MK_NULL(), nil
}

// println - acts the same as print, but appends a new line to the end.
// io.println(msg string)
var println FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for io.println, expected 1 got %v", numArgs)
	}

	str, err := printer(args)
	if err != nil {
		return nil, err
	}

	utils.Stdout(str+"\n", env.Stdout)

	return MK_NULL(), nil
}

// printf - allows for formatted statements to be printed.
// io.printf(formattedString string, args ...any)
var printf FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs < 1 {
		return nil, fmt.Errorf("unexpected number of args for io.printf, expected min 1 got %v", numArgs)
	}

	s, err := printerFormatter(args)
	if err != nil {
		return nil, err
	}

	// Write to std-out.
	utils.Stdout(s, env.Stdout)

	return MK_NULL(), nil
}

// sprintf - allows for formatted statements to be printed.
// io.sprintf(formattedString string, args ...any) string
var sprintf FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs < 1 {
		return nil, fmt.Errorf("unexpected number of args for io.sprintf, expected min 1 got %v", numArgs)
	}

	s, err := printerFormatter(args)
	if err != nil {
		return nil, err
	}

	return MK_STRING(s), nil
}

// input - reads a single line from std::in.
// io.input(message string) string
var input FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for io.input, expected 1 got %v", numArgs)
	}

	m := args[0]
	msg, isStr := m.(StringValue)
	if !isStr {
		return nil, fmt.Errorf("os.input message must be string type, got %v", msg.Type)
	}

	reader := bufio.NewReader(env.Stdin)
	utils.Stdout(msg.Value, env.Stdout)

	input, err := reader.ReadString('\n') // Read until EOF (Ctrl+D)
	if err != nil {
		return nil, err
	}

	// Remove the EOF character and any trailing newline
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)

	return MK_STRING(input), nil
}

// open - returns a new file object using the specified mode, i.e. r, w, a.
// io.open(fileName string, mode string) fileObject
var open FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	return nil, nil
}

// close - closes the specified file object.
// io.close(fileObject *fileObj)
var close FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	return nil, nil
}

// read - reads a single line from the specified file.
// io.readline(fileObject *fileObj, lineNumber int) string
var readline FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 2 {
		return nil, fmt.Errorf("unexpected number of args for io.read, expected 2 got %v", numArgs)
	}

	f := args[0]
	file, isStr := f.(StringValue)
	if !isStr {
		return nil, fmt.Errorf("io.read expectes arg1 to be of type string, %v given", file.Type)
	}

	l := args[1]
	line, isStr := l.(NumberValue)
	if !isStr {
		return nil, fmt.Errorf("io.read expectes arg2 to be of type int, %v given", line.Type)
	}

	lineNumber := line.Value

	// Get the absolute path to the file
	absPath, err := filepath.Abs(file.Value)
	if err != nil {
		return nil, err
	}

	// Open the file for reading
	fileObj, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer fileObj.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(fileObj)

	// Iterate through lines until the desired line number is reached
	currentLine := 1
	for scanner.Scan() {
		if currentLine == lineNumber {
			return MK_STRING(scanner.Text()), nil
		}
		currentLine++
	}

	// Handle errors during scanning and cases where the line number is out of bounds
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("line number %d not found in file", lineNumber)
}

// readline - reads a file line by line
// io.readlines(fileObject *fileObj) []string
var readlines FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	return nil, nil
}

// writen - writes the contents of the buffer to the specified fileObject.
// io.write(fileObject *fileObj, buffer []byte)
var write FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	return nil, nil
}

// Helper function for printf, sprintf.
func printerFormatter(args []RuntimeValue) (string, error) {

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

// Helper function for print, println.
func printer(args []RuntimeValue) (string, error) {

	builder := ""

	for _, arg := range args {

		builder += printHelper(arg)
	}

	return builder, nil
}

// Helper function, resolves types to strings.
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
