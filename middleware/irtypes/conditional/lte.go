package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Lte struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Lte) String() string {
	return fmt.Sprintf("lte %v %v %v", v.Destination, v.Lhs, v.Rhs)
}
func (Lte) Exec(context *i.IRContext) {}
