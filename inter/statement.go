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
	env.Ruba.debug(fmt.Sprintf("%s :> %v", stmt.Name, value))
	env.Set(stmt.Name, value)
	return NewInt(0)
}

func assignStmt(stmt ast.AssignStmt, env *Env) IType {
	single := len(stmt.Chain) == 0
	if env.Has(stmt.Name) {
		value := expression(stmt.Value, env)
		if single {
			env.Set(stmt.Name, value)
			env.Ruba.debug(fmt.Sprintf("%s => %v", stmt.Name, value))
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
	}
	return nil
}
