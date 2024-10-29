package main

import (
	"strconv"
	"unicode/utf8"
)

/*
Returns the first element of the slice and shifts all values up by 1.
*/
func Shift[T any](slice *[]T) T {

	if len(*slice) == 0 {
		var zero T
		return zero
	}

	first := (*slice)[0]
	*slice = (*slice)[1:]

	return first
}

/*
Returns true if the provided char is a letter, regardless of case.
*/
func IsAlpha(input string) bool {

	src, _ := utf8.DecodeRuneInString(input)

	if src >= 'A' && src <= 'Z' || src >= 'a' && src <= 'z' {
		return true
	}

	return false
}

/*
Returns true if the provided char is a integer.
*/
func IsInt(input string) bool {

	if _, err := strconv.Atoi(input); err == nil {
		return true
	}

	return false
}

/*
Is this a character element we can skip over?
*/
func IsSkippable(input string) bool {

	if input == " " || input == "\n" || input == "\t" {
		return true
	}

	return false
}
