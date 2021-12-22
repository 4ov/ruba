package inter

import (
	"fmt"

	"github.com/4ov/ruba/ast"
)

type RestArgValue struct {
	Value string
}

type Fn struct {
	Name    string
	Args    []string
	RestArg interface{}
	Proto   *Object
	Body    []interface{}
}

func (f Fn) String() string {
	return fmt.Sprintf("Fn( %v )", f.Name)
}

func (f Fn) TypeOf() int {
	return FUNCTION
}

func (f Fn) Access(child IType) IType {
	return f.Proto.Access(child)
}

func NewFn(name string, args []string, restArg interface{}, body []interface{}) *Fn {
	return &Fn{
		Name:    name,
		Args:    args,
		RestArg: restArg,
		Body:    body,
	}
}

func CallFn(f *Fn, args []ast.Expression, env *Env) IType {
	callArgs := bakeArgs(args, env)
	fnEnv := NewFnEnv(f.Name, env)
	//note: i love GH copilot
	if len(callArgs) > 0 {
		for argIndex, argName := range f.Args {
			fnEnv.Set(argName, callArgs[argIndex])
			if argIndex == len(f.Args)-1 { //last arg
				callArgs = callArgs[1:]
			}

		}
	}
	if f.RestArg != nil {
		fnEnv.Set(f.RestArg.(RestArgValue).Value, NewArray(callArgs))
	}
	return Exec(f.Body, fnEnv)

}
