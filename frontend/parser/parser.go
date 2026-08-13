package parser

import (
	"fmt"

	"goblin.org/main/frontend/ast"
	"goblin.org/main/frontend/lexer"
	"goblin.org/main/utils"
)

/*
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

type Parser struct {
	Tokens []lexer.Token
	Pos    int
	Audit  map[int]string
}

/*
Non-consuming. Returns the token at the current pointer position.
*/
func (p *Parser) peek() lexer.Token {
	return p.Tokens[p.Pos]
}

/*
Consuming. Returns the current token and increments the source pointer along.
*/
func (p *Parser) consume() lexer.Token {

	prev := p.peek()
	p.Pos++
	return prev
}

/*
Non-consuming. Will return true if _t_ matches the value of __peek__.
*/
func (p *Parser) match(t lexer.TokenType) bool {

	return p.peek().Type == t
}

/*
Consuming. Consumes the current token and increments source pointer iff __peek__ matches _t_.
*/
func (p *Parser) expect(t lexer.TokenType) (lexer.Token, error) {

	prev := p.peek()

	if &prev == nil || prev.Type != t {

		message := fmt.Sprintf("expecting token `%v`", t)
		return lexer.Token{}, fmt.Errorf("%v", message)
	}

	return p.consume(), nil
}

// Generates a formatted error message, complete with underline and error-point identification.
func (p *Parser) ErrorGenerator(message string) error {

	tmp := p.Tokens[p.Pos-1]

	m := utils.GenerateParserError(p.Audit[tmp.Line], tmp.Value, tmp.Line, tmp.Col, message)

	return fmt.Errorf("%v", m)
}

func ProduceAST(t []lexer.Token, a map[int]string) (ast.Program, error) {

	parser := Parser{

		Tokens: make([]lexer.Token, 0),
		Audit:  make(map[int]string, 0),
		Pos:    0,
	}

	// init.
	parser.Tokens = append(parser.Tokens, t...)
	parser.Audit = a

	program := ast.Program{
		Kind: "Program",
		Body: []ast.Expression{},
	}

	fmt.Printf("starting to parse tokens: %v\n", parser.Tokens)

	for parser.notEof() {

		parsed_statement, err := parser.parse_statement()
		if err != nil {
			return ast.Program{}, parser.ErrorGenerator(err.Error())
		}

		program.Body = append(program.Body, parsed_statement)
	}

	return program, nil
}

// Defines how the interpreter handles statements.
func (p *Parser) parse_statement() (ast.Expression, error) {

	switch p.peek().Type {
	case lexer.Let, lexer.Const:

		pvd, err := p.parse_var_decleration()
		if err != nil {
			return ast.Expr{}, err
		}

		return pvd, nil
	case lexer.Boolean:
		// Boolean literal coming in.
		b, err := p.parse_primary_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		return b, nil

	case lexer.Fn:

		fn, err := p.parse_fn_decleration()
		if err != nil {
			return ast.Expr{}, err
		}

		return fn, nil
	case lexer.If:

		iif, err := p.parse_if_condition()

		if err != nil {
			return ast.Expr{}, err
		}

		return iif, nil
	case lexer.While:

		while, err := p.parse_while_loop()
		if err != nil {
			return ast.Expr{}, err
		}

		return while, nil
	case lexer.For:

		while, err := p.parse_for_loop()
		if err != nil {
			return ast.Expr{}, err
		}

		return while, nil
	case lexer.Using:

		using, err := p.parse_using_decleration()
		if err != nil {
			return ast.Expr{}, err
		}

		return using, nil
	default:
		expr, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}
		return expr, nil
	}
}

// Defines how the interpreter handles experssions.
func (p *Parser) parse_expression() (ast.Expression, error) {

	assign, err := p.parse_assignment_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	return assign, nil
}

