package controlflow

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type JmpIf struct {
	Destination i.IRLabel
	Condition   i.IRValue
}

func (v *JmpIf) String() string {
	return fmt.Sprintf("jmpif %v %v", v.Condition, v.Destination)
}

func (JmpIf) Exec(context *i.IRContext) {}
