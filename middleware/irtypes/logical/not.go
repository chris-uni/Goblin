package logical

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Not struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
}

func (n *Not) Exec(context *i.IRContext) {}

func (n *Not) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(n.Lhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeBoolean {
		return fmt.Errorf("type error: not: operand of invalid type\n")
	}

	return nil
}
