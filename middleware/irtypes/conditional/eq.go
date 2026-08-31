package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Eq struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (s *Eq) String() string {
	return fmt.Sprintf("eq %v %v %v", s.Destination, s.Lhs, s.Rhs)
}

func (Eq) Exec(context *i.IRContext) {}
