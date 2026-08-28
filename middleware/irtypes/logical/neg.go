package logical

import i "goblin.org/main/middleware/irtypes"

type Neg struct {
	Destination i.IRTemporary
	Val         i.IRValue
}

func (Neg) Exec(context *i.IRContext) {}
