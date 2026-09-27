package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Sub struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (s *Sub) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	lhsVal, err := state.Resolve(s.Lhs)
	if err != nil {
		return nil, err
	}

	rhsVal, err := state.Resolve(s.Rhs)
	if err != nil {
		return nil, err
	}

	lhs, _ := i.ValueAs[int](lhsVal)
	rhs, _ := i.ValueAs[int](rhsVal)

	// Increment PC.
	state.PC++

	return i.IRNumber{Value: lhs - rhs}, nil
}

func (s *Sub) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(s.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(s.Rhs, context)
	if err != nil {
		return err
	}

	if lhsType != i.IRTypeNumber || rhsType != i.IRTypeNumber {
		return fmt.Errorf("type error: sub: operands of invalid type\n")
	}

	return nil
}

func (s *Sub) String() string {
	return fmt.Sprintf("sub %v %v %v", s.Destination, s.Lhs, s.Rhs)
}

func (s *Sub) Left() i.IRValue {
	return s.Lhs
}

func (s *Sub) Right() i.IRValue {
	return s.Rhs
}

func (s *Sub) Dest() i.IRTemporary {
	return s.Destination
}
