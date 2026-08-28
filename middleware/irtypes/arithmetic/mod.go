package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Mod struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Mod) String() string            { return fmt.Sprintf("mod %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Mod) Exec(context *i.IRContext) {}
func (v *Mod) Left() i.IRValue           { return v.Lhs }
func (v *Mod) Right() i.IRValue          { return v.Rhs }
func (v *Mod) Dest() i.IRTemporary       { return v.Destination }
