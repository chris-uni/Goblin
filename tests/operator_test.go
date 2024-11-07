package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestOperators(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source string
		want   string
	}{
		{`let x = 1;
		x++;
		println(x);`, "2\n"},
		{`let y = 1;
		y--;
		println(y);`, "0\n"},
		{`let j = 1;
		j += 1;
		println(j);`, "2\n"},
		{`let k = 1;
		k -= 1;
		println(k);`, "0\n"},
		{`let i = 1;
		i *= 2;
		println(i);`, "2\n"},
		{`let n = 1;
		n /= 1;
		println(n);`, "1\n"},
		{`let m = 1;
		m %= 1;
		println(m);`, "0\n"},
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
