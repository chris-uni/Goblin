package logical

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Neg_Expression_Invalid_Type_LHS_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	n := Neg{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: false},
	}

	out := n.Validate(&context)

	want := "type error: neg: operand of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Neg_Expression_Invalid_Type_LHS_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	n := Neg{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
	}

	out := n.Validate(&context)

	want := "type error: neg: operand of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
