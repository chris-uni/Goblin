package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Add struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Add) String() string            { return fmt.Sprintf("add %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Add) Exec(context *i.IRContext) {}
func (v *Add) Left() i.IRValue           { return v.Lhs }
func (v *Add) Right() i.IRValue          { return v.Rhs }
func (v *Add) Dest() i.IRTemporary       { return v.Destination }
