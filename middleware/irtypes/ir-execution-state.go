package irtypes

type IRExecutionState struct {
	Storage     []IRValue
	Temporaries []IRValue
	Labels      []IRLabel
	PC          int
}

type IRExecutionResult struct {
	Value IRValue
}
