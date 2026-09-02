package logical

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Neg struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
}

func (n *Neg) Exec(context *i.IRContext) {}

func (n *Neg) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(n.Lhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: neg: operand of invalid type\n")
	}

	return nil
}
