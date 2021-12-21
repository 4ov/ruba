package cgen

import (
	"fmt"
	"plang/ast"
	"strings"
)

func ExecStmts(stmts []interface{}) string {
	result := ""
	for _, stmt := range stmts {
		switch stmt := stmt.(type) {
		case ast.FnDeclareStmt:
			result += fnDeclareStmt(stmt)
		case ast.IfStmt:
			result += ifStmt(stmt)
		}
	}

	return result
}

func fnDeclareStmt(stmt ast.FnDeclareStmt) string {
	args := []string{}
	for _, arg := range stmt.Args {
		args = append(args, expression(arg))
	}

	return fmt.Sprintf("function %v(%v){%v}\n", stmt.Name, strings.Join(args, ","), ExecStmts(stmt.Body))
}

func ifStmt(stmt ast.IfStmt) string {
	mainCond := expression(stmt.Cond)
	var elifConds string
	for condI, cond := range stmt.ElseIfConds {
		elifConds += fmt.Sprintf("else if (%v){%v}", expression(cond), ExecStmts(stmt.ElseIfBodies[condI]))
	}
	return fmt.Sprintf("if(%v){%v}%v", mainCond, "BODY", elifConds)
}
