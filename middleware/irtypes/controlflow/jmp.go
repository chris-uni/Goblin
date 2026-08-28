package controlflow

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Jmp struct {
	Destination i.IRLabel
}

func (v *Jmp) String() string {
	return fmt.Sprintf("jmp %v", v.Destination)
}

func (Jmp) Exec(context *i.IRContext) {}
