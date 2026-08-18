package middleware

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
	Index map[int]int
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
func (i IRNumber) String() string    { return fmt.Sprintf("%v", i.Value) }
func (i IRString) String() string    { return fmt.Sprintf("%v", i.Value) }
func (i IRBoolean) String() string   { return strconv.FormatBool(i.Value) }

func (IRAddress) isIRValue()   {}
func (IRTemporary) isIRValue() {}
func (IRNumber) isIRValue()    {}
func (IRString) isIRValue()    {}
func (IRBoolean) isIRValue()   {}

type IRCommand interface {
	Exec(context *IRContext)
	String() string
}

type Add struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (a *Add) String() string {
	return fmt.Sprintf("add %v %v %v", a.Destination, a.Lhs, a.Rhs)
}

func (a *Add) Exec(context *IRContext) {}

type Sub struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (Sub) Exec(context *IRContext) {}

type Mul struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (a *Mul) Exec(context *IRContext) {}

type Div struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (a *Div) Exec(context *IRContext) {}

type Mod struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (a *Mod) Exec(context *IRContext) {}

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

func (Load) Exec(context *IRContext) {}

type Jmp struct {
	Destination IRLabel
}

func (Jmp) Exec(context *IRContext) {}

type JmpIf struct {
	Destination IRLabel
	Condition   IRValue
}

func (JmpIf) Exec(context *IRContext) {}

type Eq struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Eq) Exec(context *IRContext) {}

type Neq struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Neq) Exec(context *IRContext) {}

type Lt struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Lt) Exec(context *IRContext) {}

type Lte struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Lte) Exec(context *IRContext) {}

type Gt struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Gt) Exec(context *IRContext) {}

type Gte struct {
	Destination IRTemporary
	Val1        IRValue
	Val2        IRValue
}

func (Gte) Exec(context *IRContext) {}
