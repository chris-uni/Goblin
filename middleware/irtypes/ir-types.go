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

type Neg struct {
	Destination IRTemporary
	Val         IRValue
}

func (Neg) Exec(context *IRContext) {}

type Not struct {
	Destination IRTemporary
	Val         IRValue
}

func (Not) Exec(context *IRContext) {}

type And struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (And) Exec(context *IRContext) {}

type Or struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Or) Exec(context *IRContext) {}

type Store struct {
	Destination IRAddress
	Value       IRValue
}

func (s *Store) String() string {
	return fmt.Sprintf("store %v %v", s.Destination, s.Value)
}
func (Store) Exec(context *IRContext) {}

type Load struct {
	Destination IRTemporary
	Source      IRAddress
}

func (v *Load) String() string {
	return fmt.Sprintf("load %v %v", v.Destination, v.Source)
}

func (Load) Exec(context *IRContext) {}

type Jmp struct {
	Destination IRLabel
}

func (v *Jmp) String() string {
	return fmt.Sprintf("jmp %v", v.Destination)
}

func (Jmp) Exec(context *IRContext) {}

type JmpIf struct {
	Destination IRLabel
	Condition   IRValue
}

func (v *JmpIf) String() string {
	return fmt.Sprintf("jmpif %v %v", v.Destination, v.Condition)
}

func (JmpIf) Exec(context *IRContext) {}

type Eq struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (s *Eq) String() string {
	return fmt.Sprintf("eq %v %v %v", s.Destination, s.Lhs, s.Rhs)
}

func (Eq) Exec(context *IRContext) {}

type Neq struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (s *Neq) String() string {
	return fmt.Sprintf("neq %v %v %v", s.Destination, s.Lhs, s.Rhs)
}

func (Neq) Exec(context *IRContext) {}

type Lt struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Lt) String() string {
	return fmt.Sprintf("lt %v %v %v", v.Destination, v.Lhs, v.Rhs)
}

func (Lt) Exec(context *IRContext) {}

type Lte struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Lte) String() string {
	return fmt.Sprintf("lte %v %v %v", v.Destination, v.Lhs, v.Rhs)
}
func (Lte) Exec(context *IRContext) {}

type Gt struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Gt) String() string {
	return fmt.Sprintf("gt %v %v %v", v.Destination, v.Lhs, v.Rhs)
}
func (Gt) Exec(context *IRContext) {}

type Gte struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Gte) String() string {
	return fmt.Sprintf("gte %v %v %v", v.Destination, v.Lhs, v.Rhs)
}

func (Gte) Exec(context *IRContext) {}
