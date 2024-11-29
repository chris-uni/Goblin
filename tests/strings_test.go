package tests

import (
	"fmt"
	"testing"

	"goblin.org/main/program"
)

func TestSplit(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	var tests = []struct {
		source      string
		want        string
		throwsError bool
	}{
		{`using "io";
		using "strings";

		let words = strings.split("Hello world", " ");
		io.print(words[0]);`, "Hello", false},
		{`using "io";
		using "data";
		using "strings";

		let wordss = strings.split("Hello world", " ");
		let count = data.size(wordss);

		io.print(count);`, "2", false},
		{`using "io";
		using "data";
		using "strings";

		let source = "Hello world this is a test";
		let wordsss = strings.split(source, " ");
		let size = data.size(wordsss);

		for(let i = 0; i < size; i++;){

			io.print(wordsss[i]);
		}`, "Helloworldthisisatest", false},
		{`using "io";
		using "strings";

		let words = strings.split("Hello world");
		io.print(words[0]);`, "interpreter error: unexpected number of args for strings.split, expected 2 got 1", true},
		{`using "io";
		using "strings";

		let words = strings.split("Hello world", ",", "");
		io.print(words[0]);`, "interpreter error: unexpected number of args for strings.split, expected 2 got 3", true},
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
