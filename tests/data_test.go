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
		{`using "data";
		using "io";
		let aa = [];
		data.push(aa, 5, 2);
		io.print(aa);`, "interpreter error: unexpected number of args for data.push, expected 2 got 3", true},
		{`using "data";
		using "io";
		let aaa = [];
		data.push(aaa, 5, 2, 3);
		io.print(aaa);`, "interpreter error: unexpected number of args for data.push, expected 2 got 4", true},
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

func TestPut(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "data";
		using "io";
		let map = {};
		data.put(map, "one", 1);
		let size = data.size(map);
		io.print(size);`, "1", false},
		{`using "data";
		using "io";
		let m = {};
		data.put(m, "one", 1);
		io.print(m);`, "{one : 1}", false},
		{`using "data";
		using "io";
		let mapp = {
			"one": 1,
		};
		data.put(mapp, "two", 2);
		io.print(mapp);`, "{one : 1, two : 2}", false},
		{`using "data";
		using "io";
		let mm = {
			"one": 1,
		};
		data.put(mm, "two");
		io.print(mm);`, "interpreter error: unexpected number of args for data.put, expected 3 got 2", true},
		{`using "data";
		using "io";
		let mmm = {
			"one": 1,
		};
		data.put(mmm, "two", 1, 2);
		io.print(mmm);`, "interpreter error: unexpected number of args for data.put, expected 3 got 4", true},
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

func TestPop(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "data";
		using "io";

		let arr = [1, 2, 3];
		let last = data.pop(arr);
		
		io.println(last);
		io.println(arr);`, "3\n[1, 2]\n", false},
		{`using "data";
		using "io";

		let arrr = [];
		let lastt = data.pop(arrr);
		
		io.println(lastt);
		io.println(arrr);`, "interpreter error: cannot pop an empty array", true},
		{`using "data";
		using "io";

		let arrrr = [1, 2, 3];
		let lasttt = data.pop(arrrr, 1);
		
		io.println(lasttt);
		io.println(arrrr);`, "interpreter error: unexpected number of args for data.pop, expected 1 got 2", true},
		{`using "data";
		using "io";

		let arrrrr = [];
		let lastttt = data.pop();
		
		io.println(lastttt);
		io.println(arrrrr);`, "interpreter error: unexpected number of args for data.pop, expected 1 got 0", true},
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
		{`using "data";
		using "io";
		let aa = [1, 2, 3, 4, 5];
		let sArr = data.size(aa, 1);
		io.print(sArr);`, "interpreter error: unexpected number of args for data.size, expected 1 got 2", true},
		{`using "data";
		using "io";
		let aaaarr = [];
		let sArr = data.size();
		io.print(sArr);`, "interpreter error: unexpected number of args for data.size, expected 1 got 0", true},
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
