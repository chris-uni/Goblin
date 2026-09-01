package arithmetic

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Mul_Expression_Invalid_Type_LHS_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	mul := Mul{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := mul.Validate(&context)

	want := "type error: mul: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Invalid_Type_LHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	mul := Mul{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := mul.Validate(&context)

	want := "type error: mul: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Invalid_Type_RHS_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	mul := Mul{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := mul.Validate(&context)

	want := "type error: mul: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Invalid_Type_RHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	mul := Mul{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := mul.Validate(&context)

	want := "type error: mul: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Mul_Expression_Success_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	mul := Mul{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := mul.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}
