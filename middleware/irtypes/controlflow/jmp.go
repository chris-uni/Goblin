package controlflow

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Jmp struct {
	Destination i.IRLabel
}

func (j *Jmp) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	fmt.Printf("jmp to %v\n", j.Destination.PCOffset)

	state.PC = j.Destination.PCOffset

	return nil, nil
}

func (j *Jmp) Validate(context *i.IRContext) error {

	_, err := i.ResolveIRType(j.Destination, context)
	if err != nil {
		return err
	}

	return nil
}

func (j *Jmp) String() string {
	return fmt.Sprintf("jmp %v", j.Destination)
}
