package logical

import i "goblin.org/main/middleware/irtypes"

type And struct {
	Destination i.IRTemporary
	Val1        i.IRValue
	Val2        i.IRValue
}

func (a *And) Exec(context *i.IRContext) {}

func (a *And) Validate(context *i.IRContext) error {
	return nil
}
