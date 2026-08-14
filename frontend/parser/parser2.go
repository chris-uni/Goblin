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
	10. Call
	11. Member
	12. Primary
*/

package parser

import (
	"fmt"

	"goblin.org/main/frontend/ast"
	"goblin.org/main/frontend/lexer"
	"goblin.org/main/utils"
)

func (p *Parser) parseExpression() (ast.Expression, error) {
	return p.parsePrimaryExpression()
}

func (p *Parser) parseGroupedExpression() (ast.Expression, error) {

	p.consume() // '('

	expr, err := p.parseExpression()
	if err != nil {
		return ast.Expr{}, err
	}

	if _, err := p.expect(lexer.CloseParen); err != nil {
		return ast.Expr{}, nil
	}

	return expr, nil
}

func (p *Parser) parseIdentifier() (ast.Expression, error) {

	return ast.Expr{}, nil
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
		_, err := p.parseIdentifier()
		if err != nil {
			return ast.Expr{}, nil
		}

		return ast.Identifier{
			Kind:   ast.IdentifierNode,
			Symbol: p.consume().Value,
		}, nil

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

		node, err := parser.parsePrimaryExpression()
		if err != nil {
			return ast.Program{}, err
		}

		program.Body = append(program.Body, node)
	}

	return program, nil
}
