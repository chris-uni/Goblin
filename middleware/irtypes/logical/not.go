package logical

import i "goblin.org/main/middleware/irtypes"

type Not struct {
	Destination i.IRTemporary
	Val         i.IRValue
}

func (n *Not) Exec(context *i.IRContext) {}

func (n *Not) Validate(context *i.IRContext) error {
	return nil
}
