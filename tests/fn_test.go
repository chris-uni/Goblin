package tests

import (
	"os"
	"testing"

	"goblin.org/main/program"
)

// Tests a basic `if` condition.
func TestFnCall(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	file := "../source/fn_test.gob"
	expected := "Hello, World\n"

	source, err := os.ReadFile(file)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Run the program.
	_, err = program.Run(string(source), env)
	if err != nil {
		t.Errorf(err.Error())
	}

	if expected != output.String() {
		t.Errorf("expected `%v`, received `%v`", expected, output.String())
	}
}