// Parses an if statement.
// Types of conditional checks we want to support:
// if (...) { ... }											// if						DONE.
// if (...) { ... } else { ... }							// if/else					DONE.
// if (...) { ... } elseif (...) { ... } else { ... }		// if/elseif/else
// let x = (...) ? { ... } : { ... }						// ternary operator			DONE.
func (p *Parser) parse_if_condition() (ast.Expression, error) {

	// Eat 'if' keyword
	p.consume()

	// Start of if condition, expect to see the open paren.
	_, err := p.expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Capture expression inside the parens.
	expr, err := p.parse_statement()
	if err != nil {
		return nil, err
	}

	var condition ast.Expression

	// Now we should get the type of the above expression.
	binop, isBinop := expr.(ast.BinaryExpr)
	if isBinop {
		condition = binop
	}

	boolean, isBool := expr.(ast.BooleanLiteral)
	if isBool {
		condition = boolean
	}

	iden, isIden := expr.(ast.Identifier)
	if isIden {
		condition = iden
	}

	// End of if condition, expect to see the close paren.
	_, err = p.expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Start of conditional body, expect to see the open brace.
	_, err = p.expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	// Until we hit the end of the if body.
	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		stmt, err := p.parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of conditional body, expect to see the closing brace.
	_, err = p.expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	// Set the 'if' block now, we can always amend later if this
	// is a if/else.
	var iif = ast.IfCondition{
		Kind:      "IfNode",
		Condition: condition,
		Body:      body,
		ElseCatch: false,
		ElseBody:  nil,
	}

	// Checking for an 'else' at the end of the 'if'.
	if p.match(lexer.Else) {

		// Eat past the 'else' keyword.
		p.consume()

		// Start of conditional body, expect to see the open brace.
		_, err = p.expect(lexer.OpenBrace)
		if err != nil {
			return nil, err
		}

		elseBody := make([]ast.Expression, 0)

		// Until we hit the end of the if body.
		for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

			stmt, err := p.parse_statement()
			if err != nil {
				return nil, err
			}

			elseBody = append(elseBody, stmt)
		}

		// End of conditional body, expect to see the closing brace.
		_, err = p.expect(lexer.CloseBrace)
		if err != nil {
			return nil, err
		}

		iif = ast.IfCondition{
			Kind:      "IfNode",
			Condition: condition,
			Body:      body,
			ElseCatch: true,
			ElseBody:  elseBody,
		}
	}

	return iif, nil
}

// Parses a standard while loop, i.e. while( ... ){ ... }
func (p *Parser) parse_while_loop() (ast.Expression, error) {

	p.consume() // Eat past the 'while' keyword.

	// Start of while loop, expect to see the open paren.
	_, err := p.expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Capture expression inside the parens.
	expr, err := p.parse_statement()
	if err != nil {
		return nil, err
	}

	var condition ast.Expression

	// Now we should get the type of the above expression.
	binop, isBinop := expr.(ast.BinaryExpr)
	if isBinop {
		condition = binop
	}

	boolean, isBool := expr.(ast.BooleanLiteral)
	if isBool {
		condition = boolean
	}

	iden, isIden := expr.(ast.Identifier)
	if isIden {
		condition = iden
	}

	// End of if condition, expect to see the close paren.
	_, err = p.expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Start of conditional body, expect to see the open brace.
	_, err = p.expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	// Until we hit the end of the if body.
	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		stmt, err := p.parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of conditional body, expect to see the closing brace.
	_, err = p.expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	return ast.WhileLoop{
		Kind:      ast.WhileNode,
		Condition: condition,
		Body:      body,
	}, nil
}

// Parses a standard for loop, i.e. for( ... ) { ... }
func (p *Parser) parse_for_loop() (ast.Expression, error) {

	p.consume() // Eat past the 'for' keyword.

	// Start of loop head, should be an open paren there.
	_, err := p.expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Next we should see an assignment expression, i.e. 'let i = 0;'
	ass, err := p.parse_var_decleration()
	if err != nil {
		return nil, err
	}

	varDec, isVarDecleration := ass.(ast.VariableDecleration)
	if !isVarDecleration {
		return nil, fmt.Errorf("invalid assigment statement provided: %v", ass)
	}

	// Next we expect to see our binary expression, as this is how we determine if the loop should keep running.
	expr, err := p.parse_statement()
	if err != nil {
		return nil, err
	}

	// Now we should get the type of the above expression.
	binop, isBinop := expr.(ast.BinaryExpr)
	if !isBinop {
		return nil, fmt.Errorf("invalid condition in loop: %v", binop)
	}

	// Next should be another ';'.
	_, err = p.expect(lexer.EOL)
	if err != nil {
		return nil, err
	}

	// Finally, we expect to see a shorthand operator expression.
	she, err := p.parse_identifier()
	if err != nil {
		return nil, err
	}

	shorthandOp, isShorthand := she.(ast.ShorthandOperator)
	if !isShorthand {
		return nil, fmt.Errorf("invalid shorthand operator provided: %v", shorthandOp)
	}

	// End of loop header, should see ')'.
	_, err = p.expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	// Start of loop body, expect to see '{'.
	_, err = p.expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Until we hit the end of the if body.
	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		stmt, err := p.parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// Start of loop body, expect to see '{'.
	_, err = p.expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	return ast.ForLoop{
		Kind:       "ForNode",
		Assignment: varDec,
		Condition:  binop,
		Iterator:   shorthandOp,
		Body:       body,
	}, nil
}

