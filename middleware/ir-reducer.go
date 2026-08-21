/*
Goblin IR Reducer v0.1
Author: Chris J.M. Wing
Date: 21/08/2026

Input:
	AST Generated from frontend.Parser.
Output:
	A set of IRCommands that represent the original source program in GoblinIR.
*/

package middleware

import (
	"fmt"

	"goblin.org/main/frontend/ast"
)

/*
Main reducer switch-board. Orchestrates where each ast.Expression goes to be reduced down into GoblinIR.
*/
func reduceExpression(expr ast.Expression, context *IRContext) (IRResult, error) {

	switch value := expr.(type) {

	case ast.AssignmentExpr:

		ir, err := reduceAssignmentExpr(value, context)
		if err != nil {
			return IRResult{}, err
		}
		return ir, nil

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

	case ast.Identifier:

		ir, err := reduceIdentifierExpr(value, context)
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

	address := context.allocateAddress()
	context.storeSymbol(expr.Identifier, address)

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
Reduces an identifier expression down into GoblinIR.
*/
func reduceIdentifierExpr(expr ast.Identifier, context *IRContext) (IRResult, error) {

	address, ok := context.Symbols[expr.Symbol]
	if !ok {
		return IRResult{}, fmt.Errorf("undefined symbol %v\n", expr.Symbol)
	}

	temp := context.allocateTemporary()
	result := IRResult{}
	result.Commands = make([]IRCommand, 0)
	result.Commands = append(result.Commands, &Load{Destination: temp, Source: address})
	result.Value = temp

	return result, nil
}

/*
Reduces an assigmnet expression down into GoblinIR.
*/
func reduceAssignmentExpr(expr ast.AssignmentExpr, context *IRContext) (IRResult, error) {

	iden, ok := expr.Assigne.(ast.Identifier)
	if !ok {
		return IRResult{}, fmt.Errorf("invalid assignment target")
	}

	address, ok := context.Symbols[iden.Symbol]
	if !ok {
		return IRResult{}, fmt.Errorf("undefined symbol %v\n", iden.Symbol)
	}

	rhs, err := reduceExpression(expr.Value, context)
	if err != nil {
		return IRResult{}, nil
	}

	result := IRResult{
		Commands: rhs.Commands,
		Value:    address,
	}

	result.Commands = append(result.Commands, &Store{
		Destination: address,
		Value:       rhs.Value,
	})

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

	destination := context.allocateTemporary()

	switch expr.Operator {
	case "+":
		result.Commands = append(result.Commands, &Add{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		})
	case "-":
		result.Commands = append(result.Commands, &Sub{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		})
	case "*":
		result.Commands = append(result.Commands, &Mul{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		})
	case "/":
		result.Commands = append(result.Commands, &Div{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		})
	case "%":
		result.Commands = append(result.Commands, &Mod{
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

/*
Entry for Reducer called by main program.
*/
func Reduce(program ast.Program) ([]IRCommand, error) {

	context := IRContext{
		Commands:    make([]IRCommand, 0),
		Storage:     make([]IRValue, 0),
		Temporaries: make([]IRValue, 0),
		Symbols:     make(map[string]IRAddress),
		PC:          0,
	}

	for _, expr := range program.Body {

		val, err := reduceExpression(expr, &context)
		if err != nil {
			return nil, fmt.Errorf("reducer error: %v\n", err)
		}

		context.Commands = append(context.Commands, val.Commands...)
	}

	return context.Commands, nil
}

/*
Utility function for pretty-printing the generated GoblinIR.
*/
func PrintIR(prefix string, commands []IRCommand) {

	fmt.Printf("%v\n", prefix)
	for i, command := range commands {
		fmt.Printf("%d:\t%s\n", i, command)
	}
	fmt.Print("\n")
}
