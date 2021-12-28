package inter

import (
	"fmt"
	"reflect"

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

func (n NativeFn) Equal(other IType) bool {
	return reflect.DeepEqual(n, other)
}
func (n NativeFn) Gt(other IType) bool {
	return false
}
func (n NativeFn) Gte(other IType) bool {
	return false
}
func (n NativeFn) Lt(other IType) bool {
	return false
}
func (n NativeFn) Lte(other IType) bool {
	return false
}

func (n NativeFn) Bool() bool {
	return true
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
