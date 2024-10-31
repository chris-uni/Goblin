package tests

import (
	"bytes"
	"os"
	"testing"

	"goblin.org/main/program"
	"goblin.org/main/runtime"
)

var output bytes.Buffer

var env = runtime.Environment{
	Stdout:    &output,
	Variables: map[string]runtime.RuntimeValue{},
	Constants: map[string]bool{},
}

// Tests a basic `if` condition.
func TestIfCondition(t *testing.T) {

	// Setup the program env.
	env.Setup()

	file := "../source/if_test.gob"
	expected := "2 is bigger than 1\n"

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
