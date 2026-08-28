package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Lt struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Lt) String() string {
	return fmt.Sprintf("lt %v %v %v", v.Destination, v.Lhs, v.Rhs)
}

func (Lt) Exec(context *i.IRContext) {}
