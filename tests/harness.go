package tests

import (
	"bytes"
	"os"

	"goblin.org/main/runtime"
)

// Mock std::out buffer using the below byte buffer.
var output bytes.Buffer

var env = runtime.Environment{}

// Set up the test harness environment.
func HarnessSetup() {
	output = bytes.Buffer{}
	env.Stdout = &output
	env.Stdin = os.Stdin
	env.Variables = map[string]runtime.RuntimeValue{}
	env.Constants = map[string]bool{}
	env.Namespaces = map[string]runtime.Namespace{}
	env.EntryLocation = "../source"

	env.Setup()
}

func FlushBuffer() {
	output.Reset()
}
