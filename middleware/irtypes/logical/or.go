package logical

import i "goblin.org/main/middleware/irtypes"

type Or struct {
	Destination i.IRTemporary
	Val1        i.IRValue
	Val2        i.IRValue
}

func (Or) Exec(context *i.IRContext) {}
