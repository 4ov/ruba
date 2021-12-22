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

func bakeArgs(args []ast.Expression, env *Env) []IType {
	callArgs := []IType{}
	for _, expArg := range args {
		arg := expArg.(ast.ArgExpr)
		if !arg.Spreadable {
			callArgs = append(callArgs, expression(arg.Value, env))
		} else {
			callArgs = append(callArgs, expression(arg.Value, env).(*Array).Value...)
		}
	}
	return callArgs
}

func reverseEnvStack(stack []*Env) []*Env {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

func nilAsNull(value interface{}) IType {
	if value == nil {
		return NewNull()
	}
	return value.(IType)
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
