package inter

const (
	MAIN_ENV = iota
	FN_ENV
)

type Env struct {
	// Ruba  *Ruba
	Stack []*Env
	Store *Object
	Type  int
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

func (env *Env) Get(key string) IType {
	if r := env.Store.Get(key); r != nil {
		return r
	} else {
		if len(env.Stack) > 0 {
			return env.Stack[len(env.Stack)-1].Get(key)
		} else {
			return nil
		}

	}
}

func (env *Env) InFnEnv() bool {
	for _, e := range reverseEnvStack(append(env.Stack, env)) {
		if e.Type == FN_ENV {
			return true
		}
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
		Type:  FN_ENV,
		Stack: append(parent.Stack, parent),
		Store: NewObject(ObjectValue{}),
	}
}
