package irtypes

import (
	"fmt"
	"strconv"
)

type IRValue interface {
	String() string
	isIRValue()
}

type IRAddress struct {
	Index int
}

type IRTemporary struct {
	Index int
}

type IRLabel struct {
	Value    int
	PCOffset int
}

type IRNumber struct {
	Value int
}

type IRString struct {
	Value string
}

type IRBoolean struct {
	Value bool
}

func (i IRTemporary) String() string { return fmt.Sprintf("%%%v", i.Index) }
func (i IRAddress) String() string   { return fmt.Sprintf("@%v", i.Index) }
func (i IRLabel) String() string     { return fmt.Sprintf("L%v", i.Value) }
func (i IRNumber) String() string    { return fmt.Sprintf("%v", i.Value) }
func (i IRString) String() string    { return fmt.Sprintf("%v", i.Value) }
func (i IRBoolean) String() string   { return strconv.FormatBool(i.Value) }

func (IRAddress) isIRValue()   {}
func (IRTemporary) isIRValue() {}
func (IRNumber) isIRValue()    {}
func (IRLabel) isIRValue()     {}
func (IRString) isIRValue()    {}
func (IRBoolean) isIRValue()   {}

type IRCommand interface {
	Exec(context *IRContext)
	Validate(context *IRContext) error
	String() string
}

type IRBinaryCommand interface {
	IRCommand
	Dest() IRTemporary
	Left() IRValue
	Right() IRValue
}

/*
Resolves incoming IRValue type to a compariable IRType value. Will recursively resolve IRAddress and IRTemporary values.
*/
func ResolveIRType(value IRValue, context *IRContext) (IRType, error) {

	switch val := value.(type) {

	case IRNumber:
		return IRTypeNumber, nil

	case IRString:
		return IRTypeString, nil

	case IRBoolean:
		return IRTypeBoolean, nil

	case IRAddress:
		if val.Index < 0 || val.Index >= len(context.Storage) {
			return IRTypeUndefined, fmt.Errorf("undefined storage address @%d\n", val.Index)
		}
		return ResolveIRType(context.Storage[val.Index], context)

	case IRTemporary:
		if val.Index < 0 || val.Index >= len(context.Temporaries) {
			return IRTypeUndefined, fmt.Errorf("undefined temporary address %%%d\n", val.Index)
		}
		return ResolveIRType(context.Temporaries[val.Index], context)

	case IRLabel:

		if val.Value < 0 || val.Value >= len(context.Labels) {
			return IRTypeUndefined, fmt.Errorf("label[%v] index out of bounds for value %v\n", val, val.Value)
		}

		if val.PCOffset < 0 || val.PCOffset > len(context.Commands) {
			return IRTypeUndefined, fmt.Errorf("label[%v] offset out of bounds for offset %v\n", val, val.PCOffset)
		}

		return IRTypeLabel, nil

	default:
		return IRTypeUndefined, fmt.Errorf("unrecognised type %v\n", value)
	}
}
