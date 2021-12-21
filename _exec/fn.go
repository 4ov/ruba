package exec

import "fmt"

type Fn struct {
	Type
	Body      []interface{}
	Args      []Type
	RestArg   Type
	Name      string
	Prototype *Object
}

func (fn Fn) String() string {
	return fmt.Sprintf("[Function %v]", fn.Name)
}

func (fn Fn) TypeOf() int {
	return FN
}

func (fn Fn) True() bool {
	return true
}

func NewFn(name string, body []interface{}, args []Type, restArg Type) *Fn {
	result := &Fn{
		Body:      body,
		Args:      args,
		RestArg:   restArg,
		Name:      name,
		Prototype: NewObject(ObjectValue{}),
	}

	return result
}

func CallFn(fn Fn, parent *Env, callArgs []Type) Type {
	env := NewFnEnv(parent, fn.Name, fn.Args, fn.RestArg, callArgs)
	// env.SetVar("this", env.Store)

	return ExecStatements(fn.Body, env)
}
