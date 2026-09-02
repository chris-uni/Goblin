package logical

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Not_Expression_Invalid_Type_LHS_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	n := Not{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
	}

	out := n.Validate(&context)

	want := "type error: not: operand of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Not_Expression_Invalid_Type_LHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	n := Not{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
	}

	out := n.Validate(&context)

	want := "type error: not: operand of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
