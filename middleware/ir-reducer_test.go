package middleware

import (
	"fmt"
	"testing"

	"goblin.org/main/frontend/ast"
	"goblin.org/main/frontend/lexer"
	"goblin.org/main/frontend/parser"
)

func assembleAST(source string) (ast.Program, error) {

	tokens, _, err := lexer.Lex(source)
	if err != nil {
		return ast.Program{}, err
	}

	program, err := parser.ParseTokens(tokens)
	if err != nil {
		return ast.Program{}, err
	}

	return program, nil
}

/*
Arithemetic Tests
*/

func Test_Add_Expression_Numeric(t *testing.T) {

	ast, err := assembleAST(`10 + 10;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[add %0 10 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Add_Expression_String(t *testing.T) {

	ast, err := assembleAST(`"Hello, " + " world!";`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[add %0 Hello,   world!]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Sub_Expression_Numeric(t *testing.T) {

	ast, err := assembleAST(`10 - 10;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[sub %0 10 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Numeric(t *testing.T) {

	ast, err := assembleAST(`10 * 10;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[mul %0 10 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Div_Expression_Numeric(t *testing.T) {

	ast, err := assembleAST(`10 / 10;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[div %0 10 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mod_Expression_Numeric(t *testing.T) {

	ast, err := assembleAST(`10 % 10;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[mod %0 10 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

/*
Store Tests
*/
func Test_StoreValue_Single(t *testing.T) {

	ast, err := assembleAST(`let x = 10;`)
	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[store @0 10]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_StoreValue_Multiple(t *testing.T) {

	ast, err := assembleAST(`
	let x = 10;
	let y = 20;
	`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[store @0 10 store @1 20]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

/*
Load Tests
*/
func Test_LoadValue_Simple(t *testing.T) {

	ast, err := assembleAST(`
	let x = 10;
	let y = x;
	`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[store @0 10 load %0 @0 store @1 %0]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_LoadValue_Multiple(t *testing.T) {

	ast, err := assembleAST(`
	let x = 10;
	let y = 20;
	let z = x;
	let w = y;
	`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[store @0 10 store @1 20 load %0 @0 store @2 %0 load %1 @1 store @3 %1]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

/*
Assignment Tests.
*/

func Test_Assignment_Expression_Valid(t *testing.T) {

	ast, err := assembleAST(`
	let x = 10;
	x = 1;`)

	if err != nil {
		t.Error(err.Error())
	}

	out, err := Reduce(ast)
	if err != nil {
		t.Errorf(err.Error())
	}

	want := "[store @0 10 store @0 1]"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Assignment_Expression_Invalid(t *testing.T) {

	ast, err := assembleAST(`x = 1;`)

	if err != nil {
		t.Error(err.Error())
	}

	_, err = Reduce(ast)

	want := "reducer error: undefined symbol x\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
