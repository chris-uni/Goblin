package lexer

import (
	"fmt"
	"strings"

	tokens "goblin.org/main/Tokens"
	tools "goblin.org/main/Tools"
)

func Tokenize(source string) []tokens.Token {

	new_tokens := make([]tokens.Token, 0)
	elements := strings.Split(source, "")

	for len(elements) > 0 {

		// Handles single char tokens.
		if elements[0] == "(" {
			new_tokens = append(new_tokens, tokens.NewToken(tools.Shift(&elements), tokens.OpenParen))
		} else if elements[0] == ")" {
			new_tokens = append(new_tokens, tokens.NewToken(tools.Shift(&elements), tokens.CloseParen))
		} else if elements[0] == "+" || elements[0] == "-" || elements[0] == "*" || elements[0] == "/" {
			new_tokens = append(new_tokens, tokens.NewToken(tools.Shift(&elements), tokens.BinaryOperator))
		} else if elements[0] == "=" {
			new_tokens = append(new_tokens, tokens.NewToken(tools.Shift(&elements), tokens.Equals))
		} else if elements[0] == ";" {
			new_tokens = append(new_tokens, tokens.NewToken(tools.Shift(&elements), tokens.EOL))
		} else {

			// Handle multi-char token, i.e. <=

			// Build a number token...
			if tools.IsInt(elements[0]) {

				number := ""

				// While there are still values left and they are numbers.
				for len(elements) > 0 && tools.IsInt(elements[0]) {

					number += tools.Shift(&elements)
				}

				new_tokens = append(new_tokens, tokens.NewToken(number, tokens.Number))

			} else if tools.IsAlpha(elements[0]) {

				// Build an identifier token...
				value := ""

				// While there are still values left and they are numbers.
				for len(elements) > 0 && tools.IsAlpha(elements[0]) {

					value += tools.Shift(&elements)
				}

				// If the value just built was a not a reserved word, then it was an identifier.
				reserved, ok := tokens.IsReserved(value)
				if !ok {
					new_tokens = append(new_tokens, tokens.NewToken(value, tokens.Identifier))
				} else {

					// Else, it was a reserved word and we should use it.
					new_tokens = append(new_tokens, tokens.NewToken(value, reserved))
				}
			} else if tools.IsSkippable(elements[0]) {
				// Skips the current element.
				tools.Shift(&elements)
			} else {
				e := fmt.Sprintf("Unrecognised character element in source: %v \n", elements[0])
				panic(e)
			}
		}
	}

	new_tokens = append(new_tokens, tokens.NewToken("EOF", tokens.EOF))

	return new_tokens
}
