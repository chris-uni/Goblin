package conditional

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Gt struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (g *Gt) Exec(context *i.IRContext) {}

func (g *Gt) Validate(context *i.IRContext) error {

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
		return fmt.Errorf("type error: gt: operands of invalid type\n")
	}

	return nil
}

func (g *Gt) String() string {
	return fmt.Sprintf("gt %v %v %v", g.Destination, g.Lhs, g.Rhs)
}
