package tests

import (
	"os"
	"testing"

	"goblin.org/main/program"
)

// Tests a basic `if` condition.
func TestWhile(t *testing.T) {

	// Setup the program env.
	env.Setup()

	// Empty the test stdout buffer.
	flushBuffer(&output)

	file := "../source/while_test.gob"
	expected := "0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n"

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
