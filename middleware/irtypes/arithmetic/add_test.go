package arithmetic

import (
	"fmt"
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Add_Expression_Invalid_Type_LHS(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	add := Add{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := add.Validate(&context)

	want := "type error: add: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Add_Expression_Invalid_Type_RHS(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	add := Add{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := add.Validate(&context)

	want := "type error: add: operands of invalid type\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Add_Expression_Incompatiable_Types(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	add := Add{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := add.Validate(&context)

	want := "type error: add: incompatible types\n"
	got := fmt.Sprintf("%v", out)

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Add_Expression_Success_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	add := Add{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := add.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}

func Test_Add_Expression_Success_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	add := Add{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := add.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}
