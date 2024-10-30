package parser

import (
	"fmt"
	"strconv"

	ast "goblin.org/main/AST"
	lexer "goblin.org/main/Lexer"
	tokens "goblin.org/main/Tokens"
	tools "goblin.org/main/Tools"
)

var tkns = make([]tokens.Token, 0)

func NotEOF() bool {
	return tkns[0].Type != tokens.EOF
}

// Returns the current token.
func at() tokens.Token {
	return tkns[0]
}

// Returns the current token and shift the slice along.
func eat() tokens.Token {

	prev := tools.Shift(&tkns)
	return prev
}

// Same functionality as 'eat()', however also allows us to check what the next token _should_ be.
func expect(tkn tokens.TokenType) (tokens.Token, error) {

	prev := eat()

	if prev.Type != tkn {

		return tokens.Token{}, fmt.Errorf("expected %v, received %v", tkn, prev.Type)
	}

	return prev, nil
}

func BuildAst(source string) ast.Program {

	tkns = lexer.Tokenize(source)

	program := ast.Program{
		Kind: ast.ProgramNode,
		Body: []ast.Expression{},
	}

	for NotEOF() {
		program.Body = append(program.Body, parse_statement())
	}

	return program
}

func parse_statement() ast.Expression {

	// For the min we will just go straight to ParseExpression.
	return parse_expression()
}

func parse_expression() ast.Expression {

	return parse_additive_expression()
}

/*
First must evaluate the LHS of the expression.

E.g. 10 - 5 + 5

LHS = 10 - 5
RHS = LHS + 5
*/
func parse_additive_expression() ast.Expression {

	left := parse_multiplicitive_expression()

	for at().Value == "+" || at().Value == "-" {

		operator := eat().Value
		right := parse_multiplicitive_expression()

		left = ast.BinaryExpr{
			Kind:     ast.BinaryExpressionNode,
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}

	return left
}

func parse_multiplicitive_expression() ast.Expression {

	left := parse_primary_expression()

	for at().Value == "/" || at().Value == "*" || at().Value == "%" {

		operator := eat().Value
		right := parse_primary_expression()

		left = ast.BinaryExpr{
			Kind:     ast.BinaryExpressionNode,
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}

	return left
}

func parse_primary_expression() ast.Expression {

	tkn := at().Type

	switch tkn {

	case tokens.Identifier:
		return ast.Identifier{
			Kind:   ast.IdentifierNode,
			Symbol: eat().Value,
		}

	case tokens.Number:
		number, _ := strconv.ParseFloat(eat().Value, 64)
		return ast.NumericExpr{
			Kind:  ast.NumericLiteralNode,
			Value: number,
		}

	case tokens.OpenParen:
		eat() // Eat opening paren
		val := parse_expression()
		_, err := expect(tokens.CloseParen)
		if err != nil {
			fmt.Printf("%v \n", err.Error())
			return ast.Expr{}
		}
		return val

	default:
		e := fmt.Sprintf("unexpected expression: %v \n", tkn)
		panic(e)
	}
}