func (p *Parser) parse_assignment_expression() (ast.Expression, error) {

	left, err := p.parse_object_expression() // To be switched out with objects
	if err != nil {
		return ast.Expr{}, err
	}

	if p.match(lexer.Equals) {

		p.consume() // Advance past Equals token.

		value, err := p.parse_assignment_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		_, err = p.expect(lexer.EOL)
		if err != nil {
			return ast.Expr{}, err
		}

		return ast.AssignmentExpr{
			Kind:    "AssignmentExprNode",
			Assigne: left,
			Value:   value,
		}, nil
	} else if p.match(lexer.Ternary) {

		p.consume() // Advance past ternary op.

		// Now to capture the left expression.
		trueExpr, err := p.parse_expression()
		if err != nil {
			return nil, err
		}

		_, err = p.expect(lexer.Colon)
		if err != nil {
			return nil, err
		}

		// Capture rigth expression.
		falseExpr, err := p.parse_expression()
		if err != nil {
			return nil, err
		}

		return ast.TernaryCondition{
			Kind:      ast.TernaryNode,
			Condition: left,
			Left:      trueExpr,
			Right:     falseExpr,
		}, nil
	}

	return left, nil
}

// Parses a complex expression.
func (p *Parser) parse_object_expression() (ast.Expression, error) {

	// Non-map object.
	if !p.match(lexer.OpenBrace) {

		add, err := p.parse_additive_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		return add, nil
	}

	// Advances past '{'
	p.consume()

	props := make([]ast.Property, 0)

	// Continue reading unitl we get to the end of the object structure.
	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		key, err := p.expect(lexer.Identifier)
		if err != nil {
			return ast.Expr{}, err
		}

		// Allows short-hand definition, i.e.: { key, }
		if p.match(lexer.Comma) {

			p.consume() // Skip past ','.

			props = append(props, ast.Property{
				Key:  key.Value,
				Kind: "Property",
			})

			continue
		} else if p.match(lexer.CloseBrace) {

			// Allows short-hand definition, i.e.: { key }

			props = append(props, ast.Property{
				Key:  key.Value,
				Kind: "Property",
			})

			continue
		}

		_, err = p.expect(lexer.Colon)
		if err != nil {
			return ast.Expr{}, err
		}

		value, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		props = append(props, ast.Property{Kind: "Property", Value: &value, Key: key.Value})

		if !p.match(lexer.CloseBrace) {
			_, err = p.expect(lexer.Comma)
			if err != nil {
				return ast.Expr{}, err
			}
		}
	}

	_, err := p.expect(lexer.CloseBrace)
	if err != nil {
		return ast.Expr{}, err
	}

	return ast.ObjectLiteral{
		Kind:       "ObjectLiteral",
		Properties: props,
	}, nil
}

// Parses incoming functions.
func (p *Parser) parse_fn_decleration() (ast.Expression, error) {

	// Eats fn keyword
	p.consume()

	// Get the identifier name of the function.
	fnName, err := p.expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}

	// Args of the function.
	args, err := p.parse_args()
	if err != nil {
		return nil, err
	}

	params := make([]string, 0)

	for _, arg := range args {

		i, isIden := arg.(ast.Identifier)
		if !isIden {
			return nil, fmt.Errorf("expected parameters to be of type string")
		}

		// Push symbol identifier into params list.
		params = append(params, i.Symbol)
	}

	// Expect '{' at start of function body.
	_, err = p.expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Until we hit the end of the funciton body.
	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		stmt, err := p.parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of function, expect to see the closing brace.
	_, err = p.expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	function := ast.FunctionDecleration{
		Kind:   "FunctionDeclerationNode",
		Name:   fnName.Value,
		Params: params,
		Body:   body,
	}

	return function, nil
}

