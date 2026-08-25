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

/*
Comparision Operator Tests
*/

func Test_Eq_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "true" == "false";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: eq: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Eq_Expression_Invalid_Multi(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = true == "false";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: eq: incompatible types\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Neq_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "true" != "false";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: eq: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Neq_Expression_Invalid_Multi(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = true != "false";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: eq: incompatible types\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lt_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" < "world";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lt_Expression_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = false < true;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lt_Expression_Invalid_Multi(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" < 10;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lte_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" <= "world";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lte: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lte_Expression_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = false <= true;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lte: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Lte_Expression_Invalid_Multi(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" <= 10;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: lte: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" > "world";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: gt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = false > true;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: gt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_Multi(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" > 10;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: gt: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gte_Expression_Invalid_String(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = "hello" >= "world";`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: gte: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gte_Expression_Invalid_Boolean(t *testing.T) {

	rawIR, err := assembleRawIR(`let x = false >= true;`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: gte: operands of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

/*
JmpIf Tests
*/

func Test_JmpIf_Expression_String(t *testing.T) {

	rawIR, err := assembleRawIR(`if("string"){}`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: jmpif: condition of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_JmpIf_Expression_Number(t *testing.T) {

	rawIR, err := assembleRawIR(`if(10){}`)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Validate(rawIR)

	want := "validation error: type error: jmpif: condition of invalid type\n\n"
	got := fmt.Sprintf("%v", err)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
