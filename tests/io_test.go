package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestIO(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		io.println("Hello, World");`, "Hello, World\n", false},
		{`using "io";
		io.print("Hello, World");`, "Hello, World", false},
		{`using "io";
		io.print(12);`, "12", false},
		{`using "io";
		let arr = [1, 2, 3];
		io.print(arr);`, "[1, 2, 3]", false},
		{`using "io";
		let map = {
			"one": 1,
			"two": 2,
			"three": 3,
		};
		io.print(map);`, "{'one': 1, 'two': 2, 'three': 3}", false},
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
