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

func (n *Neq) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	lhsVal, err := state.Resolve(n.Lhs)
	if err != nil {
		return nil, err
	}

	rhsVal, err := state.Resolve(n.Rhs)
	if err != nil {
		return nil, err
	}

	var lhs any
	var rhs any

	switch l := lhsVal.(type) {

	case i.IRNumber:

		lhs, err = i.ValueAs[int](l)
		if err != nil {
			return nil, err
		}

		rhs, err = i.ValueAs[int](rhsVal)
		if err != nil {
			return nil, err
		}

	case i.IRString:

		lhs, err = i.ValueAs[string](l)
		if err != nil {
			return nil, err
		}

		rhs, err = i.ValueAs[string](rhsVal)
		if err != nil {
			return nil, err
		}

	case i.IRBoolean:

		lhs, err = i.ValueAs[bool](l)
		if err != nil {
			return nil, err
		}

		rhs, err = i.ValueAs[bool](rhsVal)
		if err != nil {
			return nil, err
		}
	}

	state.PC++
	return i.IRBoolean{Value: lhs != rhs}, nil
}

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

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeBoolean || lhsType == i.IRTypeString) ||
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeBoolean || rhsType == i.IRTypeString) {
		return fmt.Errorf("type error: neq: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return fmt.Errorf("type error: neq: incompatible types\n")
	}

	return nil
}

func (n *Neq) String() string {
	return fmt.Sprintf("neq %v %v %v", n.Destination, n.Lhs, n.Rhs)
}
