package logical

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type And struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (a *And) Exec(context *i.IRContext) {}

func (a *And) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(a.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(a.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeBoolean || rhsType != i.IRTypeBoolean {
		return fmt.Errorf("type error: and: operands of invalid type\n")
	}

	return nil
}
