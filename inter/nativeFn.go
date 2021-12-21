package inter

import (
	"fmt"

	"github.com/4ov/ruba/ast"
)

type NativeFn struct {
	IType
	Name  string
	Value func(env *Env, args []IType) IType
	Proto *Object
}

func (n NativeFn) TypeOf() int {
	return FUNCTION
}

func (n NativeFn) String() string {
	return fmt.Sprintf("NativeFn(%s)", n.Name)
}

func (n NativeFn) Access(child IType) IType {
	return n.Proto.Access(child)
}

func NewNativeFn(name string, fn func(env *Env, args []IType) IType) *NativeFn {
	return &NativeFn{
		Name:  name,
		Value: fn,
		Proto: NewObject(ObjectValue{
			// "length": NewInt(1),
		}),
	}
}

func CallNativeFn(target *NativeFn, args []ast.Expression, env *Env) IType {
	callArgs := []IType{}
	for _, expArg := range args {
		arg := expArg.(ast.ArgExpr)
		if !arg.Spreadable {
			callArgs = append(callArgs, expression(arg.Value, env))
		} else {
			callArgs = append(callArgs, expression(arg.Value, env).(*Array).Value...)
		}
	}

	fnEnv := NewFnEnv(target.Name, env)
	return target.Value(fnEnv, callArgs)
}
