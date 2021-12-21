package exec

import (
	"fmt"
	"os"
	"plang/ast"
	"reflect"
)

func ExecStatements(stmts []interface{}, env *Env) Type {

	var result Type = NewNull()
execution_loop:
	for _, stmt := range stmts {
		switch stmt := stmt.(type) {
		case ast.DeclareStmt:
			declareStmt(stmt, env)
		case ast.AssignStmt:
			assignStmt(stmt, env)
		case ast.FnCallStmt:
			fnCallStmt(stmt, env)
		case ast.FnDeclareStmt:
			fnDeclareStmt(stmt, env)
		case ast.ReturnStmt:
			result = returnStmt(stmt, env)
			break execution_loop
		case ast.BreakStmt:
			result = breakStmt(stmt, env)
			if result.TypeOf() == BOOL {
				break execution_loop
			}
		case ast.IfStmt:
			inFn, r := ifStmt(stmt, env)
			if inFn && r.TypeOf() != NULL {
				result = r
				break execution_loop
			} else if r.TypeOf() != NULL {
				result = r
				break execution_loop
			}
		case ast.ForInStmt:
			r := forInStmt(stmt, env)
			if r != nil {
				result = r
				break execution_loop
			}
		case ast.ExprStmt:
			exprStmt(stmt, env)
		default:
			fmt.Println(reflect.TypeOf(stmt))
			os.Exit(1)
		}
	}

	return result
}

func declareStmt(stmt ast.DeclareStmt, env *Env) Type {
	value := expression(stmt.Value, env)
	name := stmt.Name
	env.CreateVar(name, value)
	return nil
}

func assignStmt(stmt ast.AssignStmt, env *Env) Type {
	isChain := len(stmt.Chain) != 0

	if isChain {
		root := env.GetVar(stmt.Name)
		switch root := root.(type) {
		case *Array:
			root.Set(exprToType(stmt.Chain, env), expression(stmt.Value, env))
		default:
			NewError(fmt.Sprintf("Cannot set property of type %v", type_map[root.TypeOf()]))
		}
	} else {

		fmt.Printf("stmt.Chain: %v\n", stmt.Chain)
		env.SetVar(stmt.Name, expression(stmt.Value, env))
	}
	return nil
}

func fnCallStmt(stmt ast.FnCallStmt, env *Env) Type {
	rawTarget := stmt.Target
	var args []Type
	var target Type
	switch raw := rawTarget.(type) {
	case ast.DotExpr:
		target = expression(raw, env)
	case string:
		target = expression(ast.IdentExpr{
			Value: stmt.Target.(string),
		}, env)
	}

	switch target := target.(type) {
	case *NativeFn:
		for _, arg := range stmt.Args {
			if arg.(ast.ArgExpr).Spreadable {
				arr := expression(arg, env).(*Array).Value
				args = append(args, arr...)
			} else {

				args = append(args, expression(arg, env))
			}

		}
		CallNativeFn(*target, env, args)
	case *Fn:
		for _, arg := range stmt.Args {
			if arg.(ast.ArgExpr).Spreadable {
				arr := expression(arg, env).(*Array).Value
				args = append(args, arr...)
			} else {

				args = append(args, expression(arg, env))
			}

		}
		CallFn(*target, env, args)
	default:
		// fmt.Println(stmt.Target)
		NewError(fmt.Sprintf("Cannot call value of type %v", type_map[target.TypeOf()]))

	}
	return nil
}

func fnDeclareStmt(stmt ast.FnDeclareStmt, env *Env) Type {
	args := []Type{}
	var restArg Type
	for _, arg := range stmt.Args {
		args = append(args, expression(ast.StringExpr{
			Value: arg.(string),
		}, env))
	}
	if stmt.RestArg != nil {
		restArg = expression(ast.StringExpr{Value: stmt.RestArg.(string)}, env)
	}

	env.CreateVar(stmt.Name.(string), NewFn(stmt.Name.(string), stmt.Body, args, restArg))
	return nil
}

func returnStmt(stmt ast.ReturnStmt, env *Env) Type {
	if env.HasFn() {
		if stmt.Value != nil {
			return expression(stmt.Value, env)
		} else {
			return NewNull()
		}
	} else {
		NewError("Cannot use return stmt outside functions")
	}
	return NewNull()
}

func breakStmt(stmt ast.BreakStmt, env *Env) Type {
	if env.HasLoop() {
		return NewBool(true)
	} else {
		NewError("Cannot use break statement outside loop")
		return nil
	}

}

func exprStmt(stmt ast.ExprStmt, env *Env) Type {
	return expression(stmt, env)
}

func forInStmt(stmt ast.ForInStmt, env *Env) Type {
	r := loopThrough(stmt, env)
	return r
}

func ifStmt(stmt ast.IfStmt, env *Env) (bool, Type) {
	ifEnv := NewIfEnv(env)
	isFn := ifEnv.HasFn()
	if expression(stmt.Cond, env).True() {
		return isFn, ExecStatements(stmt.Body, ifEnv)
	} else {
		for condIndex, cond := range stmt.ElseIfConds {
			if expression(cond, env).True() {
				return isFn, ExecStatements(stmt.ElseIfBodies[condIndex], ifEnv)
			}
		}
		if stmt.ElseBody != nil {
			return isFn, ExecStatements(stmt.ElseBody, ifEnv)
		}
	}

	return isFn, NewNull()

}
