package conditional

import (
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Eq_Expression_Incompatiable_String_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := eq.Validate(&context)
	got := out.Error()

	want := "type error: eq: incompatible types\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Eq_Expression_Incompatiable_String_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := eq.Validate(&context)
	got := out.Error()

	want := "type error: eq: incompatible types\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Eq_Expression_Incompatiable_Number_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := eq.Validate(&context)
	got := out.Error()

	want := "type error: eq: incompatible types\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Eq_Expression_Success_Number(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRNumber{Value: 10},
		Rhs:         i.IRNumber{Value: 10},
	}

	out := eq.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}

func Test_Eq_Expression_Success_Boolean(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRBoolean{Value: true},
		Rhs:         i.IRBoolean{Value: true},
	}

	out := eq.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}

func Test_Eq_Expression_Success_String(t *testing.T) {

	context := i.IRContext{}
	irTemp := context.AllocateTemporary()

	eq := Eq{
		Destination: irTemp,
		Lhs:         i.IRString{Value: "Hello"},
		Rhs:         i.IRString{Value: "Hello"},
	}

	out := eq.Validate(&context)

	if out != nil {
		t.Errorf("\ngot %v\nwant %v\n", out, "<nil>")
	}
}
