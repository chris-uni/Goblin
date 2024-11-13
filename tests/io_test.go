package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestPrint(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		io.print("Hello, World");`, "Hello, World", false},
		{`using "io";
		io.print(12);`, "12", false},
		{`using "io";
		let arr = [1, 2, 3];
		io.print(arr);`, "[1, 2, 3]", false},
		{`using "io";
		fn printer(){
			io.print("Hello");
		}
		printer();`, "Hello", false},
		{`using "io";
		fn anotherPrinter(var){
			io.print(var);
		}
		anotherPrinter("Hello");`, "Hello", false},
		{`using "io";
		let keys = ["one", "two", "three"];
		let map = {
			"one": 1,
			"two": 2,
			"three": 3,
		};

		for(let i = 0; i < 3; i++;){

			io.print(map[keys[i]]);
		}`, "123", false},
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

func TestPrintln(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		io.println("Hello, World");`, "Hello, World\n", false},
		{`using "io";
		io.println(12);`, "12\n", false},
		{`using "io";
		let arr = [1, 2, 3];
		io.println(arr);`, "[1, 2, 3]\n", false},
		{`using "io";
		fn printer(){
			io.println("Hello");
		}
		printer();`, "Hello\n", false},
		{`using "io";
		fn anotherPrinter(var){
			io.println(var);
		}
		anotherPrinter("Hello");`, "Hello\n", false},
		{`using "io";
		let keys = ["one", "two", "three"];
		let map = {
			"one": 1,
			"two": 2,
			"three": 3,
		};

		for(let i = 0; i < 3; i++;){

			io.println(map[keys[i]]);
		}`, "1\n2\n3\n", false},
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

func TestPrintf(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		io.printf("Hello, %v", "World");`, "Hello, World", false},
		{`using "io";
		let arr = [1, 2, 3];
		io.printf("One: %v", arr[0]);`, "One: 1", false},
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

func TestSPrintf(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let i = io.sprintf("Hello, %v", "World");
		io.print(i);`, "Hello, World", false},
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
