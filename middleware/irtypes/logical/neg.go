package logical

import i "goblin.org/main/middleware/irtypes"

type Neg struct {
	Destination i.IRTemporary
	Val         i.IRValue
}

func (n *Neg) Exec(context *i.IRContext) {}

func (n *Neg) Validate(context *i.IRContext) error {
	return nil
}
