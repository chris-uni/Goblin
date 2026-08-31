package logical

import i "goblin.org/main/middleware/irtypes"

type Not struct {
	Destination i.IRTemporary
	Val         i.IRValue
}

func (Not) Exec(context *i.IRContext) {}
