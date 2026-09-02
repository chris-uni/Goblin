package controlflow

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type JmpIf struct {
	Destination i.IRLabel
	Condition   i.IRValue
}

func (ji *JmpIf) Exec(context *i.IRContext) {}

func (ji *JmpIf) Validate(context *i.IRContext) error {

	conditionType, err := i.ResolveIRType(ji.Condition, context)
	if err != nil {
		return err
	}

	_, err = i.ResolveIRType(ji.Destination, context)
	if err != nil {
		return err
	}

	// Is the condition a truthy type?
	if conditionType != i.IRTypeBoolean {
		return fmt.Errorf("type error: jmpif: condition of invalid type\n")
	}

	return nil
}

func (ji *JmpIf) String() string {
	return fmt.Sprintf("jmpif %v %v", ji.Destination, ji.Condition)
}
