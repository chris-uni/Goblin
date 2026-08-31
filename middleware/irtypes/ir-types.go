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
	String() string
}

type IRBinaryCommand interface {
	IRCommand
	Dest() IRTemporary
	Left() IRValue
	Right() IRValue
}
