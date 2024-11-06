package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestSimpleIfCondition(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source string
		want   string
	}{
		{`if (10 > 5){
			println(10);
		}
		`, "10\n"},
		{`if (10 < 20){
			println(11);
		}`, "11\n"},
		{`if (10 == 10){
			println(10);
		}`, "10\n"},
		{`if (5 > 10){
			println(5);
		}
		else {
			println(10);
		}`, "10\n"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)
			if err != nil {
				t.Errorf(err.Error())
			}

			if output.String() != tt.want {
				t.Errorf("expected `%v`, received `%v`", tt.want, output.String())
			}

			FlushBuffer()
		})
	}
}
