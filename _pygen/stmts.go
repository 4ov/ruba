package pygen

import (
	"fmt"
	"plang/ast"
	"strings"
)

func declareStmt(stmt ast.DeclareStmt) string {

	return fmt.Sprintf("%v = %v;\n", stmt.Name, expression(stmt.Value))
}

func fnCallStmt(stmt ast.FnCallStmt) string {
	args := []string{}
	for _, e := range stmt.Args {
		args = append(args, expression(e))
	}
	return fmt.Sprintf("%v(%v);\n", stmt.Name, strings.Join(args, ","))
}
