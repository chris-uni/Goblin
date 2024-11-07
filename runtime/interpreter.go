package runtime

import (
	"fmt"

	"goblin.org/main/frontend/ast"
)

func Evaluate(astNode ast.Expression, env Environment) (RuntimeValue, error) {

	// Check the type of the expression coming in for resolution later on.
	if value, ok := astNode.(ast.NumericLiteral); ok {

		return MK_NUMBER(value.Value), nil

	} else if sho, ok := astNode.(ast.ShorthandOperator); ok {

		shoVal, err := eval_shorthand_operator_expression(sho, env)
		if err != nil {
			return nil, err
		}

		return shoVal, err

	} else if bino, ok := astNode.(ast.BinaryExpr); ok {

		binop, err := eval_binary_expression(bino, env)
		if err != nil {
			return nil, err
		}

		return binop, err

	} else if w, ok := astNode.(ast.WhileLoop); ok {

		while, err := eval_while_expression(w, env)
		if err != nil {
			return nil, err
		}

		return while, err

	} else if str, ok := astNode.(ast.StringLiteral); ok {

		str, err := eval_string_expression(str, env)
		if err != nil {
			return nil, err
		}

		return str, err

	} else if b, ok := astNode.(ast.BooleanLiteral); ok {

		boolean := BooleanValue{
			Type:  Boolean,
			Value: b.Value,
		}

		return boolean, nil

	} else if iif, ok := astNode.(ast.IfCondition); ok {

		ifCondition, err := eval_if_condition_expression(iif, env)
		if err != nil {
			return nil, err
		}

		return ifCondition, err

	} else if t, ok := astNode.(ast.TernaryCondition); ok {

		ternary, err := eval_ternary_expression(t, env)
		if err != nil {
			return nil, err
		}

		return ternary, err

	} else if iden, ok := astNode.(ast.Identifier); ok {

		iden, err := eval_identifier(iden, env)
		if err != nil {
			return nil, err
		}

		return iden, nil

	} else if aomIden, ok := astNode.(ast.ArrayOrMapIdentifier); ok {

		aom, err := eval_array_or_map_identifier(aomIden, env)
		if err != nil {
			return nil, err
		}

		return aom, nil

	} else if object, ok := astNode.(ast.ObjectLiteral); ok {

		obj, err := eval_object_expr(object, env)
		if err != nil {
			return nil, err
		}

		return obj, nil

	} else if call, ok := astNode.(ast.CallExpr); ok {

		call, err := eval_call_expr(call, env)
		if err != nil {
			return nil, err
		}

		return call, nil

	} else if program, ok := astNode.(ast.Program); ok {

		prog, err := eval_program(program, env)
		if err != nil {
			return nil, err
		}

		return prog, nil

	} else if dec, ok := astNode.(ast.VariableDecleration); ok {

		varDec, err := eval_var_decleration(dec, env)
		if err != nil {
			return nil, err
		}

		return varDec, nil

	} else if arr, ok := astNode.(ast.ArrayDecleration); ok {

		arrDec, err := eval_arr_decleration(arr, env)
		if err != nil {
			return nil, err
		}

		return arrDec, nil

	} else if map_, ok := astNode.(ast.MapDecleration); ok {

		mapDec, err := eval_map_decleration(map_, env)
		if err != nil {
			return nil, err
		}

		return mapDec, nil

	} else if func_, ok := astNode.(ast.FunctionDecleration); ok {

		fn, err := eval_function_decleration(func_, env)
		if err != nil {
			return nil, err
		}

		return fn, nil

	} else if assign, ok := astNode.(ast.AssignmentExpr); ok {

		assign, err := eval_assignment_expression(assign, env)
		if err != nil {
			return nil, err
		}

		return assign, nil
	}

	return nil, fmt.Errorf("unrecognised node in source %v", astNode)
}

