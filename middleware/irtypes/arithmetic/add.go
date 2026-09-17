package arithmetic

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Add struct {
	Destination i.IRTemporary
	Lhs         i.IRValue
	Rhs         i.IRValue
}

func (a *Add) Exec(context *i.IRExecutionState) i.IRValue {

	isStringAddition := true

	_, err := i.ValueAs[string](a.Lhs)
	if err != nil {
		isStringAddition = false
	}

	if isStringAddition {
		lhs, _ := i.ValueAs[string](a.Lhs)
		rhs, _ := i.ValueAs[string](a.Rhs)

		return i.IRString{Value: lhs + rhs}
	}

	lhs, _ := i.ValueAs[int](a.Lhs)
	rhs, _ := i.ValueAs[int](a.Rhs)

	return i.IRNumber{Value: lhs + rhs}

}

func (a *Add) Validate(context *i.IRContext) error {

	// Does both the lhs and rhs of the command adhere to the commands rules?
	lhsType, err := i.ResolveIRType(a.Lhs, context)
	if err != nil {
		return err
	}

	rhsType, err := i.ResolveIRType(a.Rhs, context)
	if err != nil {
		return err
	}

	if !(lhsType == i.IRTypeNumber || lhsType == i.IRTypeString) ||
		!(rhsType == i.IRTypeNumber || rhsType == i.IRTypeString) {
		return fmt.Errorf("type error: add: operands of invalid type\n")
	}

	if lhsType != rhsType {
		return fmt.Errorf("type error: add: incompatible types\n")
	}

	return nil
}

func (a *Add) String() string {
	return fmt.Sprintf("add %v %v %v", a.Destination, a.Lhs, a.Rhs)
}

func (a *Add) Left() i.IRValue {
	return a.Lhs
}

func (a *Add) Right() i.IRValue {
	return a.Rhs
}

func (a *Add) Dest() i.IRTemporary {
	return a.Destination
}
