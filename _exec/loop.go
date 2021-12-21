package exec

import (
	"fmt"
	"plang/ast"
)

func loopThrough(stmt ast.ForInStmt, env *Env) Type {
	target := expression(stmt.Target, env)
	names := []*String{}
	for _, name := range stmt.Names {
		names = append(names, expression(ast.StringExpr{
			Value: name.(string),
		}, env).(*String))
	}

	switch target := target.(type) {
	// case *Array:
	// 	return loopArray(target, names, stmt.Body, env)
	// case *String:
	// 	loopString(target, names, stmt.Body, env)
	case *Int:
		return loopInt(target, names, stmt.Body, env)
	default:
		return NullValue
	}
}

// func loopArray(target *Array, names []*String, body []interface{}, parent *Env) {
// 	if len(names) == 1 {
// 		for index := range target.Value {
// 			env := NewForEnv(parent)
// 			env.CreateVar(names[0].Value, NewInt(index))
// 			final := ExecStatements(body, env)
// 			if final.TypeOf() != NULL {
// 				break
// 			}
// 		}
// 	} else if len(names) == 2 {
// 		for index, element := range target.Value {
// 			env := NewForEnv(parent)
// 			env.CreateVar(names[0].Value, NewInt(index))
// 			env.CreateVar(names[1].Value, element)
// 			final := ExecStatements(body, env)
// 			if final.TypeOf() != NULL {
// 				break
// 			}
// 		}
// 	} else {
// 		NewError(fmt.Sprintf("Looping accepts only 2 arguments got %v", len(names)))
// 	}
// }

// func loopString(target *String, names []*String, body []interface{}, parent *Env) {
// 	value := strings.Split(target.Value, "")
// 	if len(names) == 1 {
// 		for index := range value {
// 			env := NewForEnv(parent)
// 			env.CreateVar(names[0].Value, NewInt(index))
// 			final := ExecStatements(body, env)
// 			if final.TypeOf() != NULL {
// 				break
// 			}
// 		}
// 	} else if len(names) == 2 {
// 		for index, element := range value {
// 			env := NewForEnv(parent)
// 			env.CreateVar(names[0].Value, NewInt(index))
// 			env.CreateVar(names[1].Value, NewString(element))
// 			final := ExecStatements(body, env)
// 			if final.TypeOf() != NULL {
// 				break
// 			}
// 		}
// 	} else {
// 		NewError(fmt.Sprintf("Looping accepts only 2 arguments got %v", len(names)))
// 	}
// }

func loopInt(target *Int, names []*String, body []interface{}, parent *Env) Type {
	var result Type
	if len(names) == 1 {
		for index := 0; index < target.Value; index++ {
			env := NewForEnv(parent)
			env.CreateVar(names[0].Value, NewInt(index))
			final := ExecStatements(body, env)
			if final.TypeOf() != NULL {
				result = final
				break
			}
		}
	} else {
		NewError(fmt.Sprintf("Looping accepts only 1 arguments got %v", len(names)))
	}
	return result
}
