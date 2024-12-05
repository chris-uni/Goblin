package utils

import (
	"fmt"
	"io"
	"strconv"
)

// Removes the first element from an array and returns that removed element.
func Shift[T any](slice *[]T) T {

	if len(*slice) == 0 {
		var zero T
		return zero
	}
	firstElement := (*slice)[0]
	*slice = (*slice)[1:]

	return firstElement
}

// Converts a string into a number.
func ToNumber(str string) (int, error) {

	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("could not convert `%v` to int", str)
	}

	return num, nil
}

// Checks to see if item 'value' is in slice 'slice'.
func ContainsString(slice []string, value string) bool {

	for _, item := range slice {

		fmt.Printf("Checking %v against %v\n", item, value)
		if item == value {
			return true
		}
	}
	return false
}

// Converts a boolean into a string.
func BtoS(b bool) string {
	if b {
		return "true"
	}

	return "false"
}

// Converts a string into a boolean.
func StoB(s string) bool {
	if s == "true" {
		return true
	}

	return false
}

// Language standard output writer.
func Stdout(s string, buff io.Writer) {

	fmt.Fprintf(buff, "%v", s)
}

// Formats an error string to help with debugging.
/*
i.e.
parse error: using io
             ~~~~~~~~^
expecting 'EOL' on line 1 col 10
*/
func GenerateParserError(auditLine string, specificToken string, line int, col int, message string) string {

	origin := "parse error:"
	underlines := ""

	for i := 0; i < len(origin)+1; i++ {
		underlines += " "
	}

	for i := 0; i < len(auditLine)+1; i++ {

		if i == (col + (len(specificToken))) {
			underlines += "^"
		} else {
			underlines += "~"
		}
	}

	msg := fmt.Sprintf("%v\n%v\n%v on line %v col %v", auditLine, underlines, message, line, col)

	return msg
}
