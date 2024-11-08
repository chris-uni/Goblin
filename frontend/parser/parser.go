package parser

import (
	"fmt"

	"goblin.org/main/frontend/ast"
	"goblin.org/main/frontend/lexer"
	"goblin.org/main/utils"
)

/*
Orders of prescidence

// Assignment
// Object
// AdditiveExpr
// MultiplicitaveExpr
// Call
// Member
// PrimaryExpr
*/

var tokens = make([]lexer.Token, 0)

// Simple returns the current token.
func at() lexer.Token {
	return tokens[0]
}

// Returns the current token and shifts the pointer along to
// the next in the list.
func eat() lexer.Token {

	prev := utils.Shift[lexer.Token](&tokens)
	return prev
}

// Returns the current token and shifts the pointer along to
// the next in the list. Used in error handling to check the expected
// type of the token about to be returned.
func expect(t lexer.TokenType) (lexer.Token, error) {

	prev := utils.Shift[lexer.Token](&tokens)

	if &prev == nil || prev.Type != t {
		return lexer.Token{}, fmt.Errorf("expecting type '%v'", t)
	}

	return prev, nil
}

func ProduceAST(source string) (ast.Program, error) {

	// Convert source code into tokens.
	tokens = lexer.Tokenize(source)

	// fmt.Printf("Tokens: %v\n", tokens)

	program := ast.Program{
		Kind: "Program",
		Body: []ast.Expression{},
	}

	for notEof() {

		parsed_statement, err := parse_statement()
		if err != nil {
			return ast.Program{}, err
		}

		program.Body = append(program.Body, parsed_statement)
	}

	return program, nil
}

// Defines how the interpreter handles statements.
func parse_statement() (ast.Expression, error) {

	switch at().Type {
	case lexer.Let, lexer.Const:

		pvd, err := parse_var_decleration()
		if err != nil {
			return ast.Expr{}, err
		}

		return pvd, nil
	case lexer.Boolean:
		// Boolean literal coming in.
		b, err := parse_primary_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		return b, nil

	case lexer.Fn:

		fn, err := parse_fn_decleration()

		if err != nil {
			return ast.Expr{}, err
		}

		return fn, nil
	case lexer.If:

		iif, err := parse_if_condition()

		if err != nil {
			return ast.Expr{}, err
		}

		return iif, nil
	case lexer.While:

		while, err := parse_while_loop()
		if err != nil {
			return ast.Expr{}, err
		}

		return while, nil
	case lexer.For:

		while, err := parse_for_loop()
		if err != nil {
			return ast.Expr{}, err
		}

		return while, nil
	default:
		expr, err := parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}
		return expr, nil
	}
}

