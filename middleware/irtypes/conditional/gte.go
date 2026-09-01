package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Gte struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (g *Gte) Exec(context *i.IRContext) {}

func (g *Gte) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(g.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(g.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: gte: operands of invalid type\n")
	}

	return nil
}

func (g *Gte) String() string {
	return fmt.Sprintf("gte %v %v %v", g.Destination, g.Lhs, g.Rhs)
}
