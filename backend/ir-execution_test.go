package backend

import (
	"fmt"
	"testing"

	"goblin.org/main/frontend/lexer"
	"goblin.org/main/frontend/parser"
	"goblin.org/main/middleware"
	i "goblin.org/main/middleware/irtypes"
)

func assembleValidatedIR(source string) ([]i.IRCommand, error) {

	// Stage 1. Lex the input.
	tokens, _, err := lexer.Lex(source)
	if err != nil {
		return nil, err
	}

	// Stage 2. Produce the Abstract Syntax Tree.
	program, err := parser.ParseTokens(tokens)
	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err.Error())
	}

	// Stage 3. Reduce to GoblinIR (reducer, validator, optimiser [tbc]).
	goblinIR, err := middleware.OrchestrateIRLayer(program)
	if err != nil {
		return nil, fmt.Errorf("ir error: %v\n", err.Error())
	}

	return goblinIR, nil
}

func Test_Simple_Numeric_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{10}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_String_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return "hello";`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{hello}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Boolean_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return false;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{false}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Lt_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 < 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{false}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Lte_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 <= 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{false}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Gt_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 > 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{true}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Gte_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 >= 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{true}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Eq_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 == 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{false}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Simple_Neq_Return(t *testing.T) {

	validatedIR, err := assembleValidatedIR(`return 10 != 5;`)
	if err != nil {
		t.Errorf(err.Error())
	}

	out, err := Execution(validatedIR)

	want := `{true}`
	got := fmt.Sprintf("%v", out)

	if want != got {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
