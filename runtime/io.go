package runtime

import (
	"bufio"
	"fmt"
	"os"
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
	- when opening files, file locations should be relative to the main.gob file, not the Goblin interpreter.
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

	numArgs := len(args)
	if numArgs != 2 {
		return nil, fmt.Errorf("unexpected number of args for io.open, expected 2 got %v", numArgs)
	}

	fp, isStr := args[0].(StringValue)
	if !isStr {
		return nil, fmt.Errorf("file path: string expected, got %v", fp)
	}

	m, isStr := args[1].(StringValue)
	if !isStr {
		return nil, fmt.Errorf("file mode: string expected, got %v", m)
	}

	mode := fileOpenFlags(m.Value)
	filePath := env.EntryLocation + "/" + fp.Value

	file, err := os.OpenFile(filePath, mode, 0644) // Adjust permissions as needed
	if err != nil {
		return nil, err
	}

	initCursorValue := 1
	return FileObjectValue{
		Path:          filePath,
		File:          file,
		Mode:          mode,
		IsOpen:        true,
		CursorPointer: &initCursorValue,
	}, nil
}

// close - closes the specified file object.
// io.close(fileObject *fileObj)
var close FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for io.close, expected 1 got %v", numArgs)
	}

	f := args[0]
	fileObj, isFileObj := f.(FileObjectValue)
	if !isFileObj {
		return nil, fmt.Errorf("io.close expectes arg1 to be of type fileObject, %v given", fileObj.Type)
	}

	// Close the underlying file.
	if err := fileObj.File.Close(); err != nil {
		return nil, err
	}

	fileObj.IsOpen = false

	return MK_NULL(), nil
}

// read - reads a single line from the specified file.
// io.readline(fileObject *fileObj, lineNumber int) string
var readline FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 2 {
		return nil, fmt.Errorf("unexpected number of args for io.readline, expected 2 got %v", numArgs)
	}

	f := args[0]
	fileObj, isFileObj := f.(FileObjectValue)
	if !isFileObj {
		return nil, fmt.Errorf("io.readline expectes arg1 to be of type fileObject, %v given", fileObj.Type)
	}

	l := args[1]
	line, isStr := l.(NumberValue)
	if !isStr {
		return nil, fmt.Errorf("io.readline expectes arg2 to be of type int, %v given", line.Type)
	}

	// Only allow this to work if file opened in appropriate mode.
	if fileObj.IsOpen && (fileObj.Mode == os.O_RDONLY || fileObj.Mode == os.O_RDWR) {

		val, err := fileReader(fileObj, line.Value)
		if err != nil {
			return nil, err
		}

		return MK_STRING(val), nil
	}

	// File was opened in a non-read mode.
	return nil, fmt.Errorf("file: %v not opened in a valid read-mode", fileObj.Path)
}

// readline - reads a file line by line, uses internal file pointer to continue pointing to next
// line. Returns \r\n when line limit is reached.
// io.readlines(fileObject *fileObj) string
var readlines FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	numArgs := len(args)
	if numArgs != 1 {
		return nil, fmt.Errorf("unexpected number of args for io.readlines, expected 1 got %v", numArgs)
	}

	f := args[0]
	fileObj, isFileObj := f.(FileObjectValue)
	if !isFileObj {
		return nil, fmt.Errorf("io.readlines expectes arg1 to be of type fileObject, %v given", fileObj.Type)
	}

	// Only allow this to work if file opened in appropriate mode.
	if fileObj.IsOpen && (fileObj.Mode == os.O_RDONLY || fileObj.Mode == os.O_RDWR) {

		val, err := fileReader(fileObj, *fileObj.CursorPointer)
		if err != nil {
			return nil, err
		}

		*fileObj.CursorPointer++

		return MK_STRING(val), nil
	}

	return nil, nil
}

// write - writes the contents of the buffer to the specified fileObject.
// io.write(fileObject *fileObj, buffer []byte)
var write FunctionCall = func(args []RuntimeValue, env Environment) (RuntimeValue, error) {

	return nil, nil
}

func fileReader(file FileObjectValue, lineNumber int) (string, error) {

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file.File)

	// Iterate through lines until the desired line number is reached
	currentLine := 1
	for scanner.Scan() {
		if currentLine == lineNumber {
			return scanner.Text(), nil
		}
		currentLine++
	}

	// Handle errors during scanning and cases where the line number is out of bounds
	if err := scanner.Err(); err != nil {
		return "\r\n", err
	}

	return "\r\n", nil
}

// Helper function for open
func fileOpenFlags(mode string) int {

	var flags int

	if strings.Contains(mode, "r") {
		flags = os.O_RDONLY
	}
	if strings.Contains(mode, "w") {
		flags = os.O_WRONLY
	}
	if strings.Contains(mode, "+") {
		flags = os.O_RDWR
	}

	return flags
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
	} else if fileObj, ok := arg.(FileObjectValue); ok {

		builder += fileObj.Path
	}

	return builder
}
