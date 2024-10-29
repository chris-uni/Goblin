package Tokens

import "strings"

type TokenType string

var (
	Number         TokenType = "Number"
	Identifier     TokenType = "Identifier"
	Equals         TokenType = "Equals"
	OpenParen      TokenType = "OpenParen"
	CloseParen     TokenType = "CloseParen"
	BinaryOperator TokenType = "BinaryOperator"
	Let            TokenType = "Let"
	EOL            TokenType = "EOL"

	ReservedWords = map[string]TokenType{
		"let": Let,
	}
)

/*
Determines if the provided input is a reserved word.
*/
func IsReserved(input string) (TokenType, bool) {
	c, ok := ReservedWords[strings.ToLower(input)]

	return c, ok
}

type Token struct {
	Value string
	Type  TokenType
}

/*
Returns a new Token.
*/
func NewToken(_value string, _type TokenType) Token {

	return Token{
		Value: _value,
		Type:  _type,
	}
}
