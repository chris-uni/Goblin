package lexer

import (
	"fmt"
	"strings"
	"unicode"

	"goblin.org/main/utils"
)

func Tokenize(sourceCode string) ([]Token, map[int]string) {

	tokens := make(Tokens, 0)
	audit := make(map[int]string)
	auditBuilder := ""

	src := strings.Split(sourceCode, "")

	line := 1
	col := 0

	for len(src) > 0 {

		if src[0] == "\n" {
			// Capture the line we just lexed.
			audit[line] = auditBuilder

			// Reset the builder
			auditBuilder = ""

			// Increment the line count.
			line++

			// Reset the col counter.
			col = 0
		}
		if src[0] == " " {
			auditBuilder += src[0]
			col++
		}

		if src[0] == "(" {
			auditBuilder += src[0]
			tokens = append(tokens, token(OpenParen, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == ")" {
			auditBuilder += src[0]
			tokens = append(tokens, token(CloseParen, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "{" {
			auditBuilder += src[0]
			tokens = append(tokens, token(OpenBrace, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "}" {
			auditBuilder += src[0]
			tokens = append(tokens, token(CloseBrace, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "[" {
			auditBuilder += src[0]
			tokens = append(tokens, token(OpenBracket, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "]" {
			auditBuilder += src[0]
			tokens = append(tokens, token(CloseBracket, utils.Shift[string](&src), line, col))
			col++

		} else if src[0] == "+" {

			if (src[1] == "+") || (src[1] == "=") {
				// Shorthand ++ or +=
				auditBuilder += src[0] + src[1]
				op := fmt.Sprintf("%v%v", utils.Shift[string](&src), utils.Shift[string](&src))
				tokens = append(tokens, token(ShorthandOperator, op, line, col))
				col += 2
			} else {
				// Standard + BinOp.
				auditBuilder += src[0]
				tokens = append(tokens, token(BinaryOperator, utils.Shift[string](&src), line, col))
				col++
			}
		} else if src[0] == "-" {

			if (src[1] == "-") || (src[1] == "=") {
				// Shorthand -- or -=
				auditBuilder += src[0] + src[1]
				op := fmt.Sprintf("%v%v", utils.Shift[string](&src), utils.Shift[string](&src))
				tokens = append(tokens, token(ShorthandOperator, op, line, col))
				col += 2
			} else {
				// Standard - BinOp.
				auditBuilder += src[0]
				tokens = append(tokens, token(BinaryOperator, utils.Shift[string](&src), line, col))
				col++
			}
		} else if src[0] == "/" || src[0] == "*" || src[0] == "%" {

			// Shorthand operator or standard BinaryOperator?
			if src[1] == "=" {
				// Shorthand operator.
				auditBuilder += src[0] + src[1]
				op := fmt.Sprintf("%v%v", utils.Shift[string](&src), utils.Shift[string](&src))
				tokens = append(tokens, token(ShorthandOperator, op, line, col))
				col += 2
			} else {
				// Standard BinOp.
				auditBuilder += src[0]
				tokens = append(tokens, token(BinaryOperator, utils.Shift[string](&src), line, col))
				col++
			}
		} else if src[0] == ">" || src[0] == "<" {
			auditBuilder += src[0]
			tokens = append(tokens, token(ConditionalOperator, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "=" && src[1] != "=" {
			auditBuilder += src[0]
			tokens = append(tokens, token(Equals, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == ";" {
			auditBuilder += src[0]
			tokens = append(tokens, token(EOL, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == ":" {
			auditBuilder += src[0]
			tokens = append(tokens, token(Colon, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "," {
			auditBuilder += src[0]
			tokens = append(tokens, token(Comma, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "." {
			auditBuilder += src[0]
			tokens = append(tokens, token(Period, utils.Shift[string](&src), line, col))
			col++
		} else if src[0] == "?" {
			auditBuilder += src[0]
			tokens = append(tokens, token(Ternary, utils.Shift[string](&src), line, col))
			col++
		} else {

			// Multicharacter tokens (<=, >=...)

			if src[0] == "=" && src[1] == "=" {

				auditBuilder += src[0] + src[1]

				// This is an '==' operator.
				symbol := utils.Shift[string](&src)
				symbol += utils.Shift[string](&src)

				tokens = append(tokens, token(Equality, symbol, line, col))
				col += 2

			} else if isInt(src[0]) {
				// Builds a number token.
				num := ""

				for len(src) > 0 && isInt(src[0]) {
					num += utils.Shift[string](&src)
				}

				tokens = append(tokens, token(Number, num, line, col))
				auditBuilder += num
				col += len(num)

			} else if isQuote(src[0]) {

				// Start of a string literal.
				str := ""

				// Shift past '"'.
				utils.Shift[string](&src)

				for len(src) > 0 && !isQuote(src[0]) {
					c := utils.Shift[string](&src)
					str += c
				}

				// Shift past '"'.
				utils.Shift[string](&src)

				tokens = append(tokens, token(String, str, line, col))
				auditBuilder += str
				col += (len(str) + 2) // Lenght of string + 2 for the quotes either side.

			} else if isAlpha(src[0]) {
				// Builds an identifier token.
				iden := ""

				for len(src) > 0 && isAlpha(src[0]) {
					iden += utils.Shift[string](&src)
				}

				// Check for reserved keyword.
				t, ok := Keywords[iden]
				if !ok {

					// If not exist, check to see is this is a bool value.
					bVal, err := truthValue(iden)
					if err == nil {
						bsv := utils.BtoS(bVal)
						tokens = append(tokens, token(Boolean, bsv, line, col))
						auditBuilder += bsv
						col += len(bsv)
					} else {

						// Really is an identifier.
						tokens = append(tokens, token(Identifier, iden, line, col))
						auditBuilder += iden
						col += len(iden)
					}
				} else {
					tokens = append(tokens, token(t, iden, line, col))
					auditBuilder += iden
					col += len(iden)
				}

			} else if isSkippable(src[0]) {
				// Skips to next character.
				utils.Shift(&src)
			} else {
				fmt.Printf("Unrecognised token in source: %v \n", src[0])
				panic("Program exit.")
			}
		}
	}

	// Add in the EOF token.
	tokens = append(tokens, token(EOF, "EOF", line, col))

	// Add the final of the lexer audit.
	audit[line] = auditBuilder

	return tokens, audit
}

// Checks to see if we are starting a new string.
func isQuote(src string) bool {

	q := []rune(src)
	return q[0] == '"'
}

// Checks to see if the src[0] contains alpha characters only.
func isAlpha(src string) bool {

	r := []rune(src)

	return unicode.IsLetter(r[0])
}

// Checks to see if the src[0] is numeric.
func isInt(src string) bool {

	r := []rune(src)

	return unicode.IsDigit(r[0])
}

// Determined that this is a bool value, so getting its value.
func truthValue(src string) (bool, error) {

	if src == "true" {
		return true, nil
	} else if src == "false" {
		return false, nil
	}

	// Dont need an error message here, as not being a boolean value means
	// this is actually an identifier
	return false, fmt.Errorf("")
}

// Checks to see if this is a skippable token.
func isSkippable(src string) bool {
	return src == " " || src == "\n" || src == "\t" || src == "\r"
}

// Builds and returns a new token.
func token(tknType TokenType, value string, line int, col int) Token {

	token := Token{
		Type:  tknType,
		Value: value,
		Line:  line,
		Col:   col,
	}

	return token
}