/*
Handles either:
  - lex x = 10;
  - let x;
  - const y = 9;
*/
func (p *Parser) parse_var_decleration() (ast.Expression, error) {

	// true:  const x = 10;
	// false: let x = 10;
	isConst := p.consume().Type == lexer.Const
	identifier, err := p.expect(lexer.Identifier)
	if err != nil {
		return ast.Expr{}, err
	}

	if p.match(lexer.EOL) {

		// Consume the next token.
		p.consume()

		if isConst {
			// Current token is an EOL however trying to define const. Error.
			return ast.Expr{}, p.ErrorGenerator("no value provided for const decleration")
		}

		// E.g. 'let x;'
		return ast.VariableDecleration{
			Kind:       "VariableDeclerationNode",
			Constant:   isConst,
			Identifier: identifier.Value,
			Value:      ast.Expr{},
		}, nil
	}

	// Now we are checking 'let x = 10;'
	_, err = p.expect(lexer.Equals)
	if err != nil {
		return ast.Expr{}, err
	}

	// In the case of 'let x = [];', an array is being declared.
	if p.match(lexer.OpenBracket) {

		// Eat the opening bracket.
		p.consume()

		// Attempt to capture all the expressions inside the array.
		array_decleration, err := p.parse_array_decleration(identifier.Value, isConst)
		if err != nil {
			return ast.Expr{}, err
		}

		return array_decleration, nil

	} else if p.match(lexer.OpenBrace) {

		// In the case of 'let x = {};', a map is being declared.

		// Eat the opening brace.
		p.consume()

		// Attempt to capture all the expressions inside the array.
		map_decleration, err := p.parse_map_decleration(identifier.Value, isConst)
		if err != nil {
			return ast.Expr{}, err
		}

		return map_decleration, nil
	}

	// Standard variable decleration, i.e. 'let x = 10;'

	value, err := p.parse_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	decleration := ast.VariableDecleration{
		Kind:       "VariableDeclerationNode",
		Value:      value,
		Identifier: identifier.Value,
		Constant:   isConst,
	}

	_, isCallExpr := decleration.Value.(ast.CallExpr)
	if isCallExpr {
		return decleration, nil
	}

	_, err = p.expect(lexer.EOL)
	if err != nil {
		return ast.Expr{}, err
	}

	return decleration, nil
}

// Parses a statement that declares a new map.
func (p *Parser) parse_map_decleration(identifier string, isConst bool) (ast.Expression, error) {

	keyValuePairs := make(map[ast.Expression]ast.Expression, 0)

	for !p.match(lexer.CloseBrace) && !p.match(lexer.EOF) {

		// Capture the key defined.
		key, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		// Key must be of type IComparable
		if !isComparableType(key) {
			return nil, fmt.Errorf("invalid type provided for map key: %v", key)
		}

		// Next we expect to see a ':'.
		_, err = p.expect(lexer.Colon)
		if err != nil {
			return nil, err
		}

		// Capture the value defined.
		value, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		// Need to make sure the keys are unique.
		if _, ok := keyValuePairs[key]; ok {
			return nil, fmt.Errorf("maps keys should be unique: %v", key)
		}

		// Store the new key/value pair.
		keyValuePairs[key] = value

		// Next we expect to see a ','.
		_, err = p.expect(lexer.Comma)
		if err != nil {
			return nil, err
		}
	}

	// End of map body, expect to see a closing bracket.
	_, err := p.expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	// End of array decleration, expect to see an EOL.
	_, err = p.expect(lexer.EOL)
	if err != nil {
		return nil, err
	}

	return ast.MapDecleration{
		Kind:       "MapDeclerationNode",
		Identifier: identifier,
		Value:      keyValuePairs,
		Constant:   isConst,
	}, nil
}

// Is the key type one of the valid types Goblin allows for its keys?
func isComparableType(val any) bool {

	switch any(val).(type) {
	case ast.NumericLiteral, ast.StringLiteral, ast.BooleanLiteral:
		return true
	default:
		return false
	}
}

// Parses a statement that declares a new array.
func (p *Parser) parse_array_decleration(identifier string, isConst bool) (ast.Expression, error) {

	expressions := make([]ast.Expression, 0)

	for p.match(lexer.CloseBracket) && !p.match(lexer.EOF) {

		value, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		expressions = append(expressions, value)

		if p.match(lexer.CloseBracket) {
			break
		} else {
			_, err = p.expect(lexer.Comma)
			if err != nil {
				return nil, err
			}
		}
	}

	// End of array body, expect to see a closing bracket.
	_, err := p.expect(lexer.CloseBracket)
	if err != nil {
		return nil, err
	}

	// End of array decleration, expect to see an EOL.
	_, err = p.expect(lexer.EOL)
	if err != nil {
		return nil, err
	}

	decleration := ast.ArrayDecleration{
		Kind:       "ArrayDeclerationNode",
		Value:      expressions,
		Identifier: identifier,
		Constant:   isConst,
	}

	return decleration, nil
}

