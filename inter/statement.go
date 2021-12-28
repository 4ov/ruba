package inter

import (
	"fmt"

	"github.com/4ov/ruba/ast"
)

func exprStmt(stmt ast.ExprStmt, env *Env) IType {
	return expression(stmt, env)
}

func declareStmt(stmt ast.DeclareStmt, env *Env) IType {
	value := expression(stmt.Value, env)
	// env.Ruba.debug(fmt.Sprintf("%s :> %v", stmt.Name, value))
	env.Set(stmt.Name, value)
	return NewInt(0)
}

func assignStmt(stmt ast.AssignStmt, env *Env) IType {
	single := len(stmt.Chain) == 0
	if env.Has(stmt.Name) {
		value := expression(stmt.Value, env)
		if single {
			env.Set(stmt.Name, value)
			// env.Ruba.debug(fmt.Sprintf("%s => %v", stmt.Name, value))
		} else {
			root := env.Get(stmt.Name)
			nestAssign(root, bakeArray(stmt.Chain, env), value)
		}
	} else {
		panic(fmt.Sprintf("%s is not defined", stmt.Name))
	}

	return nil
}

func fnCallStmt(stmt ast.FnCallStmt, env *Env) IType {
	var target IType
	//I GUESS THIS IS A BUG

	switch raw := stmt.Target.(type) {
	case string:
		target = env.Get(raw)
	default:
		target = expression(raw, env)
	}

	switch target := target.(type) {
	case *NativeFn:
		return CallNativeFn(target, stmt.Args, env)
	case *Fn:
		return CallFn(target, stmt.Args, env)
	default:
		panic(fmt.Sprintf("%s is not a function", target))
	}
}

func fnDeclareStmt(stmt ast.FnDeclareStmt, env *Env) IType {
	args := []string{}
	var restArg interface{}
	if stmt.RestArg != nil {
		restArg = RestArgValue{
			Value: stmt.RestArg.(string),
		}
	}
	for _, arg := range stmt.Args {
		args = append(args, arg.(string))
	}

	env.Set(stmt.Name.(string), NewFn(stmt.Name.(string), args, restArg, stmt.Body))

	return nil
}

func returnStmt(stmt ast.ReturnStmt, env *Env) IType {
	if env.InFnEnv() {
		if stmt.Value == nil {
			return NewNull()
		}
		return expression(stmt.Value, env)
	} else {
		panic("Cannot use return statement outside of function")
	}

}

func ifStmt(stmt ast.IfStmt, env *Env) IType {
	ifEnv := NewIfEnv(env)
	if expression(stmt.Cond, env).Bool() {
		return Exec(stmt.Body, ifEnv)
	}

	return nil
}
