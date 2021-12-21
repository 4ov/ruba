package checker

import "plang/ast"

type Ident struct {
	Value string
	Type  string
}

type Memory struct {
	Idents map[string]int
}

func Check(program ast.Program) {
	memory := Memory{
		Idents: make(map[string]int),
	}

	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case ast.DeclareStmt:
			declareStmt(stmt, &memory)
			// case ast.FnCallStmt:
			// 	fnCallStmt(stmt, &memory)
		}
	}

}
