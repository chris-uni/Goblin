package parser

import (
	"fmt"
	"testing"

	"goblin.org/main/frontend/lexer"
)

func generateLexerTokens(source string) ([]lexer.Token, error) {

	tokens, _, err := lexer.Lex(source)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func Test_Primary_NumberExpression(t *testing.T) {

	tokens, err := generateLexerTokens(`10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{NumericLiteralNode 10}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Primary_StringExpression(t *testing.T) {

	tokens, err := generateLexerTokens(`"Hello, World!";`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{StringLiteralNode Hello, World!}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Primary_BooleanTExpression(t *testing.T) {

	tokens, err := generateLexerTokens(`true;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BooleanLiteralNode true}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Primary_BooleanFExpression(t *testing.T) {

	tokens, err := generateLexerTokens(`false;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BooleanLiteralNode false}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Primary_IdentifierExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`x;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{IdentifierNode x}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Primary_GroupedExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`(10 + 10);`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} +}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_VariableDeclerationExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`let x = 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{VariableDeclarationNode {NumericLiteralNode 10} false x}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_AssignmentExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`x = 20;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{AssignmentExprNode {NumericLiteralNode 20} {IdentifierNode x}}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_AdditionExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 + 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 8})

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} +}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_SubtractionExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 - 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} -}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_DivisionExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 / 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} /}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_MultiplicationExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 * 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} *}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_ModulusExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 % 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 10} %}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_Associativity_1Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 - 5 - 2;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf("\n%v\n", err)
	}

	want := "{Program [{BinaryExprNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 5} -} {NumericLiteralNode 2} -}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_Associativity_2Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 / 5 / 2;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf("\n%v\n", err)
	}

	want := "{Program [{BinaryExprNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 5} /} {NumericLiteralNode 2} /}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_Precedence_1Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 + 5 * 2;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf("\n%v\n", err)
	}

	want := "{Program [{BinaryExprNode {NumericLiteralNode 10} {BinaryExprNode {NumericLiteralNode 5} {NumericLiteralNode 2} *} +}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_Precedence_2Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 * 5 - 2;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf("\n%v\n", err)
	}

	want := "{Program [{BinaryExprNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 5} *} {NumericLiteralNode 2} -}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Artehmetic_GroupingExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`(10 + 5) * 2;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := ParseTokens(tokens)
	if err != nil {
		t.Errorf("\n%v\n", err)
	}

	want := "{Program [{BinaryExprNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 5} +} {NumericLiteralNode 2} *}]}"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_1Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`let = 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "expecting token `Identifier`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_2Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`let x 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "expecting token `=`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_3Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 = x;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "expecting lhs of assignment to be identifier type, got {NumericLiteralNode 10}"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_4Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 +;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "unexpected token found during parsing ';'"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_5Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`let x = ;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "unexpected token found during parsing ';'"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_6Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`(10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "expecting token `)`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_7Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens, err := generateLexerTokens(`10 +5`)
	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = ParseTokens(tokens)

	want := "expecting token `;`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_If_Simple_Expression(t *testing.T) {

	var tests = []struct {
		in   string
		want string
	}{

		{"if(10 < 4){}", "{Program [{IfNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 4} <} [] false []}]}"},
		{`if (10 < 4){
				let y = 100;
			}`, "{Program [{IfNode {BinaryExprNode {NumericLiteralNode 10} {NumericLiteralNode 4} <} [{VariableDeclarationNode {NumericLiteralNode 100} false y}] false []}]}"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.in)
		t.Run(testname, func(t *testing.T) {

			tkns, err := generateLexerTokens(tt.in)
			prog, err := ParseTokens(tkns)
			out := fmt.Sprintf("%v", prog)

			if err != nil {
				t.Errorf(err.Error())
			}

			if out != tt.want {
				t.Errorf("\ngot %v\nwant %v", out, tt.want)
			}
		})
	}
}
