package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Neq struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (s *Neq) String() string {
	return fmt.Sprintf("neq %v %v %v", s.Destination, s.Lhs, s.Rhs)
}

func (Neq) Exec(context *i.IRContext) {}
