package conditional

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Gt_Expression_Invalid_Type_LHS_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	gt := Gt{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := gt.Validate(&context)

	want := "type error: gt: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_Type_LHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	gt := Gt{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := gt.Validate(&context)

	want := "type error: gt: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_Type_RHS_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	gt := Gt{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := gt.Validate(&context)

	want := "type error: gt: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Invalid_Type_RHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	gt := Gt{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := gt.Validate(&context)

	want := "type error: gt: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Gt_Expression_Success_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	gt := Gt{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := gt.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}
