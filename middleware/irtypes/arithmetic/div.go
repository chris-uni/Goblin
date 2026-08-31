package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Div struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (v *Div) String() string            { return fmt.Sprintf("div %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Div) Exec(context *i.IRContext) {}
func (v *Div) Left() i.IRValue           { return v.Lhs }
func (v *Div) Right() i.IRValue          { return v.Rhs }
func (v *Div) Dest() i.IRTemporary       { return v.Destination }
