package middleware

import (
	"fmt"

	"goblin.org/main/frontend/ast"
)

func reduceExpression(expr ast.Expression, context *IRContext) (IRResult, error) {

	switch value := expr.(type) {

	case ast.VariableDecleration:

		ir, err := reduceVariableDecleration(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil

	case ast.BinaryExpr:

		ir, err := reduceBinaryExpr(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil

	case ast.NumericLiteral:

		ir, err := reduceNumericExpr(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil

	case ast.StringLiteral:

		ir, err := reduceStringExpr(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil

	case ast.BooleanLiteral:

		ir, err := reduceBooleanExpr(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil
	}

	return IRResult{}, nil
}

/*
Reduces a VariableDeclerationExpr down into GoblinIR.
*/
func reduceVariableDecleration(expr ast.VariableDecleration, context *IRContext) (IRResult, error) {

	value, err := reduceExpression(expr.Value, context)
	if err != nil {
		return IRResult{}, err
	}

	address := IRAddress{
		Index: len(context.Storage),
	}

	result := IRResult{
		Commands: value.Commands,
	}

	result.Commands = append(
		result.Commands,
		&Store{
			Destination: address,
			Value:       value.Value,
		},
	)

	return result, nil
}

/*
Reduces a BinaryExpr down into GoblinIR.
*/
func reduceBinaryExpr(expr ast.BinaryExpr, context *IRContext) (IRResult, error) {

	lhs, err := reduceExpression(expr.Left, context)
	if err != nil {
		return IRResult{}, err
	}

	rhs, err := reduceExpression(expr.Right, context)
	if err != nil {
		return IRResult{}, err
	}

	result := IRResult{}

	result.Commands = append(result.Commands, lhs.Commands...)
	result.Commands = append(result.Commands, rhs.Commands...)

	destination := IRTemporary{
		Index: len(context.Temporaries),
	}

	switch expr.Operator {
	case "+":
		result.Commands = append(result.Commands, &Add{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		})
	}

	result.Value = destination
	return result, nil
}

/*
Reduces a NumericExpr down into GoblinIR.
*/
func reduceNumericExpr(expr ast.NumericLiteral, _ *IRContext) (IRResult, error) {

	result := IRResult{}
	result.Value = IRNumber{Value: expr.Value}

	return result, nil
}

/*
Reduces a StringExpr down into GoblinIR.
*/
func reduceStringExpr(expr ast.StringLiteral, _ *IRContext) (IRResult, error) {

	result := IRResult{}
	result.Value = IRString{Value: expr.Value}

	return result, nil
}

/*
Reduces a BooleanExpr down into GoblinIR.
*/
func reduceBooleanExpr(expr ast.BooleanLiteral, _ *IRContext) (IRResult, error) {

	result := IRResult{}
	result.Value = IRBoolean{Value: expr.Value}

	return result, nil
}

func Reduce(program ast.Program) ([]IRCommand, error) {

	context := IRContext{
		Commands:    make([]IRCommand, 0),
		Storage:     make([]IRValue, 0),
		Temporaries: make([]IRValue, 0),
		PC:          0,
	}

	for _, expr := range program.Body {

		val, err := reduceExpression(expr, &context)
		if err != nil {
			return nil, err
		}

		context.Commands = append(context.Commands, val.Commands...)
	}

	return context.Commands, nil
}

func PrintIR(commands []IRCommand) {
	for i, command := range commands {
		fmt.Printf("%d:\t%s\n", i, command)
	}
}
