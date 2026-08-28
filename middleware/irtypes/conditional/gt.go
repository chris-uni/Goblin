package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Gt struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Gt) String() string {
	return fmt.Sprintf("gt %v %v %v", v.Destination, v.Lhs, v.Rhs)
}
func (Gt) Exec(context *i.IRContext) {}
