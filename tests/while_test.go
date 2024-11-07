package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestWhile(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source string
		want   string
	}{
		{`let i = 0;
		while (i < 5) {
			println(i);
			i++;
		}`, "0\n1\n2\n3\n4\n"},
		{`let j = 5;
		while (j > 0) {
			println(j);
			j--;
		}`, "5\n4\n3\n2\n1\n"},
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
