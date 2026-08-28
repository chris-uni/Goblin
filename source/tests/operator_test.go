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
		{`using "io";
		let x = 1;
		x++;
		io.println(x);`, "2\n"},
		{`using "io";
		let y = 1;
		y--;
		io.println(y);`, "0\n"},
		{`using "io";
		let j = 1;
		j += 1;
		io.println(j);`, "2\n"},
		{`using "io";
		let k = 1;
		k -= 1;
		io.println(k);`, "0\n"},
		{`using "io";
		let i = 1;
		i *= 2;
		io.println(i);`, "2\n"},
		{`using "io";
		let n = 1;
		n /= 1;
		io.println(n);`, "1\n"},
		{`using "io";
		let m = 1;
		m %= 1;
		io.println(m);`, "0\n"},
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
