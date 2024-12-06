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
		{`using "io";
		if (10 < 20){
			io.print(11);
		}`, "11"},
		{`using "io";
		if (10 > 5){
			io.print(10);
		}`, "10"},
		{`using "io";
		if (10 == 10){
			io.println(10);
		}`, "10\n"},
		{`using "io";
		if (5 > 10){
			io.println(5);
		}
		else {
			io.println(10);
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

			fmt.Printf("Output buffer: %v\n", env.Stdout)
			FlushBuffer()
		})
	}
}
