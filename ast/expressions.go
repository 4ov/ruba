package ast

type Expression interface{}

type NumberExpr struct {
	Expression
	Value string
}

type IntExpr struct {
	Expression
	Value string
}

type FloatExpr struct {
	Expression
	Value string
}

type ArgExpr struct {
	Expression
	Value      Expression
	Spreadable bool
}

type NegativeExpr struct {
	Expression
	Value Expression
}

type StringExpr struct {
	Expression
	Value string
}

type BoolExpr struct {
	Expression
	Value bool
}

type NullExpr struct {
	Expression
}

type ArrayExpr struct {
	Expression
	Elements []Expression
}

type ObjectExpr struct {
	Expression
	Elements map[Expression]Expression
}

type ObjectField struct {
	Expression
	Key   Expression
	Value Expression
}

type IdentExpr struct {
	Expression
	Value string
}

type IndexExpr struct {
	Expression
	Parent Expression
	Child  Expression
}

type InternalConstExpr struct {
	Expression
	Value string
}

type MathExpr struct {
	Expression
	Left     Expression
	Right    Expression
	Operator string
}

type EqualExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type NotEqualExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type GtExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type GteExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type NotExpr struct {
	Expression
	Value Expression
}

type SubExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type DivExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type MulExpr struct {
	Expression
	Left  Expression
	Right Expression
}

type DotExpr struct {
	Expression
	Parent Expression
	Child  string
}

type FnCallExpr struct {
	Expression
	Target Expression
	Args   []Expression
}

type IncrExpr struct {
	Expression
	Value Expression
}

type IfExpr struct {
	Expression
	Cond  Expression
	True  Expression
	False Expression
}
