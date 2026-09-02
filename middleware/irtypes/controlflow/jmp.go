package controlflow

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Jmp struct {
	Destination i.IRLabel
}

func (j *Jmp) Exec(context *i.IRContext) {}

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
