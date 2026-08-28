package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Mul struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Mul) String() string            { return fmt.Sprintf("mul %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Mul) Exec(context *i.IRContext) {}
func (v *Mul) Left() i.IRValue           { return v.Lhs }
func (v *Mul) Right() i.IRValue          { return v.Rhs }
func (v *Mul) Dest() i.IRTemporary       { return v.Destination }
