package exec

import "fmt"

const (
	MAIN_ENV = iota
	FN_ENV
	IF_ENV
	LOOP_ENV
)

func reserveStack(stack []*Env) []*Env {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

type Env struct {
	Name  string
	Type  int
	Stack []*Env
	Store *Object
}

func (env *Env) CreateVar(key string, val Type) {
	if !env.Store.Has(key) {
		env.Store.Set(key, val)
	} else {
		panic("already defined")
	}
}

func (env *Env) GetVar(key string) Type {
	var result *Type
	var found bool = false
	if value := env.Store.Get(key); value != nil {
		result = &value
		found = true
	} else {
	lookup:
		for _, level := range reserveStack(env.Stack) {
			if value := level.Store.Get(key); value != nil {
				result = &value
				found = true
				break lookup
			}
		}
	}

	if !found {
		NewError(fmt.Sprintf("%v is not defined", key))
	}
	return *result
}

func (env *Env) SetVar(key string, newValue Type) {
	var found bool
	if value := env.Store.Get(key); value != nil {
		env.Store.Set(key, newValue)
		found = true
	} else {
	lookup:
		for _, level := range env.Stack {
			if value := level.Store.Get(key); value != nil {
				level.Store.Set(key, value)
				found = true
				break lookup
			}
		}
	}
	if !found {
		panic("not defined")
	}

}

func (env *Env) HasFn() bool {
	if env.Type == FN_ENV {
		return true
	} else {
		for _, env := range env.Stack {
			if env.Type == FN_ENV {
				return true
			}
		}
		return false
	}
}

func (env *Env) HasLoop() bool {
	if env.Type == LOOP_ENV {
		return true
	} else {
		for _, env := range env.Stack {
			if env.Type == LOOP_ENV {
				return true
			}
		}
		return false
	}
}

func NewNativeFnEnv(parent *Env, name string, args []Type) *Env {
	checkStack(parent)
	result := &Env{
		Name:  name,
		Type:  FN_ENV,
		Stack: append(parent.Stack, parent),
		Store: NewObject(ObjectValue{
			"args": NewArray(args),
		}),
	}

	return result
}

func NewFnEnv(parent *Env, name string, argNames []Type, restArg Type, callArgs []Type) *Env {
	checkStack(parent)

	store := NewObject(ObjectValue{})
	leftArgs := callArgs

	// for argIndex, arg := range argNames {
	// 	argName := arg.(*String)
	// 	store.Set(argName.Value, callArgs[argIndex])
	// 	leftArgs = leftArgs[1:]
	// }

	if len(callArgs) > 0 {
		for argIndex, arg := range argNames {
			argName := arg.(*String)
			store.Set(argName.Value, callArgs[argIndex])
			leftArgs = leftArgs[1:]
		}
	} else {
		for _, arg := range argNames {
			argName := arg.(*String)
			store.Set(argName.Value, NullValue)
			// leftArgs = leftArgs[1:]
		}
	}

	if restArg != nil {
		restArgName := restArg.(*String).Value
		store.Set(restArgName, NewArray(leftArgs))
	}

	result := &Env{
		Name:  name,
		Type:  FN_ENV,
		Stack: append(parent.Stack, parent),
		Store: store,
	}

	return result
}

func NewIfEnv(parent *Env) *Env {
	checkStack(parent)
	store := NewObject(ObjectValue{})

	result := &Env{
		Name:  "<anonymouse>",
		Type:  IF_ENV,
		Stack: append(parent.Stack, parent),
		Store: store,
	}

	return result
}

func NewForEnv(parent *Env) *Env {
	checkStack(parent)
	store := NewObject(ObjectValue{})
	result := &Env{
		Name:  "<unnamed>",
		Type:  LOOP_ENV,
		Stack: append(parent.Stack, parent),
		Store: store,
	}

	return result
}

func checkStack(env *Env) {
	if len(env.Stack) == (MAX_STACK_SIZE - 1) {
		NewError("max stack size reached")
	}
}
