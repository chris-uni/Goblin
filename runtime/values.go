package runtime

import (
	"fmt"

	"goblin.org/main/frontend/ast"
)

type ValueType string

const (
	Number      ValueType = "Number"
	Array       ValueType = "Array"
	Map         ValueType = "Map"
	Null        ValueType = "Null"
	Boolean     ValueType = "Boolean"
	String      ValueType = "String"
	Object      ValueType = "Object"
	NativeFn    ValueType = "NativeFn"
	UserFn      ValueType = "UserFn"
	Conditional ValueType = "Contitional"
)

type RuntimeValue interface {
	runtime()
}

type RntmVal struct {
	Type  ValueType
	Value string
}

func (rtm RntmVal) runtime() {}

type NullValue struct {
	Type  ValueType
	Value string
}

func (n NullValue) runtime() {}

type NumberValue struct {
	Type  ValueType
	Value int
}

func (n NumberValue) runtime() {
	fmt.Printf("%v\n", n.Value)
}

type ArrayValue struct {
	Type  ValueType
	Value []RuntimeValue
}

func (a ArrayValue) runtime() {
	fmt.Printf("%v\n", a.Value)
}

type MapValue struct {
	Type  ValueType
	Value map[RuntimeValue]RuntimeValue
}

func (m MapValue) runtime() {
	fmt.Printf("%v\n", m.Value)
}

type StringValue struct {
	Type  ValueType
	Value string
}

func (s StringValue) runtime() {}

type BooleanValue struct {
	Type  ValueType
	Value bool
}

func (b BooleanValue) runtime() {}

type ObjectVal struct {
	Type       ValueType
	Properties map[string]RuntimeValue
}

func (o ObjectVal) runtime() {}

type FunctionCall func(args []RuntimeValue, env Environment) RuntimeValue

type NativeFunction struct {
	Type ValueType
	Call FunctionCall
}

func (n NativeFunction) runtime() {}

type UserFunction struct {
	Type   ValueType
	Name   string
	Params []string
	DecEnv Environment
	Body   []ast.Expression
}

func (f UserFunction) runtime() {}

type IfConditional struct {
	Type      string
	Condition ast.BinaryExpr
	Body      []ast.Expression
}

func (i IfConditional) runtime() {}

type WhileValue struct {
	Type      string
	Condition ast.BinaryExpr
	Body      []ast.Expression
}

func (w WhileValue) runtime() {}

/* -- MACROS --*/

func MK_NUMBER(n int) NumberValue {

	return NumberValue{
		Type:  "Number",
		Value: n,
	}
}

func MK_ARRAY(elements []RuntimeValue) ArrayValue {

	return ArrayValue{
		Type:  "Array",
		Value: elements,
	}
}

func MK_MAP(elements map[RuntimeValue]RuntimeValue) MapValue {

	return MapValue{
		Type:  "Map",
		Value: elements,
	}
}

func MK_STRING(s string) StringValue {

	return StringValue{
		Type:  "String",
		Value: s,
	}
}

func MK_NULL() NullValue {

	return NullValue{
		Type:  "Null",
		Value: "null",
	}
}

func MK_BOOL(b bool) BooleanValue {

	return BooleanValue{
		Type:  "Boolean",
		Value: b,
	}
}

func MK_NATIVE_FN(call FunctionCall) NativeFunction {

	return NativeFunction{
		Type: "NativeFn",
		Call: call,
	}
}
