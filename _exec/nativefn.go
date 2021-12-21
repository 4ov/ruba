package exec

import "fmt"

type NativeValue = func(env *Env) Type

type NativeFn struct {
	Type
	Value     NativeValue
	Name      string
	Prototype *Object
}

func (fn NativeFn) String() string {
	return fmt.Sprintf("[Function %v]", fn.Name)
}

func (fn NativeFn) TypeOf() int {
	return FN
}

func (fn NativeFn) True() bool {
	return true
}

func NewNativeFn(name string, fn NativeValue) *NativeFn {

	result := &NativeFn{
		Name:  name,
		Value: fn,
	}
	return result
}

func CallNativeFn(fn NativeFn, parent *Env, args []Type) Type {
	env := NewNativeFnEnv(parent, fn.Name, args)
	return fn.Value(env)

}
