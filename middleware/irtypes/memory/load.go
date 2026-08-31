package memory

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Load struct {
	Destination i.IRTemporary
	Source      i.IRAddress
}

func (v *Load) String() string {
	return fmt.Sprintf("load %v %v", v.Destination, v.Source)
}

func (Load) Exec(context *i.IRContext) {}
