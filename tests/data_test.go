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

func TestSize(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "data";
		using "io";
		let arr = [1, 2, 3, 4, 5];
		let sArr = data.size(arr);
		io.print(sArr);`, "5", false},
		{`using "data";
		using "io";
		let map = {
			"foo": 10,
			"bar": 20,
			"baz": 30,
		};
		let sMap = data.size(map);
		io.print(sMap);`, "3", false},
		{`using "data";
		using "io";
		let arrr = [];
		let sArrr = data.size(arrr);
		io.print(sArrr);`, "0", false},
		{`using "data";
		using "io";
		let arrrr = [];
		data.push(arrr, 1);
		data.push(arrr, 2);
		data.push(arrr, 3);
		let sArrrr = data.size(arrr);
		io.print(sArrrr);`, "3", false},
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
