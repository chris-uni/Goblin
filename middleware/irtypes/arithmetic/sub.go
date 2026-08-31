package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Sub struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Sub) String() string            { return fmt.Sprintf("sub %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Sub) Exec(context *i.IRContext) {}
func (v *Sub) Left() i.IRValue           { return v.Lhs }
func (v *Sub) Right() i.IRValue          { return v.Rhs }
func (v *Sub) Dest() i.IRTemporary       { return v.Destination }
