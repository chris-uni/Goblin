package main

import (
	"fmt"
	"strings"

	"goblin.org/main/Tokens"
)

func Tokenize(source string) []Tokens.Token {

	tokens := make([]Tokens.Token, 0)
	elements := strings.Split(source, "")

	for len(elements) > 0 {

		// Handles single char tokens.
		if elements[0] == "(" {
			tokens = append(tokens, Tokens.NewToken(Shift(&elements), Tokens.OpenParen))
		} else if elements[0] == ")" {
			tokens = append(tokens, Tokens.NewToken(Shift(&elements), Tokens.CloseParen))
		} else if elements[0] == "+" || elements[0] == "-" || elements[0] == "*" || elements[0] == "/" {
			tokens = append(tokens, Tokens.NewToken(Shift(&elements), Tokens.BinaryOperator))
		} else if elements[0] == "=" {
			tokens = append(tokens, Tokens.NewToken(Shift(&elements), Tokens.Equals))
		} else if elements[0] == ";" {
			tokens = append(tokens, Tokens.NewToken(Shift(&elements), Tokens.EOL))
		} else {

			// Handle multi-char token, i.e. <=

			// Build a number token...
			if IsInt(elements[0]) {

				number := ""

				// While there are still values left and they are numbers.
				for len(elements) > 0 && IsInt(elements[0]) {

					number += Shift(&elements)
				}

				tokens = append(tokens, Tokens.NewToken(number, Tokens.Number))

			} else if IsAlpha(elements[0]) {

				// Build an identifier token...
				value := ""

				// While there are still values left and they are numbers.
				for len(elements) > 0 && IsAlpha(elements[0]) {

					value += Shift(&elements)
				}

				// If the value just built was a not a reserved word, then it was an identifier.
				reserved, ok := Tokens.IsReserved(value)
				if !ok {
					tokens = append(tokens, Tokens.NewToken(value, Tokens.Identifier))
				} else {

					// Else, it was a reserved word and we should use it.
					tokens = append(tokens, Tokens.NewToken(value, reserved))
				}
			} else if IsSkippable(elements[0]) {
				// Skips the current element.
				Shift(&elements)
			} else {
				e := fmt.Sprintf("Unrecognised character element in source: %v \n", elements[0])
				panic(e)
			}
		}
	}

	return tokens
}
