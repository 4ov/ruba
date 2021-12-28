package inter

import "fmt"

const (
	MAIN_ENV = iota
	FN_ENV
	IF_ENV
)

type Env struct {
	// Ruba  *Ruba
	Parent *Env
	Store  *Object
	Type   int
}

func (env *Env) CreateVar(name string, value IType) {
	env.Store.Set(name, value)
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

// func (env *Env) Get(key string) IType {
// 	if r := env.Store.Get(key); r != nil {
// 		return r
// 	} else {
// 		if len(env.Stack) > 0 {
// 			fmt.Println()
// 			return reverseEnvStack(env.Stack)[0].Get(key)
// 		} else {
// 			panic(fmt.Sprintf("%s not found", key))
// 		}
// 	}
// }

func (env *Env) Get(key string) IType {
	var result IType
	if r := env.Store.Get(key); r != nil {
		return r
	} else {
		if env.Parent != nil {
			result = env.Parent.Get(key)
		}
	}
	if result == nil {
		panic(fmt.Sprintf("%s not found", key))
	}
	return result
}

func (env *Env) InFnEnv() bool {
	if env.Type == FN_ENV {
		return true
	}
	if env.Parent != nil {
		return env.Parent.InFnEnv()
	}
	return false
}

func NewEnv() *Env {
	return &Env{
		Type:  MAIN_ENV,
		Store: NewObject(ObjectValue{}),
	}
}

func NewFnEnv(name string, parent *Env) *Env {
	return &Env{
		Type:   FN_ENV,
		Parent: parent,
		// Stack: append(parent.Stack, parent),
		Store: NewObject(ObjectValue{}),
	}
}

func NewIfEnv(parent *Env) *Env {
	return &Env{
		Type: IF_ENV,
		// Stack: append(parent.Stack, parent),
		Parent: parent,
		Store:  NewObject(ObjectValue{}),
	}
}
