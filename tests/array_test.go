package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestSimpleArray(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let arr = [1, 2, 3, 4, 5];
		io.println(arr[2]);`, "3\n", false},
		{`using "io";
		let arrr = [1, 2, 3, 4, 5];
		io.println(arrr[6]);`, "interpreter error: index out of bounds for index 6", true},
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
