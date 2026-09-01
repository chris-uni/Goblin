package logical

import i "goblin.org/main/middleware/irtypes"

type Or struct {
	Destination i.IRTemporary
	Val1        i.IRValue
	Val2        i.IRValue
}

func (o *Or) Exec(context *i.IRContext) {}

func (o *Or) Validate(context *i.IRContext) error {
	return nil
}
