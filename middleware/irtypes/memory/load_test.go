package memory

import (
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Load_Invalid_Source(t *testing.T) {

	context := i.IRContext{}

	irTemp := context.AllocateTemporary()

	load := Load{
		Destination: irTemp,
		Source:      i.IRAddress{},
	}

	out := load.Validate(&context)
	got := out.Error()

	want := "undefined storage address @0\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}

func Test_Load_Invalid_Destination(t *testing.T) {

	context := i.IRContext{}

	irAddr := context.AllocateAddress()

	load := Load{
		Destination: i.IRTemporary{},
		Source:      irAddr,
	}

	out := load.Validate(&context)
	got := out.Error()

	want := "unrecognised type <nil>\n"

	if got != want {
		t.Errorf("\ngot %v\nwant %v\n", got, want)
	}
}
