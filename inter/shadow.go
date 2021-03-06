package inter

import (
	"github.com/4ov/ruba/ast"
)

func Exec(tree []interface{}, env *Env) IType {
	var result IType = nil
exec:
	for _, stmt := range tree {
		switch stmt := stmt.(type) {
		case ast.DeclareStmt:
			result = declareStmt(stmt, env)
		case ast.AssignStmt:
			result = assignStmt(stmt, env)
		case ast.FnCallStmt:
			fnCallStmt(stmt, env)
		case ast.FnDeclareStmt:
			fnDeclareStmt(stmt, env)
		case ast.ReturnStmt:
			result = returnStmt(stmt, env)
			break exec
		case ast.IfStmt:
			result = ifStmt(stmt, env)
			if result != nil {
				break exec
			}
		case ast.ExprStmt:
			result = exprStmt(stmt, env)

		}
	}
	return result
}
