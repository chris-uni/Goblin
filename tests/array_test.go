package tests

import (
	"os"
	"testing"

	"goblin.org/main/program"
)

// Tests a basic `if` condition.
func TestArray(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	file := "../source/array_test.gob"
	expected := "10\n34\n56\n12\n78\n"

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
