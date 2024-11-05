package tests

import (
	"bytes"

	"goblin.org/main/runtime"
)

// Mock up the stdout buffer using the below byte buffer.
var output bytes.Buffer

var env = runtime.Environment{}

// Set up the test harness environment.
func HarnessSetup() {
	output = bytes.Buffer{}
	env.Stdout = &output
	env.Variables = map[string]runtime.RuntimeValue{}
	env.Constants = map[string]bool{}

	env.Setup()
}
