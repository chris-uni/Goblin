package logical

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Or struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (o *Or) Exec(context *i.IRContext) {}

func (o *Or) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(o.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(o.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeBoolean || rhsType != i.IRTypeBoolean {
		return fmt.Errorf("type error: or: operands of invalid type\n")
	}

	return nil
}