// Defines how the interpreter handles additive expressions.
func (p *Parser) parse_additive_expression() (ast.Expression, error) {

	left, err := p.parse_multiplicitive_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for p.peek().Value == "+" || p.peek().Value == "-" || p.peek().Value == "<" || p.peek().Value == ">" || p.peek().Value == "==" || p.peek().Value == "!=" {

		operator := p.consume().Value

		right, err := p.parse_multiplicitive_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		left = ast.BinaryExpr{
			Kind:     "BinaryExprNode",
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}

	return left, nil
}

func (p *Parser) parse_call_member_expression() (ast.Expression, error) {

	member, err := p.parse_member_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	// '(' found, go into a call expression.
	if p.match(lexer.OpenParen) {

		val, err := p.parse_call_expression(member)
		if err != nil {
			return ast.Expr{}, err
		}

		return val, err
	}

	return member, nil
}

func (p *Parser) parse_call_expression(caller ast.Expression) (ast.CallExpr, error) {

	args, err := p.parse_args()
	if err != nil {
		return ast.CallExpr{}, err
	}

	call_expr := ast.CallExpr{
		Kind:   "CallExpression",
		Caller: caller,
		Args:   args,
	}

	// At another '('.
	if p.match(lexer.OpenParen) {

		call_expr, err = p.parse_call_expression(call_expr)
		if err != nil {
			return ast.CallExpr{}, err
		}
	} else {
		_, err := p.expect(lexer.EOL)
		if err != nil {
			return ast.CallExpr{}, err
		}
	}

	return call_expr, nil
}

func (p *Parser) parse_args() ([]ast.Expression, error) {

	_, err := p.expect(lexer.OpenParen)
	if err != nil {
		return []ast.Expression{}, err
	}

	var args []ast.Expression

	if p.match(lexer.CloseParen) {
		// Return an empty array.
		args = []ast.Expression{}
	} else {
		argsList, err := p.parse_args_list()
		if err != nil {
			return []ast.Expression{}, err
		}

		args = argsList
	}

	_, err = p.expect(lexer.CloseParen)
	if err != nil {
		return []ast.Expression{}, err
	}

	return args, nil
}

// Handles the following, e.g. foo(x = 5, v = "bar")
func (p *Parser) parse_args_list() ([]ast.Expression, error) {

	args := make([]ast.Expression, 0)

	arg1, err := p.parse_assignment_expression()
	if err != nil {
		return []ast.Expression{}, err
	}

	args = append(args, arg1)

	for p.match(lexer.Comma) && (p.consume() != lexer.Token{}) {

		expr, err := p.parse_assignment_expression()
		if err != nil {
			return []ast.Expression{}, err
		}

		args = append(args, expr)
	}

	return args, nil
}

// Parses a 'using' directive.
func (p *Parser) parse_using_decleration() (ast.Expression, error) {

	// Move past the 'using' keyword.
	p.consume()

	val, err := p.parse_statement()
	if err != nil {
		return nil, err
	}

	str, ok := val.(ast.StringLiteral)
	if !ok {
		return nil, p.ErrorGenerator("string type required with using directives")
	}

	// Always expect to see a ';' after a using directive.
	_, err = p.expect(lexer.EOL)
	if err != nil {
		return nil, err
	}

	return ast.NamespaceDecleration{
		Kind: "NamespaceDecleration",
		Name: str.Value,
	}, nil
}

// Parses how to access member fields from an object.
func (p *Parser) parse_member_expression() (ast.Expression, error) {

	object, err := p.parse_primary_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for p.match(lexer.Period) || p.match(lexer.OpenBracket) {

		// Gives us access to current operator, either '.' or '['
		opp := p.consume()
		var computed bool

		// Get the Identifier.
		prop, err := p.parse_primary_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		if opp.Type == lexer.Period {

			// Not a computed expression.

			computed = false

			_, isIden := prop.(ast.Identifier)
			if !isIden {
				return ast.Expr{}, fmt.Errorf("cannot use dot operator without rhs being an indentifier")
			}

		} else {
			// This should allow us to do chaining.
			computed = true

			prop, err = p.parse_expression()
			if err != nil {
				return ast.Expr{}, err
			}

			_, err = p.expect(lexer.CloseBracket)
			if err != nil {
				return ast.Expr{}, err
			}
		}

		object = ast.MemberExpr{
			Kind:     "MemberExpressionNode",
			Object:   object,
			Property: prop,
			Computed: computed,
		}
	}

	return object, nil
}

// Defines how the interpreter handles multiplicitive expressions.
func (p *Parser) parse_multiplicitive_expression() (ast.Expression, error) {

	left, err := p.parse_call_member_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for p.peek().Value == "*" || p.peek().Value == "/" || p.peek().Value == "%" {

		operator := p.consume().Value

		right, err := p.parse_call_member_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		left = ast.BinaryExpr{
			Kind:     "BinaryExprNode",
			Left:     left,
			Right:    right,
			Operator: operator,
		}
	}

	return left, nil
}

func (p *Parser) parse_identifier() (ast.Expression, error) {

	// Normal identifier, or array identifier?
	// Normal -> x
	// Array -> x[0]
	// Map -> x["foo"]
	// Shorthand Operator -> x++ or x--

	identifier := p.consume() // Capture the identifier value

	if p.match(lexer.OpenBracket) {
		p.consume() // Eat the open bracket.

		// Capture index, but we need to parse it as it could be a number or an identifier.
		index, err := p.parse_expression()
		if err != nil {
			return nil, err
		}

		// End of array/map body, expect to see a closing bracket.
		_, err = p.expect(lexer.CloseBracket)
		if err != nil {
			return nil, err
		}

		return ast.ArrayOrMapIdentifier{
			Kind:   "ArrayOrMapIdentifierNode",
			Symbol: identifier.Value,
			Index:  index,
		}, nil
	} else if p.match(lexer.ShorthandOperator) {

		// Capture the operator type.
		opp := p.consume()

		// Depending on shorthand operator used:
		// x++;
		// x += 1;
		// Need to handle accordingly.

		if opp.Value == "++" || opp.Value == "--" {

			// ++, --

			// End of statement.
			_, err := p.expect(lexer.EOL)
			if err != nil {
				return nil, err
			}

			return ast.ShorthandOperator{
				Kind: "ShorthandOperatorNode",
				Left: identifier.Value,
				Right: ast.NumericLiteral{
					Kind:  "NumberNode",
					Value: 0,
				},
				Operator: opp.Value,
			}, nil

		} else {

			rhs, err := p.parse_expression()
			if err != nil {
				return ast.Expr{}, nil
			}

			// End of statement.
			_, err = p.expect(lexer.EOL)
			if err != nil {
				return nil, err
			}

			return ast.ShorthandOperator{
				Kind:     "ShorthandOperatorNode",
				Left:     identifier.Value,
				Right:    rhs,
				Operator: opp.Value,
			}, nil
		}
	} else {
		// Standard identifier.
		return ast.Identifier{
			Kind:   "IdentifierNode",
			Symbol: identifier.Value,
		}, nil
	}
}

// Defines how the interpreter handles primary expressions.
func (p *Parser) parse_primary_expression() (ast.Expression, error) {

	tk := p.peek().Type

	switch tk {
	case lexer.Identifier:
		// Some form of Identifier coming in.
		iden, err := p.parse_identifier()
		if err != nil {
			return ast.Expr{}, nil
		}

		return iden, nil
	case lexer.Boolean:
		return ast.BooleanLiteral{
			Kind:  "BooleanLiteralNode",
			Value: utils.StoB(p.consume().Value),
		}, nil
	case lexer.String:
		return ast.StringLiteral{
			Kind:  "StringLiteralNode",
			Value: p.consume().Value,
		}, nil
	case lexer.Number:
		// Convert the tokens string value into a int.
		val, err := utils.ToNumber(p.consume().Value)
		if err != nil {
			return ast.Expr{}, err
		}

		return ast.NumericLiteral{
			Kind:  "NumericLiteralNode",
			Value: val,
		}, nil

	case lexer.OpenParen:
		p.consume() // Consume to remove.
		v, err := p.parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}
		value := v
		_, err = p.expect(lexer.CloseParen) // Consume to remove.
		if err != nil {
			return ast.Expr{}, err
		}

		return value, nil

	default:
		message := fmt.Sprintf("unexpected token found during parsing '%v'", p.peek().Value)
		return ast.Expr{}, fmt.Errorf("%v", message)
	}
}

// Checks to see if we have hit the end of the file.
func (p *Parser) notEof() bool {
	return p.Tokens[p.Pos].Type != lexer.EOF
}