// Interpreter entry. Evaluates an entire program.
func eval_program(prog ast.Program, env Environment) (RuntimeValue, error) {

	var lastEval RuntimeValue

	for _, expr := range prog.Body {

		tmp, err := Evaluate(expr, env)
		if err != nil {
			return nil, err
		}

		lastEval = tmp
	}

	return lastEval, nil
}

// Evaluates the provided identifier.
func eval_identifier(iden ast.Identifier, env Environment) (RuntimeValue, error) {

	val, err := env.Lookup(iden.Symbol)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// Evaluates either an array or map index accessor.
func eval_array_or_map_identifier(arr ast.ArrayOrMapIdentifier, env Environment) (RuntimeValue, error) {

	i, err := Evaluate(arr.Index, env)
	if err != nil {
		return nil, err
	}

	ds, err := env.ArrayOrMapLookup(arr.Symbol, i)
	if err != nil {
		return nil, err
	}

	return ds, nil
}

// Evaluates complex object assignments such as 'let foo = {x: 10};'
func eval_object_expr(obj ast.ObjectLiteral, env Environment) (RuntimeValue, error) {

	object := ObjectVal{Type: "Object", Properties: map[string]RuntimeValue{}}

	for _, prop := range obj.Properties {

		key := prop.Key
		value := prop.Value

		var runtimeVal RuntimeValue

		if value == nil {
			v, err := env.Lookup(key)
			if err != nil {
				return nil, err
			}
			runtimeVal = v
		} else {
			e, err := Evaluate(*value, env)
			if err != nil {
				return nil, err
			}
			runtimeVal = e
		}

		object.Properties[key] = runtimeVal
	}

	return object, nil
}

func eval_ternary_expression(t ast.TernaryCondition, env Environment) (RuntimeValue, error) {

	// Capture the condition.
	cond, err := Evaluate(t.Condition, env)
	if err != nil {
		return nil, err
	}

	boolean, isBoolean := cond.(BooleanValue)
	if !isBoolean {
		return nil, fmt.Errorf("ternary expressions should evaluate to a boolean value, got %v", cond)
	}

	// Default to do nothing if this is just an empty function definition.
	var result RuntimeValue = MK_NULL()

	// If true.
	if boolean.Value {

		result, err = Evaluate(t.Left, env)
		if err != nil {
			return nil, err
		}
	} else {
		result, err = Evaluate(t.Right, env)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Evaluates a standard while loop.
func eval_while_expression(w ast.WhileLoop, env Environment) (RuntimeValue, error) {

	// We need to determine the type of the if's condition.
	binop, isBinop := w.Condition.(ast.BinaryExpr)
	boolean, isBool := w.Condition.(ast.BooleanLiteral)
	iden, isIden := w.Condition.(ast.Identifier)

	var isConditionTrue bool

	if isBinop {

		left, err := Evaluate(binop.Left, env)
		if err != nil {
			return nil, err
		}

		right, err := Evaluate(binop.Right, env)
		if err != nil {
			return nil, err
		}

		lhs, ok1 := left.(NumberValue)
		rhs, ok2 := right.(NumberValue)

		if ok1 && ok2 {
			b, err := eval_numeric_boolean_expression(lhs, rhs, binop.Operator)
			isConditionTrue = b.Value
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("conditions must be of same type, got %v %v", lhs, rhs)
		}
	} else if isBool {
		isConditionTrue = boolean.Value
	} else if isIden {

		b, err := env.Lookup(iden.Symbol)
		if err != nil {
			return nil, err
		}

		boolean, ok := b.(BooleanValue)
		if !ok {
			return nil, fmt.Errorf("if statement expressions must evaluate to a bool value, got %v", b)
		}

		isConditionTrue = boolean.Value
	}

	// Do we evaluate the conditional body or not?
	if isConditionTrue {

		for _, stmt := range w.Body {

			_, err := Evaluate(stmt, env)
			if err != nil {
				return nil, err
			}
		}

		eval_while_expression(w, env)
	}

	return MK_NULL(), nil
}

func eval_shorthand_operator_expression(sho ast.ShorthandOperator, env Environment) (RuntimeValue, error) {

	left, err := env.Lookup(sho.Left)
	if err != nil {
		return nil, err
	}

	lhs, lok := left.(NumberValue)

	if lok {

		currentValue := lhs.Value

		if sho.Operator == "++" || sho.Operator == "--" {
			// Simple Shorthand (x++;)

			if sho.Operator == "++" {
				currentValue++
			} else {
				currentValue--
			}

		} else {
			// Complex Shorthand (x += 1;)

			right, err := Evaluate(sho.Right, env)
			if err != nil {
				return nil, err
			}

			rhs, rok := right.(NumberValue)
			if rok {

				if sho.Operator == "*=" {
					currentValue *= rhs.Value
				} else if sho.Operator == "/=" {
					currentValue /= rhs.Value
				} else if sho.Operator == "%=" {
					currentValue %= rhs.Value
				} else if sho.Operator == "+=" {
					currentValue += rhs.Value
				} else if sho.Operator == "-=" {
					currentValue -= rhs.Value
				}

			} else {
				return nil, fmt.Errorf("invalid type used for operator %v", sho.Operator)
			}
		}

		newValue, err := env.Update(sho.Left, MK_NUMBER(currentValue))
		if err != nil {
			return nil, err
		}

		return newValue, nil

	} else {
		return nil, fmt.Errorf("invalid type used for operator %v", sho.Operator)
	}

	return nil, fmt.Errorf("invalid operator %v", sho.Operator)
}

// Evaluates an if condition, i.e. if (10 > 5) { ... }
func eval_if_condition_expression(iif ast.IfCondition, env Environment) (RuntimeValue, error) {

	// We need to determine the type of the if's condition.
	binop, isBinop := iif.Condition.(ast.BinaryExpr)
	boolean, isBool := iif.Condition.(ast.BooleanLiteral)
	iden, isIden := iif.Condition.(ast.Identifier)

	var isConditionTrue bool

	if isBinop {

		left, err := Evaluate(binop.Left, env)
		if err != nil {
			return nil, err
		}

		right, err := Evaluate(binop.Right, env)
		if err != nil {
			return nil, err
		}

		lhs, ok1 := left.(NumberValue)
		rhs, ok2 := right.(NumberValue)

		if ok1 && ok2 {
			b, err := eval_numeric_boolean_expression(lhs, rhs, binop.Operator)
			isConditionTrue = b.Value
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("conditions must be of same type, got %v %v", lhs, rhs)
		}
	} else if isBool {
		isConditionTrue = boolean.Value
	} else if isIden {

		b, err := env.Lookup(iden.Symbol)
		if err != nil {
			return nil, err
		}

		boolean, ok := b.(BooleanValue)
		if !ok {
			return nil, fmt.Errorf("if statement expressions must evaluate to a bool value, got %v", b)
		}

		isConditionTrue = boolean.Value
	}

	// Default to do nothing if this is just an empty function definition.
	var result RuntimeValue = MK_NULL()

	// Do we evaluate the conditional body or not?
	if isConditionTrue {

		for _, stmt := range iif.Body {

			r, err := Evaluate(stmt, env)
			if err != nil {
				return nil, err
			}

			// TO-DO: Currently this will make it so functions return values regardless of if there is a 'return' statement at the end.
			// We maye want to change this in the future.
			result = r
		}

	} else {
		if iif.ElseCatch && iif.ElseBody != nil {

			for _, stmt := range iif.ElseBody {

				r, err := Evaluate(stmt, env)
				if err != nil {
					return nil, err
				}

				// TO-DO: Currently this will make it so functions return values regardless of if there is a 'return' statement at the end.
				// We maye want to change this in the future.
				result = r
			}
		}
	}

	// No.
	return result, nil
}

// Evaluates a string expression.
func eval_string_expression(str ast.StringLiteral, env Environment) (RuntimeValue, error) {

	return MK_STRING(str.Value), nil
}

// Evaluates complex object assignments such as 'let foo = {x: 10};'
func eval_call_expr(expr ast.CallExpr, env Environment) (RuntimeValue, error) {

	args := make([]RuntimeValue, 0)

	// For all args passed in, evaluate and store.
	for _, arg := range expr.Args {

		val, err := Evaluate(arg, env)
		if err != nil {
			return nil, err
		}

		args = append(args, val)
	}

	fn, err := Evaluate(expr.Caller, env)
	if err != nil {
		return nil, err
	}

	// User calling built-in funciton.
	nativeFunc, isFn := fn.(NativeFunction)
	if isFn {
		// Build the function call.
		var func_ FunctionCall = nativeFunc.Call
		result := func_(args, env)

		return result, nil
	}

	// User calling user-defined function.
	userFunc, isFn := fn.(UserFunction)
	if isFn {

		newScope := Environment{
			Stdout:    env.Stdout, // Atm same stdout as main scope, however we would change to bytes.Buffer to give each new scope its own output buffer.
			Parent:    &userFunc.DecEnv,
			Constants: map[string]bool{},
			Variables: map[string]RuntimeValue{},
		}

		// Does number of provided params match the expected params?
		providedParamCount := len(expr.Args)
		expectingParamCount := len(userFunc.Params)
		if expectingParamCount != providedParamCount {
			return nil, fmt.Errorf("incorrect number of params specified for fn %v, got %v want %v", userFunc.Name, providedParamCount, expectingParamCount)
		}

		// Make vars for params list.
		for i := 0; i < expectingParamCount; i++ {
			varName := userFunc.Params[i]
			newScope.Declare(varName, args[i], false)
		}

		// Default to do nothing if this is just an empty function definition.
		var result RuntimeValue = MK_NULL()

		for _, stmt := range userFunc.Body {

			r, err := Evaluate(stmt, newScope)
			if err != nil {
				return nil, err
			}

			// TO-DO: Currently this will make it so functions return values regardless of if there is a 'return' statement at the end.
			// We may want to change this in the future.
			result = r
		}

		return result, nil
	}

	return nil, fmt.Errorf("unexpected value in place of function: %v", fn)
}

// Evaluates a function call.
func eval_function_decleration(f ast.FunctionDecleration, env Environment) (RuntimeValue, error) {

	fn := UserFunction{
		Type:   "UserFn",
		Name:   f.Name,
		Params: f.Params,
		DecEnv: env,
		Body:   f.Body,
	}

	val, err := env.Declare(f.Name, fn, true)
	if err != nil {
		return RntmVal{}, err
	}

	return val, nil
}

// Evaluates either a 'let' or 'const' decleration statement.
func eval_var_decleration(dec ast.VariableDecleration, env Environment) (RuntimeValue, error) {

	value, err := Evaluate(dec.Value, env)
	if err != nil {
		return nil, err
	}

	// Incase the value here returns null.
	if value == nil {
		value = MK_NULL()
	}

	decleration, err := env.Declare(dec.Identifier, value, dec.Constant)
	if err != nil {
		return nil, err
	}

	return decleration, nil
}

// Evaluates an array decleration.
func eval_arr_decleration(arr ast.ArrayDecleration, env Environment) (RuntimeValue, error) {

	values := make([]RuntimeValue, 0)

	for _, val := range arr.Value {

		v, err := Evaluate(val, env)
		if err != nil {
			return nil, err
		}

		values = append(values, v)
	}

	decleration, err := env.DeclareArray(arr.Identifier, values, arr.Constant)
	if err != nil {
		return nil, err
	}

	return decleration, nil
}

// Evaluates a map decleration.
func eval_map_decleration(map_ ast.MapDecleration, env Environment) (RuntimeValue, error) {

	mapValues := make(map[RuntimeValue]RuntimeValue, 0)

	for k, v := range map_.Value {

		// Evaluate the provided key.
		key, err := Evaluate(k, env)
		if err != nil {
			return nil, err
		}

		// Evaluate the provided value.
		value, err := Evaluate(v, env)
		if err != nil {
			return nil, err
		}

		mapValues[key] = value
	}

	decleration, err := env.DeclareMap(map_.Identifier, mapValues, map_.Constant)
	if err != nil {
		return nil, err
	}

	return decleration, nil
}

// Evaluates a binary expression.
func eval_binary_expression(binop ast.BinaryExpr, env Environment) (RuntimeValue, error) {

	left, err := Evaluate(binop.Left, env)
	if err != nil {
		return nil, err
	}

	right, err := Evaluate(binop.Right, env)
	if err != nil {
		return nil, err
	}

	lhs, ok1 := left.(NumberValue)
	rhs, ok2 := right.(NumberValue)

	if ok1 && ok2 {

		// Is this a mathemetical expression?
		if binop.Operator == "+" || binop.Operator == "-" ||
			binop.Operator == "/" || binop.Operator == "*" ||
			binop.Operator == "%" {
			return eval_numeric_expression(NumberValue{Type: "Number", Value: lhs.Value}, NumberValue{Type: "Number", Value: rhs.Value}, binop.Operator)

			// Or is this a boolean (logical) expression?
		} else if binop.Operator == ">" || binop.Operator == "<" || binop.Operator == ">=" || binop.Operator == "<=" || binop.Operator == "==" || binop.Operator == "!=" {
			return eval_numeric_boolean_expression(NumberValue{Type: "Number", Value: lhs.Value}, NumberValue{Type: "Number", Value: rhs.Value}, binop.Operator)
		}

	}

	// Encountered a null value.
	return MK_NULL(), nil
}

// Evaluates a numeric expression.
func eval_numeric_expression(lhs NumberValue, rhs NumberValue, opp string) (NumberValue, error) {

	result := 0

	if opp == "+" {
		result = lhs.Value + rhs.Value
	} else if opp == "-" {
		result = lhs.Value - rhs.Value
	} else if opp == "*" {
		result = lhs.Value * rhs.Value
	} else if opp == "/" {

		if rhs.Value != 0 {
			result = lhs.Value / rhs.Value
		} else {
			result = 0
		}

	} else if opp == "%" {
		result = lhs.Value % rhs.Value
	} else {
		return NumberValue{}, fmt.Errorf("invalid binop provided: %v", opp)
	}

	return MK_NUMBER(result), nil
}

// Returns a boolean evaluiation of a numeric expression. E.g. 10 < 100 === true.
func eval_numeric_boolean_expression(lhs NumberValue, rhs NumberValue, opp string) (BooleanValue, error) {

	var b bool = false

	if opp == ">" {
		b = lhs.Value > rhs.Value
	} else if opp == "<" {
		b = lhs.Value < rhs.Value
	} else if opp == ">=" {
		b = lhs.Value >= rhs.Value
	} else if opp == "<=" {
		b = lhs.Value <= rhs.Value
	} else if opp == "==" {
		b = lhs.Value == rhs.Value
	} else if opp == "!=" {
		b = lhs.Value != rhs.Value
	}

	return BooleanValue{
		Type:  Boolean,
		Value: b,
	}, nil
}

// Evaluates an assignment expression, e.g. x = 10
func eval_assignment_expression(node ast.AssignmentExpr, env Environment) (RuntimeValue, error) {

	// Checking that node isnt an Identifier.
	iden, identifier := node.Assigne.(ast.Identifier)

	if !identifier {

		return nil, fmt.Errorf("invalid lhs in expression: %v", node.Assigne)
	}

	eval, err := Evaluate(node.Value, env)
	if err != nil {
		return nil, err
	}

	assign, err := env.Assign(iden.Symbol, eval)
	if err != nil {
		return nil, err
	}
	return assign, nil
}
