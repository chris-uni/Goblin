package middleware

import (
	"fmt"
	"testing"
)

func assembleRawIR(source string) ([]IRCommand, error) {

	ast, err := assembleAST(source)

	if err != nil {
		return nil, err
	}

	rawIR, err := Reduce(ast)
	if err != nil {
		return nil, err
	}

	return rawIR, nil
}

/*
Arithemetic Tests
*/

func Test_Add_Expression_Numeric_Invalid_MultiType(t *testing.T) {

	rawIR, err := assembleRawIR(`10 + "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: add: incompatible types\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Add_Expression_Numeric_Invalid_BooleanType(t *testing.T) {

	rawIR, err := assembleRawIR(`true + true;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: add: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Sub_Expression_Numeric_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`"hello" - "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: sub: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Sub_Expression_Numeric_Invalid_MultiType(t *testing.T) {

	rawIR, err := assembleRawIR(`10 - "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: sub: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Sub_Expression_Numeric_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`true - false;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: sub: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Numeric_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`"hello" * "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mul: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Numeric_Invalid_MultiType(t *testing.T) {

	rawIR, err := assembleRawIR(`10 * "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mul: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Numeric_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`true * false;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mul: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Div_Expression_Numeric_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`"hello" / "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: div: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Div_Expression_Numeric_Invalid_MultiType(t *testing.T) {

	rawIR, err := assembleRawIR(`10 / "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: div: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Div_Expression_Numeric_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`true / false;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: div: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mod_Expression_Numeric_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`"hello" % "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mod: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mod_Expression_Numeric_Invalid_MultiType(t *testing.T) {

	rawIR, err := assembleRawIR(`10 % "hello";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mod: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mod_Expression_Numeric_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`true % false;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: mod: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
