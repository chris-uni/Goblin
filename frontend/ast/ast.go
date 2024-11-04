package ast

type NodeType string

const (
	// Statements.
	ProgramNode             NodeType = "ProgramNode"
	VariableDeclerationNode NodeType = "VariableDeclerationNode"
	FunctionDeclerationNode NodeType = "FunctionDeclerationNode"
	ArrayDeclerationNode    NodeType = "ArrayDeclerationNode"
	MapDeclerationNode      NodeType = "MapDeclerationNode"

	// Expressions.
	BinaryExprNode       NodeType = "BinaryExprNode"
	AssingmentExprNode   NodeType = "AssignmentExprNode"
	CallExprNode         NodeType = "CallExprNode"
	UnaryExprNode        NodeType = "UnaryExprNode"
	FuncDeclerationNode  NodeType = "FuncDeclerationNode"
	MemberExpressionNode NodeType = "MemberExpressionNode"
	CallExpression       NodeType = "CallExpression"

	// Literals.
	NumericLiteralNode  NodeType = "NumericLiteralNode"
	StringLiteralNode   NodeType = "StringLiteralNode"
	BooleanLiteralNode  NodeType = "BooleanLiteralNode"
	IdentifierNode      NodeType = "IdentifierNode"
	ArrayIdentifierNode NodeType = "ArrayIdentifierNode"
	PropertyNode        NodeType = "PropertyNode"
	ObjectLiteralNode   NodeType = "ObjectLiteralNode"

	// Conditionals
	IfNode      NodeType = "IfNode"
	TernaryNode NodeType = "TernaryNode"
	WhileNode   NodeType = "WhileNode"

	// Misc.
	UnknownNode NodeType = "UnknownNode"
)

type Expression interface {
	expr()
}

type Expr struct {
	Kind NodeType
}

func (e Expr) expr() {}

type Program struct {
	Kind NodeType
	Body []Expression
}

func (p Program) expr() {}

type VariableDecleration struct {
	Kind       NodeType
	Value      Expression
	Constant   bool
	Identifier string
}

func (v VariableDecleration) expr() {}

type ArrayDecleration struct {
	Kind       NodeType
	Value      []Expression
	Constant   bool
	Identifier string
}

func (a ArrayDecleration) expr() {}

type MapDecleration struct {
	Kind       NodeType
	Identifier string
	Value      map[Expression]Expression
	Constant   bool
}

func (m MapDecleration) expr() {}

type FunctionDecleration struct {
	Kind   NodeType
	Params []string
	Name   string
	Body   []Expression
}

func (f FunctionDecleration) expr() {}

type BinaryExpr struct {
	Kind     NodeType
	Left     Expression
	Right    Expression
	Operator string
}

func (b BinaryExpr) expr() {}

type CallExpr struct {
	Kind   NodeType
	Args   []Expression
	Caller Expression
}

func (c CallExpr) expr() {}

type MemberExpr struct {
	Kind     NodeType
	Object   Expression
	Property Expression
	Computed bool
}

func (m MemberExpr) expr() {}

type AssignmentExpr struct {
	Kind    NodeType
	Value   Expression
	Assigne Expression
}

func (a AssignmentExpr) expr() {}

type Identifier struct {
	Kind   NodeType
	Symbol string
}

func (i Identifier) expr() {}

type ArrayOrMapIdentifier struct {
	Kind   NodeType
	Symbol string
	Index  Expression
}

func (aom ArrayOrMapIdentifier) expr() {}

type NumericLiteral struct {
	Kind  NodeType
	Value int
}

func (n NumericLiteral) expr() {}

type BooleanLiteral struct {
	Kind  NodeType
	Value bool
}

func (b BooleanLiteral) expr() {}

type StringLiteral struct {
	Kind  NodeType
	Value string
}

func (b StringLiteral) expr() {}

type IfCondition struct {
	Kind      NodeType
	Condition Expression
	Body      []Expression
	ElseCatch bool
	ElseBody  []Expression
}

func (n IfCondition) expr() {}

type TernaryCondition struct {
	Kind      NodeType
	Condition Expression
	Left      Expression
	Right     Expression
}

func (t TernaryCondition) expr() {}

type WhileLoop struct {
	Kind      NodeType
	Condition Expression
	Body      []Expression
}

func (w WhileLoop) expr() {}

type Property struct {
	Kind  string
	Key   string
	Value *Expression
}

type ObjectLiteral struct {
	Kind       NodeType
	Properties []Property
}

func (o ObjectLiteral) expr() {}

type Unknown struct {
	Kind NodeType
}

func (u Unknown) expr() {}
