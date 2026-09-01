package controlflow

import (
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_JmpIf_Condition_Invalid_Type_String(t *testing.T) {

	context := i.IRContext{}
	irLabel := context.AllocateLabel()

	jmpif := JmpIf{
		Destination: irLabel,
		Condition:   i.IRString{Value: "Hello"},
	}

	out := jmpif.Validate(&context)
	got := out.Error()

	want := "type error: jmpif: condition of invalid type\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_JmpIf_Condition_Invalid_Type_Number(t *testing.T) {

	context := i.IRContext{}
	irLabel := context.AllocateLabel()

	jmpif := JmpIf{
		Destination: irLabel,
		Condition:   i.IRNumber{Value: 10},
	}

	out := jmpif.Validate(&context)
	got := out.Error()

	want := "type error: jmpif: condition of invalid type\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
