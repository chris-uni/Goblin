package logical

import i "goblin.org/main/middleware/irtypes"

type And struct {
	Destination i.IRTemporary
	Val1        i.IRValue
	Val2        i.IRValue
}

func (And) Exec(context *i.IRContext) {}
