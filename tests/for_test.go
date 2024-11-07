package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestForLoop(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`for(let i = 0; i < 5; i++;){
			println(i);
		}`, "0\n1\n2\n3\n4\n", false},
		{`for(let i = 5; i > 0; i--;){
			println(i);
		}`, "5\n4\n3\n2\n1\n", false},
		{`let arr = [1, 2, 3, 4, 5];
		for(let i = 0; i < 5; i++;){
			println(arr[i]);
		}`, "1\n2\n3\n4\n5\n", false},
		{`let arrr = ["foo", "bar", "foobar"];
		for(let i = 0; i < 3; i++;){
			let val = arrr[i];
			println(val);
		}`, "foo\nbar\nfoobar\n", false},
		{`let arrrr = ["foo", "bar", "foobar"];
		let map = {
			"foo": 10,
			"bar": 20,
			"foobar": 30,
		};

		for(let i = 0; i < 3; i++;){
			let key = arrrr[i];
			let val = map[key];
			println(val);
		}`, "10\n20\n30\n", false},
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
