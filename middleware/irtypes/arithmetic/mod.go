package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Mod struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (m *Mod) Exec(context *i.IRContext) {}

func (m *Mod) Validate(context *i.IRContext) error {
	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(m.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(m.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: mod: operands of invalid type\n")
	}

	return nil
}

func (m *Mod) String() string {
	return fmt.Sprintf("mod %v %v %v", m.Destination, m.Lhs, m.Rhs)
}

func (m *Mod) Left() i.IRValue {
	return m.Lhs
}

func (m *Mod) Right() i.IRValue {
	return m.Rhs
}

func (m *Mod) Dest() i.IRTemporary {
	return m.Destination
}
