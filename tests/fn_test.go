package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestFunctions(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`fn Print() {
			println("Hello, World");
		}
		Print();`, "Hello, World\n", false},
		{`fn testPrint(a, b){
			let x = a + b;
			println(x);
		}
		testPrint(1, 3);`, "4\n", false},
		{`fn Adder(a, b, c){
			let x = a + b;
			println(x);
		}
		Adder(1, 3);`, "interpreter error: incorrect number of params specified for fn Adder, got 2 want 3", true},
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
