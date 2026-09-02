package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Div struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (d *Div) Exec(context *i.IRContext) {}

func (d *Div) Validate(context *i.IRContext) error {
	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(d.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(d.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: div: operands of invalid type\n")
	}

	return nil
}

func (d *Div) String() string {
	return fmt.Sprintf("div %v %v %v", d.Destination, d.Lhs, d.Rhs)
}

func (d *Div) Left() i.IRValue {
	return d.Lhs
}

func (d *Div) Right() i.IRValue {
	return d.Rhs
}

func (d *Div) Dest() i.IRTemporary {
	return d.Destination
}
