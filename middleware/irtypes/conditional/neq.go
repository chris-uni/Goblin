package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Neq struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (n *Neq) Exec(context *i.IRContext) {}

func (n *Neq) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(n.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(n.Rhs, context)
	if err != nil {
		return err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeBoolean) &&
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeBoolean) {
		return fmt.Errorf("type error: eq: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return fmt.Errorf("type error: eq: incompatible types\n")
	}

	return nil
}

func (n *Neq) String() string {
	return fmt.Sprintf("neq %v %v %v", n.Destination, n.Lhs, n.Rhs)
}
