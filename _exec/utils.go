package exec

import "plang/ast"

func exprToType(exprs []ast.Expression, env *Env) []Type {
	result := []Type{}
	for _, expr := range exprs {
		result = append(result, expression(expr, env))
	}

	return result
}
