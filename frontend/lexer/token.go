package lexer

import "fmt"

type Token struct {
	Type  TokenType
	Value string
}

type Tokens []Token

func (t Token) ToString() string {
	return fmt.Sprintf("Type: '%v', Value: '%v'", t.Type, t.Value)
}

func (t Tokens) ToString() string {

	builder := ""
	for _, token := range t {
		builder += token.ToString()
	}

	return builder
}

type TokenType string

const (
	Number     TokenType = "Number"
	Identifier TokenType = "Identifier"
	Boolean    TokenType = "Boolean"
	String     TokenType = "String"

	// Symbols.
	Equals       TokenType = "Equals"
	OpenParen    TokenType = "OpenParen"
	CloseParen   TokenType = "CloseParen"
	Comma        TokenType = "Comma"
	Colon        TokenType = "Colon"
	OpenBrace    TokenType = "OpenBrace"
	CloseBrace   TokenType = "CloseBrace"
	OpenBracket  TokenType = "OpenBracket"
	CloseBracket TokenType = "CloseBracket"
	Period       TokenType = "Period"
	Equality     TokenType = "Equality"
	Ternary      TokenType = "Ternary"

	// Operators.
	BinaryOperator      TokenType = "BinaryOperator"
	ConditionalOperator TokenType = "ConditionalOperator"

	// Keywords.
	Let   TokenType = "Let"   // declaring new variables
	Const TokenType = "Const" // declaring new constants
	Fn    TokenType = "Fn"    // declaring new functions
	If    TokenType = "If"    // standard if condition
	Else  TokenType = "Else"  // standard else condition
	While TokenType = "While"

	// End of Line.
	EOL TokenType = "EOL"
	// End of file.
	EOF TokenType = "EOF"
)

// Map of languages keywords.
var Keywords = map[string]TokenType{
	"let":   Let,
	"const": Const,
	"fn":    Fn,
	"if":    If,
	"else":  Else,
	"while": While,
}