// Defines how the interpreter handles experssions.
func parse_expression() (ast.Expression, error) {

	assign, err := parse_assignment_expression()
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
func parse_if_condition() (ast.Expression, error) {

	// Eat 'if' keyword
	eat()

	// Start of if condition, expect to see the open paren.
	_, err := expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Capture expression inside the parens.
	expr, err := parse_statement()
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
	_, err = expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Start of conditional body, expect to see the open brace.
	_, err = expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	// Until we hit the end of the if body.
	for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

		stmt, err := parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of conditional body, expect to see the closing brace.
	_, err = expect(lexer.CloseBrace)
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
	if at().Type == lexer.Else {

		// Eat past the 'else' keyword.
		eat()

		// Start of conditional body, expect to see the open brace.
		_, err = expect(lexer.OpenBrace)
		if err != nil {
			return nil, err
		}

		elseBody := make([]ast.Expression, 0)

		// Until we hit the end of the if body.
		for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

			stmt, err := parse_statement()
			if err != nil {
				return nil, err
			}

			elseBody = append(elseBody, stmt)
		}

		// End of conditional body, expect to see the closing brace.
		_, err = expect(lexer.CloseBrace)
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
func parse_while_loop() (ast.Expression, error) {

	eat() // Eat past the 'while' keyword.

	// Start of while loop, expect to see the open paren.
	_, err := expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Capture expression inside the parens.
	expr, err := parse_statement()
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
	_, err = expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Start of conditional body, expect to see the open brace.
	_, err = expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	// Until we hit the end of the if body.
	for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

		stmt, err := parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of conditional body, expect to see the closing brace.
	_, err = expect(lexer.CloseBrace)
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
func parse_for_loop() (ast.Expression, error) {

	eat() // Eat past the 'for' keyword.

	// Start of loop head, should be an open paren there.
	_, err := expect(lexer.OpenParen)
	if err != nil {
		return nil, err
	}

	// Next we should see an assignment expression, i.e. 'let i = 0;'
	ass, err := parse_var_decleration()
	if err != nil {
		return nil, err
	}

	varDec, isVarDecleration := ass.(ast.VariableDecleration)
	if !isVarDecleration {
		return nil, fmt.Errorf("invalid assigment statement provided: %v", ass)
	}

	// Next we expect to see our binary expression, as this is how we determine if the loop should keep running.
	expr, err := parse_statement()
	if err != nil {
		return nil, err
	}

	// Now we should get the type of the above expression.
	binop, isBinop := expr.(ast.BinaryExpr)
	if !isBinop {
		return nil, fmt.Errorf("invalid condition in loop: %v", binop)
	}

	// Next should be another ';'.
	_, err = expect(lexer.EOL)
	if err != nil {
		return nil, err
	}

	// Finally, we expect to see a shorthand operator expression.
	she, err := parse_identifier()
	if err != nil {
		return nil, err
	}

	shorthandOp, isShorthand := she.(ast.ShorthandOperator)
	if !isShorthand {
		return nil, fmt.Errorf("invalid shorthand operator provided: %v", shorthandOp)
	}

	// End of loop header, should see ')'.
	_, err = expect(lexer.CloseParen)
	if err != nil {
		return nil, err
	}

	// Start of loop body, expect to see '{'.
	_, err = expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Until we hit the end of the if body.
	for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

		stmt, err := parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// Start of loop body, expect to see '{'.
	_, err = expect(lexer.CloseBrace)
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

func parse_assignment_expression() (ast.Expression, error) {

	left, err := parse_object_expression() // To be switched out with objects
	if err != nil {
		return ast.Expr{}, err
	}

	if at().Type == lexer.Equals {

		eat() // Advance past Equals token.

		value, err := parse_assignment_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		_, err = expect(lexer.EOL)
		if err != nil {
			return ast.Expr{}, err
		}

		return ast.AssignmentExpr{
			Kind:    "AssignmentExprNode",
			Assigne: left,
			Value:   value,
		}, nil
	} else if at().Type == lexer.Ternary {

		eat() // Advance past ternary op.

		// Now to capture the left expression.
		trueExpr, err := parse_expression()
		if err != nil {
			return nil, err
		}

		_, err = expect(lexer.Colon)
		if err != nil {
			return nil, err
		}

		// Capture rigth expression.
		falseExpr, err := parse_expression()
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
func parse_object_expression() (ast.Expression, error) {

	// Non-map object.
	if at().Type != lexer.OpenBrace {

		add, err := parse_additive_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		return add, nil
	}

	// Advances past '{'
	eat()

	props := make([]ast.Property, 0)

	// Continue reading unitl we get to the end of the object structure.
	for notEof() && at().Type != lexer.CloseBrace {

		key, err := expect(lexer.Identifier)
		if err != nil {
			return ast.Expr{}, err
		}

		// Allows short-hand definition, i.e.: { key, }
		if at().Type == lexer.Comma {

			eat() // Skip past ','.

			props = append(props, ast.Property{
				Key:  key.Value,
				Kind: "Property",
			})

			continue
		} else if at().Type == lexer.CloseBrace {

			// Allows short-hand definition, i.e.: { key }

			props = append(props, ast.Property{
				Key:  key.Value,
				Kind: "Property",
			})

			continue
		}

		_, err = expect(lexer.Colon)
		if err != nil {
			return ast.Expr{}, err
		}

		value, err := parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		props = append(props, ast.Property{Kind: "Property", Value: &value, Key: key.Value})

		if at().Type != lexer.CloseBrace {
			_, err = expect(lexer.Comma)
			if err != nil {
				return ast.Expr{}, err
			}
		}
	}

	_, err := expect(lexer.CloseBrace)
	if err != nil {
		return ast.Expr{}, err
	}

	return ast.ObjectLiteral{
		Kind:       "ObjectLiteral",
		Properties: props,
	}, nil
}

// Parses incoming functions.
func parse_fn_decleration() (ast.Expression, error) {

	// Eats fn keyword
	eat()

	// Get the identifier name of the function.
	fnName, err := expect(lexer.Identifier)
	if err != nil {
		return nil, err
	}

	// Args of the function.
	args, err := parse_args()
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
	_, err = expect(lexer.OpenBrace)
	if err != nil {
		return nil, err
	}

	body := make([]ast.Expression, 0)

	// Until we hit the end of the funciton body.
	for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

		stmt, err := parse_statement()
		if err != nil {
			return nil, err
		}

		body = append(body, stmt)
	}

	// End of function, expect to see the closing brace.
	_, err = expect(lexer.CloseBrace)
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
func parse_var_decleration() (ast.Expression, error) {

	// true:  const x = 10;
	// false: let x = 10;
	isConst := eat().Type == lexer.Const
	identifier, err := expect(lexer.Identifier)
	if err != nil {
		return ast.Expr{}, err
	}

	if at().Type == lexer.EOL {

		// Consume the next token.
		eat()

		if isConst {
			// Current token is an EOL however trying to define const. Error.
			return ast.Expr{}, fmt.Errorf("no value provided for const decleration")
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
	_, err = expect(lexer.Equals)
	if err != nil {
		return ast.Expr{}, err
	}

	// In the case of 'let x = [];', an array is being declared.
	if at().Type == lexer.OpenBracket {

		// Eat the opening bracket.
		eat()

		// Attempt to capture all the expressions inside the array.
		array_decleration, err := parse_array_decleration(identifier.Value, isConst)
		if err != nil {
			return ast.Expr{}, err
		}

		return array_decleration, nil

	} else if at().Type == lexer.OpenBrace {

		// In the case of 'let x = {};', a map is being declared.

		// Eat the opening brace.
		eat()

		// Attempt to capture all the expressions inside the array.
		map_decleration, err := parse_map_decleration(identifier.Value, isConst)
		if err != nil {
			return ast.Expr{}, err
		}

		return map_decleration, nil
	}

	// Standard variable decleration, i.e. 'let x = 10;'

	value, err := parse_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	decleration := ast.VariableDecleration{
		Kind:       "VariableDeclerationNode",
		Value:      value,
		Identifier: identifier.Value,
		Constant:   isConst,
	}

	_, err = expect(lexer.EOL)
	if err != nil {
		return ast.Expr{}, err
	}

	return decleration, nil
}

// Parses a statement that declares a new map.
func parse_map_decleration(identifier string, isConst bool) (ast.Expression, error) {

	keyValuePairs := make(map[ast.Expression]ast.Expression, 0)

	for at().Type != lexer.CloseBrace && at().Type != lexer.EOF {

		// Capture the key defined.
		key, err := parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		// Key must be of type IComparable
		if !isComparableType(key) {
			return nil, fmt.Errorf("invalid type provided for map key: %v", key)
		}

		// Next we expect to see a ':'.
		_, err = expect(lexer.Colon)
		if err != nil {
			return nil, err
		}

		// Capture the value defined.
		value, err := parse_expression()
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
		_, err = expect(lexer.Comma)
		if err != nil {
			return nil, err
		}
	}

	// End of map body, expect to see a closing bracket.
	_, err := expect(lexer.CloseBrace)
	if err != nil {
		return nil, err
	}

	// End of array decleration, expect to see an EOL.
	_, err = expect(lexer.EOL)
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
func parse_array_decleration(identifier string, isConst bool) (ast.Expression, error) {

	expressions := make([]ast.Expression, 0)

	for at().Type != lexer.CloseBracket && at().Type != lexer.EOF {

		value, err := parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}

		expressions = append(expressions, value)

		if at().Type == lexer.CloseBracket {
			break
		} else {
			_, err = expect(lexer.Comma)
			if err != nil {
				return nil, err
			}
		}
	}

	// End of array body, expect to see a closing bracket.
	_, err := expect(lexer.CloseBracket)
	if err != nil {
		return nil, err
	}

	// End of array decleration, expect to see an EOL.
	_, err = expect(lexer.EOL)
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
func parse_additive_expression() (ast.Expression, error) {

	left, err := parse_multiplicitive_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for at().Value == "+" || at().Value == "-" || at().Value == "<" || at().Value == ">" || at().Value == "==" {

		operator := eat().Value

		right, err := parse_multiplicitive_expression()
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

func parse_call_member_expression() (ast.Expression, error) {

	member, err := parse_member_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	// '(' found, go into a call expression.
	if at().Type == lexer.OpenParen {

		val, err := parse_call_expression(member)
		if err != nil {
			return ast.Expr{}, err
		}

		return val, err
	}

	return member, nil
}

func parse_call_expression(caller ast.Expression) (ast.CallExpr, error) {

	args, err := parse_args()
	if err != nil {
		return ast.CallExpr{}, err
	}

	call_expr := ast.CallExpr{
		Kind:   "CallExpression",
		Caller: caller,
		Args:   args,
	}

	// At another '('.
	if at().Type == lexer.OpenParen {

		call_expr, err = parse_call_expression(call_expr)
		if err != nil {
			return ast.CallExpr{}, err
		}
	} else {
		_, err := expect(lexer.EOL)
		if err != nil {
			return ast.CallExpr{}, err
		}
	}

	return call_expr, nil
}

func parse_args() ([]ast.Expression, error) {

	_, err := expect(lexer.OpenParen)
	if err != nil {
		return []ast.Expression{}, err
	}

	var args []ast.Expression

	if at().Type == lexer.CloseParen {
		// Return an empty array.
		args = []ast.Expression{}
	} else {
		argsList, err := parse_args_list()
		if err != nil {
			return []ast.Expression{}, err
		}

		args = argsList
	}

	_, err = expect(lexer.CloseParen)
	if err != nil {
		return []ast.Expression{}, err
	}

	return args, nil
}

// Handles the following, e.g. foo(x = 5, v = "bar")
func parse_args_list() ([]ast.Expression, error) {

	args := make([]ast.Expression, 0)

	arg1, err := parse_assignment_expression()
	if err != nil {
		return []ast.Expression{}, err
	}

	args = append(args, arg1)

	for at().Type == lexer.Comma && (eat() != lexer.Token{}) {

		expr, err := parse_assignment_expression()
		if err != nil {
			return []ast.Expression{}, err
		}

		args = append(args, expr)
	}

	return args, nil
}

// Parses how to access member fields from an object.
func parse_member_expression() (ast.Expression, error) {

	object, err := parse_primary_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for at().Type == lexer.Period || at().Type == lexer.OpenBracket {

		// Gives us access to current operator, either '.' or '('
		opp := eat()
		var computed bool

		// Get the Identifier.
		prop, err := parse_primary_expression()
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

			prop, err = parse_expression()
			if err != nil {
				return ast.Expr{}, err
			}

			_, err = expect(lexer.CloseBracket)
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
func parse_multiplicitive_expression() (ast.Expression, error) {

	left, err := parse_call_member_expression()
	if err != nil {
		return ast.Expr{}, err
	}

	for at().Value == "*" || at().Value == "/" || at().Value == "%" {

		operator := eat().Value

		right, err := parse_call_member_expression()
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

func parse_identifier() (ast.Expression, error) {

	// Normal identifier, or array identifier?
	// Normal -> x
	// Array -> x[0]
	// Map -> x["foo"]
	// Shorthand Operator -> x++ or x--

	identifier := eat() // Capture the identifier value

	if at().Type == lexer.OpenBracket {
		eat() // Eat the open bracket.

		// Capture index, but we need to parse it as it could be a number or an identifier.
		index, err := parse_expression()
		if err != nil {
			return nil, err
		}

		// End of array/map body, expect to see a closing bracket.
		_, err = expect(lexer.CloseBracket)
		if err != nil {
			return nil, err
		}

		return ast.ArrayOrMapIdentifier{
			Kind:   "ArrayOrMapIdentifierNode",
			Symbol: identifier.Value,
			Index:  index,
		}, nil
	} else if at().Type == lexer.ShorthandOperator {

		// Capture the operator type.
		opp := eat()

		// Depending on shorthand operator used:
		// x++;
		// x += 1;
		// Need to handle accordingly.

		if opp.Value == "++" || opp.Value == "--" {

			// ++, --

			// End of statement.
			_, err := expect(lexer.EOL)
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

			rhs, err := parse_expression()
			if err != nil {
				return ast.Expr{}, nil
			}

			// End of statement.
			_, err = expect(lexer.EOL)
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
func parse_primary_expression() (ast.Expression, error) {

	tk := at().Type

	switch tk {
	case lexer.Identifier:
		// Some form of Identifier coming in.
		iden, err := parse_identifier()
		if err != nil {
			return ast.Expr{}, nil
		}

		return iden, nil
	case lexer.Boolean:
		return ast.BooleanLiteral{
			Kind:  "BooleanLiteralNode",
			Value: utils.StoB(eat().Value),
		}, nil
	case lexer.String:
		return ast.StringLiteral{
			Kind:  "StringLiteralNode",
			Value: eat().Value,
		}, nil
	case lexer.Number:
		// Convert the tokens string value into a int.
		val, err := utils.ToNumber(eat().Value)
		if err != nil {
			return ast.Expr{}, err
		}

		return ast.NumericLiteral{
			Kind:  "NumericLiteralNode",
			Value: val,
		}, nil

	case lexer.OpenParen:
		eat() // Consume to remove.
		v, err := parse_expression()
		if err != nil {
			return ast.Expr{}, err
		}
		value := v
		_, err = expect(lexer.CloseParen) // Consume to remove.
		if err != nil {
			return ast.Expr{}, err
		}

		return value, nil

	default:
		return ast.Expr{}, fmt.Errorf("unexpected token found during parsing - %v", at().ToString())
	}
}

// Checks to see if we have hit the end of the file.
func notEof() bool {
	return tokens[0].Type != lexer.EOF
}
