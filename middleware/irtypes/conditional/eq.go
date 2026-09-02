package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Eq struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (e *Eq) Exec(context *i.IRContext) {}

func (e *Eq) Validate(context *i.IRContext) error {
	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(e.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(e.Rhs, context)
	if err != nil {
		return err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeBoolean || lhsType == i.IRTypeString) ||
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeBoolean || rhsType == i.IRTypeString) {
		return fmt.Errorf("type error: eq: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return fmt.Errorf("type error: eq: incompatible types\n")
	}

	return nil
}

func (e *Eq) String() string {
	return fmt.Sprintf("eq %v %v %v", e.Destination, e.Lhs, e.Rhs)
}
