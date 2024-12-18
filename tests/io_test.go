package tests

import (
	"fmt"
	"os"
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
		{`using "io";
		io.print();`, "interpreter error: unexpected number of args for io.print, expected 1 got 0", true},
		{`using "io";
		io.print("hello", "world");`, "interpreter error: unexpected number of args for io.print, expected 1 got 2", true},
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
		{`using "io";
		io.println();`, "interpreter error: unexpected number of args for io.println, expected 1 got 0", true},
		{`using "io";
		io.println("hello", "world");`, "interpreter error: unexpected number of args for io.println, expected 1 got 2", true},
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
		{`using "io";
		io.printf();`, "interpreter error: unexpected number of args for io.printf, expected min 1 got 0", true},
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
		{`using "io";
		let j = io.sprintf();`, "interpreter error: unexpected number of args for io.sprintf, expected min 1 got 0", true},
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

func TestInput(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		stdin       string
		want        string
		throwsError bool
	}{
		{`using "io";
		let i = io.input("Message: ");
		io.print(i);`, "Test message\n", "Message: Test message", false},
		{`using "io";
		let j = io.input();
		io.print(j);`, "\n", "interpreter error: unexpected number of args for io.input, expected 1 got 0", true},
		{`using "io";
		let k = io.input("Message: ", "Another message");
		io.print(k);`, "\n", "interpreter error: unexpected number of args for io.input, expected 1 got 2", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Create a pipe to simulate stdin
			reader, writer, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			defer reader.Close()

			// Replace os.Stdin with the reader end of the pipe
			env.Stdin = reader

			// Write the test input to the writer end of the pipe
			go func() {
				writer.Write([]byte(tt.stdin))
				writer.Close()
			}()

			// Run the program.
			_, err = program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf("%v - %v", err.Error(), output.String())
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

func TestOpen(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let f = io.open("test.txt", "r");
		io.print(f);`, "../source/test.txt", false},
		{`using "io";
		let f = io.open("doesNotExist.txt", "r");
		io.print(f);`, "interpreter error: open ../source/doesNotExist.txt: no such file or directory", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf("%v - %v", err.Error(), output.String())
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

func TestReadLine(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let f = io.open("test.txt", "r");
		let line = io.readline(f, 1);
		io.print(line);`, "Hello, World!", false},
		{`using "io";
		let fr = io.open("test.txt", "r");
		let liner = io.readline(fr, 3);
		io.print(liner);`, "interpreter error: line number 3 not found in ../source/test.txt", true},
		{`using "io";
		let fw = io.open("test.txt", "w");
		let linew = io.readline(fw, 1);
		io.print(linew);`, "interpreter error: file: ../source/test.txt not opened in a valid read-mode", true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf("%v - %v", err.Error(), output.String())
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

func TestReadLines(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf("%v - %v", err.Error(), output.String())
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

func TestClose(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		let f = io.open("test.txt", "r");

		let line = io.readline(f, 1);
		io.println(line);
		io.close(f);

		let anotherLine = io.readline(f, 1);`, "interpreter error: read ../source/test.txt: file already closed", true},
		{`using "io";
		let file = io.open("test.txt", "r");

		let lline = io.readline(file, 1);
		io.println(lline);`, "Hello, World!\n", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.source, tt.want)
		t.Run(testname, func(t *testing.T) {

			// Run the program.
			_, err := program.Run(string(tt.source), env)

			if !tt.throwsError {

				// When tests aren't supposed to throw an error.
				if err != nil {
					t.Errorf("%v - %v", err.Error(), output.String())
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
