package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Lt struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (l *Lt) Exec(context *i.IRContext) {}

func (l *Lt) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(l.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(l.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: lt: operands of invalid type\n")
	}

	return nil
}

func (l *Lt) String() string {
	return fmt.Sprintf("lt %v %v %v", l.Destination, l.Lhs, l.Rhs)
}
