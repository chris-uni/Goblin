package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Gte struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Gte) String() string {
	return fmt.Sprintf("gte %v %v %v", v.Destination, v.Lhs, v.Rhs)
}

func (Gte) Exec(context *i.IRContext) {}
