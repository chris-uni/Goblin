package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestPush(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "data";
		using "io";
		let arr = [];
		data.push(arr, 1);
		io.print(arr);`, "[1]", false},
		{`using "data";
		using "io";
		let arrr = [1, 2, 3, 4];
		data.push(arrr, 5);
		io.print(arrr);`, "[1, 2, 3, 4, 5]", false},
		{`using "data";
		using "io";
		let arrrr = [];
		data.push(arrrr, 5);
		io.print(arrrr[0]);`, "5", false},
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
