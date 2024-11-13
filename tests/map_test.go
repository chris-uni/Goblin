package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestSimpleMap(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let map = {
			"foo": 10,
			"bar": 20,
		};
		io.println(map["foo"]);`, "10\n", false},
		{`using "io";
		let mapp = {
			"foo": 10,
			"bar": 20,
		};
		io.println(mapp["baz"]);`, "interpreter error: key `{String baz}` does not exist for map: mapp", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf(err.Error())
				}

				if output.String() != tt.want {
					t.Errorf("expected `%v`, received `%v`", tt.want, output.String())
				}
			} else {

				// When tests are supposed to throw an error.
				if err.Error() != tt.want {
					t.Errorf("expected `%v`, received `%v`", tt.want, err.Error())
				}
			}

			FlushBuffer()
		})
	}
}
