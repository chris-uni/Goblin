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
	i "goblin.org/main/middleware/irtypes"
	a "goblin.org/main/middleware/irtypes/arithmetic"
	c "goblin.org/main/middleware/irtypes/conditional"
	f "goblin.org/main/middleware/irtypes/controlflow"
	m "goblin.org/main/middleware/irtypes/memory"
)

/*
Main reducer switch-board. Orchestrates where each ast.Expression goes to be reduced down into GoblinIR.
*/
func reduceExpression(expr ast.Expression, context *i.IRContext) (i.IRResult, error) {

	switch value := expr.(type) {

	case ast.IfCondition:

		ir, err := reduceIfExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.AssignmentExpr:

		ir, err := reduceAssignmentExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.VariableDecleration:

		ir, err := reduceVariableDecleration(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.BinaryExpr:

		ir, err := reduceBinaryExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.Identifier:

		ir, err := reduceIdentifierExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.NumericLiteral:

		ir, err := reduceNumericExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.StringLiteral:

		ir, err := reduceStringExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	case ast.BooleanLiteral:

		ir, err := reduceBooleanExpr(value, context)
		if err != nil {
			return i.IRResult{}, err
		}
		return ir, nil

	default:
		return i.IRResult{}, fmt.Errorf("unknonwn expression found %v\n", expr)
	}
}

/*
Reduces an IfExpression down into GoblinIR.
*/
func reduceIfExpr(expr ast.IfCondition, context *i.IRContext) (i.IRResult, error) {

	condition, err := reduceExpression(expr.Condition, context)
	if err != nil {
		return i.IRResult{}, err
	}

	result := i.IRResult{
		Commands: condition.Commands,
	}

	// If Body label.
	bodyLabel := context.AllocateLabel()

	jmpOffset := context.PC + 2 // 1 for the jmpif and 1 for jmp commands
	if len(expr.Body) == 0 {
		jmpOffset = context.PC
	}

	bodyLabel.PCOffset = jmpOffset

	jmpIf := &f.JmpIf{
		Destination: bodyLabel,
		Condition:   condition.Value,
	}

	context.Push(jmpIf)
	result.Commands = append(result.Commands, jmpIf)

	endLabel := context.AllocateLabel()

	jmp := &f.Jmp{
		Destination: endLabel,
	}
	context.Push(jmp)
	result.Commands = append(result.Commands, jmp)

	var bodyCmds []i.IRCommand

	for _, stmt := range expr.Body {

		res, err := reduceExpression(stmt, context)
		if err != nil {
			return i.IRResult{}, err
		}

		bodyCmds = append(bodyCmds, res.Commands...)
	}

	jmpOffset = bodyLabel.PCOffset + len(bodyCmds)
	if len(expr.Body) == 0 {
		jmpOffset = bodyLabel.PCOffset
	}

	jmp.Destination.PCOffset = jmpOffset

	result.Commands = append(result.Commands, bodyCmds...)
	return result, nil
}

/*
Reduces a VariableDeclerationExpr down into GoblinIR.
*/
func reduceVariableDecleration(expr ast.VariableDecleration, context *i.IRContext) (i.IRResult, error) {

	value, err := reduceExpression(expr.Value, context)
	if err != nil {
		return i.IRResult{}, err
	}

	address := context.AllocateAddress()
	context.StoreSymbol(expr.Identifier, address)

	result := i.IRResult{
		Commands: value.Commands,
	}

	cmd := &m.Store{
		Destination: address,
		Value:       value.Value,
	}

	context.Push(cmd)
	result.Commands = append(result.Commands, cmd)

	context.Storage[address.Index] = value.Value

	return result, nil
}

/*
Reduces an identifier expression down into GoblinIR.
*/
func reduceIdentifierExpr(expr ast.Identifier, context *i.IRContext) (i.IRResult, error) {

	address, ok := context.Symbols[expr.Symbol]
	if !ok {
		return i.IRResult{}, fmt.Errorf("undefined symbol %v\n", expr.Symbol)
	}

	temp := context.AllocateTemporary()
	result := i.IRResult{}
	result.Commands = make([]i.IRCommand, 0)
	cmd := &m.Load{Destination: temp, Source: address}
	result.Commands = append(result.Commands, cmd)
	result.Value = temp

	context.Push(cmd)

	return result, nil
}

/*
Reduces an assigmnet expression down into GoblinIR.
*/
func reduceAssignmentExpr(expr ast.AssignmentExpr, context *i.IRContext) (i.IRResult, error) {

	iden, ok := expr.Assigne.(ast.Identifier)
	if !ok {
		return i.IRResult{}, fmt.Errorf("invalid assignment target")
	}

	address, ok := context.Symbols[iden.Symbol]
	if !ok {
		return i.IRResult{}, fmt.Errorf("undefined symbol %v\n", iden.Symbol)
	}

	rhs, err := reduceExpression(expr.Value, context)
	if err != nil {
		return i.IRResult{}, nil
	}

	result := i.IRResult{
		Commands: rhs.Commands,
		Value:    address,
	}

	cmd := &m.Store{
		Destination: address,
		Value:       rhs.Value,
	}

	context.Push(cmd)
	result.Commands = append(result.Commands, cmd)

	return result, nil
}

/*
Reduces a BinaryExpr down into GoblinIR.
*/
func reduceBinaryExpr(expr ast.BinaryExpr, context *i.IRContext) (i.IRResult, error) {

	lhs, err := reduceExpression(expr.Left, context)
	if err != nil {
		return i.IRResult{}, err
	}

	rhs, err := reduceExpression(expr.Right, context)
	if err != nil {
		return i.IRResult{}, err
	}

	result := i.IRResult{}

	result.Commands = append(result.Commands, lhs.Commands...)
	result.Commands = append(result.Commands, rhs.Commands...)

	destination := context.AllocateTemporary()

	switch expr.Operator {
	case "+":
		cmd := &a.Add{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRNumber{}

	case "-":
		cmd := &a.Sub{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRNumber{}

	case "*":
		cmd := &a.Mul{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRNumber{}

	case "/":
		cmd := &a.Div{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRNumber{}

	case "%":
		cmd := &a.Mod{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRNumber{}

	case ">":
		cmd := &c.Gt{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}

	case ">=":
		cmd := &c.Gte{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}

	case "<":
		cmd := &c.Lt{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}

	case "<=":
		cmd := &c.Lte{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}

	case "==":
		cmd := &c.Eq{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}

	case "!=":
		cmd := &c.Neq{
			Destination: destination,
			Lhs:         lhs.Value,
			Rhs:         rhs.Value,
		}

		context.Push(cmd)
		result.Commands = append(result.Commands, cmd)
		context.Temporaries[destination.Index] = i.IRBoolean{}
	}

	result.Value = destination
	return result, nil
}

/*
Reduces a NumericExpr down into GoblinIR.
*/
func reduceNumericExpr(expr ast.NumericLiteral, _ *i.IRContext) (i.IRResult, error) {

	result := i.IRResult{}
	result.Value = i.IRNumber{Value: expr.Value}

	return result, nil
}

/*
Reduces a StringExpr down into GoblinIR.
*/
func reduceStringExpr(expr ast.StringLiteral, _ *i.IRContext) (i.IRResult, error) {

	result := i.IRResult{}
	result.Value = i.IRString{Value: expr.Value}

	return result, nil
}

/*
Reduces a BooleanExpr down into GoblinIR.
*/
func reduceBooleanExpr(expr ast.BooleanLiteral, _ *i.IRContext) (i.IRResult, error) {

	result := i.IRResult{}
	result.Value = i.IRBoolean{Value: expr.Value}

	return result, nil
}

/*
Entry for Reducer called by main program.
*/
func Reduce(program ast.Program, context *i.IRContext) ([]i.IRCommand, error) {

	for _, expr := range program.Body {

		_, err := reduceExpression(expr, context)
		if err != nil {
			return nil, fmt.Errorf("reducer error: %v\n", err)
		}
	}

	return context.Commands, nil
}

/*
Utility function for pretty-printing the generated GoblinIR.
*/
func PrintIR(prefix string, commands []i.IRCommand) {

	fmt.Printf("%v\n", prefix)
	for i, command := range commands {
		fmt.Printf("%d:\t%s\n", i, command)
	}
	fmt.Print("\n")
}
