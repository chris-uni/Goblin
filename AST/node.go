package ast

type NodeType string

const (
	ProgramNode          NodeType = "Program"
	NumericLiteralNode   NodeType = "NumericLiteral"
	IdentifierNode       NodeType = "Identifier"
	BinaryExpressionNode NodeType = "BinaryExpression"
)

type Expression interface {
	expr()
}

type Expr struct {
	Kind NodeType
}

func (e Expr) expr() {}

// -- Program
type Program struct {
	Kind NodeType
	Body []Expression
}

func (p Program) expr() {}

// -- BinaryExpr
type BinaryExpr struct {
	Kind     NodeType
	Left     Expression
	Right    Expression
	Operator string
}

func (b BinaryExpr) expr() {}

// -- Identifier

type Identifier struct {
	Kind   NodeType
	Symbol string
}

func (i Identifier) expr() {}

// -- NumericExpr
type NumericExpr struct {
	Kind  NodeType
	Value float64
}

func (n NumericExpr) expr() {}
