package tests

import (
	"os"
	"testing"

	"goblin.org/main/program"
)

// Tests a basic `if` condition.
func TestMap(t *testing.T) {

	// Setup the program env.
	HarnessSetup()

	file := "../source/map_test.gob"
	expected := "10\n30\n"

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
