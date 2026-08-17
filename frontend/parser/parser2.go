/*
Goblin Parser v0.1
Author: Chris J.M. Wing

Input:
	Slice of Goblin source-code represented by Goblins Lexer as Lexical Tokens.
Output:
	An AST Tree representing the Goblin program.

Guarantees:
	Orders of prescidence

	1. Assignment
	2. Ternary
	3. Logical OR
	4. Logical AND
	5. Equality
	6. Comparision
	7. Additive
	8. Multiplicative
	9. Unary
	10. Call			} Part of the parsePostfix collection
	11. Member			}
	12. Primary
*/

package parser

import (
	"fmt"

	"goblin.org/main/frontend/ast"
	"goblin.org/main/frontend/lexer"
	"goblin.org/main/utils"
)

func (p *Parser) parseStatement() (ast.Expression, error) {

	if p.match(lexer.Let) {
		return p.parseVariableDecleration()
	} else {
		return p.parseExpression()
	}
}

func (p *Parser) parseExpression() (ast.Expression, error) {
	return p.parseAssignmentExpression()
}

/*
-- 1.0 Assignment Section --
*/
func (p *Parser) parseAssignmentExpression() (ast.Expression, error) {

	lhs, err := p.parseAdditiveExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	// Not necessairly an assignment, so return what we have.
	if !p.match(lexer.Equals) {

		return lhs, nil
	}

	// Else this is an assignment, so enforce that lhs is an identifier.
	_, isIdentifier := lhs.(ast.Identifier)
	if !isIdentifier {
		return ast.Expr{}, fmt.Errorf("epecting lhs of assignment to be identifier type, got %v\n", lhs)
	}

	p.consume() // Consume the '='

	rhs, err := p.parseExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	return ast.AssignmentExpr{
		Kind:    ast.AssingmentExprNode,
		Value:   rhs,
		Assigne: lhs,
	}, nil
}

/*
-- 1.1 Variable Assignment Section --
*/
func (p *Parser) parseVariableDecleration() (ast.Expression, error) {

	p.consume() // Consume the 'let keyword'.

	lhs, err := p.expect(lexer.Identifier)
	if err != nil {
		return ast.Expr{}, err
	}

	_, err = p.expect(lexer.Equals) // Expecting an '=' symbol here.
	if err != nil {
		return ast.Expr{}, err
	}

	rhs, err := p.parseExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	return ast.VariableDecleration{
		Kind:       ast.VariableDeclarationNode,
		Value:      rhs,
		Identifier: lhs.Value,
		Constant:   false,
	}, nil
}

/*
-- 7. Additive Expression Section --
*/
func (p *Parser) parseAdditiveExpression() (ast.Expression, error) {

	lhs, err := p.parseMultiplicativeExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	for p.match(lexer.BinaryOperator) {

		op := p.peek()
		if op.Value != "+" && op.Value != "-" {
			break
		}

		p.consume() // Consume either the '+' or '-'.

		rhs, err := p.parseMultiplicativeExpression()
		if err != nil {
			return ast.Expr{}, err
		}

		lhs = ast.BinaryExpr{
			Kind:     ast.BinaryExprNode,
			Left:     lhs,
			Right:    rhs,
			Operator: op.Value,
		}
	}

	return lhs, nil
}

/*
-- 8. Multiplicative Expression Section --
*/
func (p *Parser) parseMultiplicativeExpression() (ast.Expression, error) {

	lhs, err := p.parsePrimaryExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	for p.match(lexer.BinaryOperator) {

		op := p.peek()
		if op.Value != "*" && op.Value != "/" && op.Value != "%" {
			break
		}

		p.consume() // Consume either the '*', '/' or '%'.

		rhs, err := p.parsePrimaryExpression()
		if err != nil {
			return ast.Expr{}, err
		}

		lhs = ast.BinaryExpr{
			Kind:     ast.BinaryExprNode,
			Left:     lhs,
			Right:    rhs,
			Operator: op.Value,
		}
	}

	return lhs, nil
}

/*
-- 12. Primary Expression Section --
*/
func (p *Parser) parseGroupedExpression() (ast.Expression, error) {

	p.consume() // '('

	expr, err := p.parseExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	if _, err := p.expect(lexer.CloseParen); err != nil {
		return ast.Expr{}, err
	}

	return expr, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {

	return ast.Identifier{
		Kind:   ast.IdentifierNode,
		Symbol: p.consume().Value,
	}, nil
}

func (p *Parser) parsePrimaryExpression() (ast.Expression, error) {

	switch tkn := p.peek().Type; {

	case p.match(lexer.Number):
		val, err := utils.ToNumber(p.consume().Value)
		if err != nil {
			return ast.Expr{}, err
		}

		return ast.NumericLiteral{
			Kind:  ast.NumericLiteralNode,
			Value: val,
		}, nil

	case p.match(lexer.String):
		return ast.StringLiteral{
			Kind:  ast.StringLiteralNode,
			Value: p.consume().Value,
		}, nil

	case tkn == lexer.Boolean:
		return ast.BooleanLiteral{
			Kind:  ast.BooleanLiteralNode,
			Value: utils.StoB(p.consume().Value),
		}, nil

	case tkn == lexer.Identifier:
		val, err := p.parseIdentifier()
		if err != nil {
			return ast.Expr{}, nil
		}

		return val, err

	case p.match(lexer.OpenParen):
		val, err := p.parseGroupedExpression()
		if err != nil {
			return ast.Expr{}, err
		}

		return val, nil

	default:
		message := fmt.Sprintf("unexpected token found during parsing '%v'", p.consume().Value)
		return ast.Expr{}, fmt.Errorf("%v", message)
	}
}

/*
-- Main Token parsing entry.
*/
func ParseTokens(tokens []lexer.Token) (ast.Program, error) {

	parser := Parser{
		Tokens: tokens,
		Audit:  make(map[int]string, 0),
		Pos:    0,
	}

	program := ast.Program{
		Kind: "Program",
		Body: []ast.Expression{},
	}

	for !parser.match(lexer.EOF) {

		node, err := parser.parseStatement()
		if err != nil {
			return ast.Program{}, err
		}

		program.Body = append(program.Body, node)

		// Each statement should be terminated by an EOL ';'.
		_, err = parser.expect(lexer.EOL)
		if err != nil {
			return ast.Program{}, err
		}
	}

	return program, nil
}
