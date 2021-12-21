package inter

import (
	"github.com/4ov/ruba/ast"
)

func nestAssign(parent IType, chain []IType, value IType) IType {
	switch parent := parent.(type) {
	case *Int:
		panic("cannot assign to int")
	case *String:
		panic("cannot assign to string")
	case *Object:
		parent.Assign(chain, value)
	}

	return nil
}

func bakeArray(array []ast.Expression, env *Env) []IType {
	var result []IType
	for _, expr := range array {
		result = append(result, expression(expr, env))
	}
	return result
}

// func keyable(value IType) bool {
// 	switch value := value.(type) {
// 	case Int:
// 		return true
// 		// case :
// 		// 	return true
// 	}
// 	return false
// }
