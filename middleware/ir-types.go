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

func (c *IRContext) allocateAddress() IRAddress {

	storage := IRAddress{
		Index: len(c.Storage),
	}

	c.Storage = append(c.Storage, nil)
	return storage
}

func (c *IRContext) storeSymbol(name string, address IRAddress) {

	c.Symbols[name] = address
}

type IRTemporary struct {
	Index int
}

func (c *IRContext) allocateTemporary() IRTemporary {

	temporary := IRTemporary{
		Index: len(c.Temporaries),
	}

	c.Temporaries = append(c.Temporaries, nil)
	return temporary
}

type IRLabel struct {
	Value    int
	PCOffset int
}

func (c *IRContext) allocateLabel() IRLabel {

	label := IRLabel{
		Value: len(c.Labels),
	}

	c.Labels = append(c.Labels, label)
	return label
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

type Add struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Add) String() string          { return fmt.Sprintf("add %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Add) Exec(context *IRContext) {}
func (v *Add) Left() IRValue           { return v.Lhs }
func (v *Add) Right() IRValue          { return v.Rhs }
func (v *Add) Dest() IRTemporary       { return v.Destination }

type Sub struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Sub) String() string          { return fmt.Sprintf("sub %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Sub) Exec(context *IRContext) {}
func (v *Sub) Left() IRValue           { return v.Lhs }
func (v *Sub) Right() IRValue          { return v.Rhs }
func (v *Sub) Dest() IRTemporary       { return v.Destination }

type Mul struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Mul) String() string          { return fmt.Sprintf("mul %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Mul) Exec(context *IRContext) {}
func (v *Mul) Left() IRValue           { return v.Lhs }
func (v *Mul) Right() IRValue          { return v.Rhs }
func (v *Mul) Dest() IRTemporary       { return v.Destination }

type Div struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Div) String() string          { return fmt.Sprintf("div %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Div) Exec(context *IRContext) {}
func (v *Div) Left() IRValue           { return v.Lhs }
func (v *Div) Right() IRValue          { return v.Rhs }
func (v *Div) Dest() IRTemporary       { return v.Destination }

type Mod struct {
	Destination IRTemporary
	Lhs         IRValue
	Rhs         IRValue
}

func (v *Mod) String() string          { return fmt.Sprintf("mod %v %v %v", v.Destination, v.Lhs, v.Rhs) }
func (v *Mod) Exec(context *IRContext) {}
func (v *Mod) Left() IRValue           { return v.Lhs }
func (v *Mod) Right() IRValue          { return v.Rhs }
func (v *Mod) Dest() IRTemporary       { return v.Destination }

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

type Lbl struct {
	Destination int
}

func (v *Lbl) String() string {
	return fmt.Sprintf("L%v:", v.Destination)
}

func (Lbl) Exec(context *IRContext) {}

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
