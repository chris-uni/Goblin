package memory

import (
	"testing"

	i "goblin.org/main/middleware/irtypes"
)

func Test_Store_Invalid_Destination(t *testing.T) {

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
