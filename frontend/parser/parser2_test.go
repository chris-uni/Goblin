package parser

import (
	"fmt"
	"testing"

	"goblin.org/main/frontend/lexer"
)

func Test_Primary_NumberExpression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.String, Value: "Hello, World!", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Boolean, Value: "true", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Boolean, Value: "false", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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
	tokens = append(tokens, lexer.Token{Type: lexer.OpenParen, Value: "(", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 1})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.CloseParen, Value: ")", Line: 1, Col: 8})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Let, Value: "let", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Equals, Value: "=", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 8})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 3})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Equals, Value: "=", Line: 1, Col: 2})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "20", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 7})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "-", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 8})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "/", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 8})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "*", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 8})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "%", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 8})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "-", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "-", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "2", Line: 1, Col: 9})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 10})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 11})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "/", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "/", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "2", Line: 1, Col: 9})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 10})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 11})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "*", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "2", Line: 1, Col: 9})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 10})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 11})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "*", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "-", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "2", Line: 1, Col: 9})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 10})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 11})

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
	tokens = append(tokens, lexer.Token{Type: lexer.OpenParen, Value: "(", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 1})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.CloseParen, Value: ")", Line: 1, Col: 7})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "*", Line: 1, Col: 9})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "2", Line: 1, Col: 11})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 12})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 13})

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
	tokens = append(tokens, lexer.Token{Type: lexer.Let, Value: "let", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Equals, Value: "=", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 8})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 9})

	_, err := ParseTokens(tokens)

	want := "expecting token `Identifier`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_2Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Let, Value: "let", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 8})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 9})

	_, err := ParseTokens(tokens)

	want := "expecting token `=`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_3Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Equals, Value: "=", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 7})

	_, err := ParseTokens(tokens)

	want := "expecting lhs of assignment to be identifier type, got {NumericLiteralNode 10}"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_4Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 5})

	_, err := ParseTokens(tokens)

	want := "unexpected token found during parsing ';'"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_5Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Let, Value: "let", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Identifier, Value: "x", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.Equals, Value: "=", Line: 1, Col: 6})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 4})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 5})

	_, err := ParseTokens(tokens)

	want := "unexpected token found during parsing ';'"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_6Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.OpenParen, Value: "(", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 1})
	tokens = append(tokens, lexer.Token{Type: lexer.EOL, Value: ";", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 4})

	_, err := ParseTokens(tokens)

	want := "expecting token `)`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Invalid_7Expression(t *testing.T) {

	tokens := make([]lexer.Token, 0)
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "10", Line: 1, Col: 0})
	tokens = append(tokens, lexer.Token{Type: lexer.BinaryOperator, Value: "+", Line: 1, Col: 3})
	tokens = append(tokens, lexer.Token{Type: lexer.Number, Value: "5", Line: 1, Col: 5})
	tokens = append(tokens, lexer.Token{Type: lexer.EOF, Value: string(lexer.EOF), Line: 1, Col: 4})

	_, err := ParseTokens(tokens)

	want := "expecting token `;`"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
