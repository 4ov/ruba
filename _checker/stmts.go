package checker

import (
	"plang/ast"
)

func declareStmt(stmt ast.DeclareStmt, memory *Memory) {
	expr := expression(stmt.Value, memory)
	memory.Idents[stmt.Name] = expr
	// fmt.Println(memory)

}

// func fnCallStmt(stmt ast.FnCallStmt, memory *Memory) string {
// 	args := []string{}
// 	for _, e := range stmt.Args {
// 		args = append(args, expression(e))
// 	}
// 	return fmt.Sprintf("%v(%v);\n", stmt.Name, strings.Join(args, ","))
// }
