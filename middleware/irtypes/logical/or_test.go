package logical

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Or_Expression_Invalid_Type_LHS_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	o := Or{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := o.Validate(&context)

	want := "type error: or: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Or_Expression_Invalid_Type_LHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	o := Or{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := o.Validate(&context)

	want := "type error: or: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Or_Expression_Invalid_Type_RHS_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	o := Or{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := o.Validate(&context)

	want := "type error: or: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Or_Expression_Invalid_Type_RHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	o := Or{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := o.Validate(&context)

	want := "type error: or: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
