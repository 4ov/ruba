package inter

import "fmt"

type Env struct {
	Ruba  *Ruba
	Stack []*Env
	Store *Object
}

func (env *Env) Set(key string, value IType) {
	env.Store.Set(key, value)
}

func (env *Env) Has(key string) bool {
	if r := env.Store.Get(key); r != nil {
		return true
	} else {
		return false
	}
}

func (env *Env) Get(key string) IType {
	if r := env.Store.Get(key); r != nil {
		return r
	} else {
		panic(fmt.Sprintf("%s is not defined", key))
	}
}

func NewEnv() *Env {
	return &Env{
		Store: NewObject(ObjectValue{}),
	}
}

func NewFnEnv(name string, parent *Env) *Env {
	return &Env{
		Stack: append(parent.Stack, parent),
	}
}
