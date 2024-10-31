package tests

import (
	"bytes"

	"goblin.org/main/runtime"
)

// Mock up the stdout buffer using the below byte buffer.
var output bytes.Buffer

var env = runtime.Environment{
	Stdout:    &output,
	Variables: map[string]runtime.RuntimeValue{},
	Constants: map[string]bool{},
}

// Empties the current harness buffer.
func flushBuffer(o *bytes.Buffer) {
	*o = bytes.Buffer{}
}
