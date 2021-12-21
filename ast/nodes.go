package ast

const (
	NUMBERT = iota
	IDENTT
	STRINGT
	BOOLT
)

type Program struct {
	Statements []interface{}
}

type DeclareStmt struct {
	Name   string
	Value  Expression
	Public bool
	Line   int
	Col    int
}

type ImportStmt struct {
	Name  string
	Alias string
}

type AssignStmt struct {
	Name  string
	Chain []Expression
	Value Expression

	Line int
	Col  int
}

type NestAssignStmt struct {
	Root     IdentExpr
	Children []Expression
	Value    Expression
	// Line     int
	// Col      int
}

type IndexAssignStmt struct {
	Parent   IdentExpr
	Children []Expression
	Value    Expression
	Line     int
	Col      int
}

type FnCallStmt struct {
	Target Expression
	Args   []Expression
}

type ReturnStmt struct {
	Value Expression
}

type BreakStmt struct {
	Value Expression
}

type IfStmt struct {
	Cond         Expression
	ElseIfConds  []Expression
	ElseIfBodies [][]interface{}
	ElseBody     []interface{}
	Body         []interface{}
}

type FnDeclareStmt struct {
	Name    Expression
	Args    []Expression
	RestArg Expression
	Body    []interface{}
}

type ForInStmt struct {
	Names  []Expression
	Target Expression
	Body   []interface{}
}

type ForCondStmt struct {
	Cond     Expression
	Body     []interface{}
	ElseBody []interface{}
}

type ExprStmt = Expression
