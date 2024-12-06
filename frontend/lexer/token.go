package lexer

import "fmt"

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
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
	Equals       TokenType = "="
	OpenParen    TokenType = "("
	CloseParen   TokenType = ")"
	Comma        TokenType = ","
	Colon        TokenType = ":"
	OpenBrace    TokenType = "{"
	CloseBrace   TokenType = "}"
	OpenBracket  TokenType = "["
	CloseBracket TokenType = "]"
	Period       TokenType = "."
	Equality     TokenType = "=="
	Ternary      TokenType = "?"

	// Operators.
	BinaryOperator      TokenType = "BinaryOperator"      // e.g. '+, -, /, *, etc'
	ConditionalOperator TokenType = "ConditionalOperator" // e.g. '==, <=, >=, etc'
	ShorthandOperator   TokenType = "ShorthandOperator"   // e.g. '++, --'

	// Keywords.
	Let   TokenType = "Let"   // declaring new variables
	Const TokenType = "Const" // declaring new constants
	Fn    TokenType = "Fn"    // declaring new functions
	If    TokenType = "If"    // standard if condition
	Else  TokenType = "Else"  // standard else condition
	While TokenType = "While" // standard while loop
	For   TokenType = "For"   // standard for loop
	Using TokenType = "Using"

	// End of Line.
	EOL TokenType = ";"
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
	"for":   For,
	"using": Using,
}
