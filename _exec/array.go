package exec

import (
	"fmt"
	"strings"
)

type ArrayValue = []Type

type Array struct {
	Type
	Value     ArrayValue
	Prototype *Object
}

func (v Array) String() string {
	return fmt.Sprint(v.Value)
}

func (v Array) TypeOf() int {
	return ARRAY
}

func (v Array) True() bool {
	return true
}

func (v Array) Access(child string) Type {
	value := v.Prototype.Get(child)
	return value
}

func (v *Array) Get(index Type) {

}

func (v *Array) Set(index []Type, value Type) {
	if len(index) == 1 {
		switch index := index[0].(type) {
		case *Int:
			if index.Value >= 0 {
				if len(v.Value) > index.Value {
					v.Value[index.Value] = value
				} else {
					NewError("Assigment out of range")
				}
			} else {
				NewError("Assigment out of range")
			}

		default:
			NewError(fmt.Sprintf("Value of type %v is only settable by int %v values", type_map[ARRAY], type_map[NUMBER]))
		}
	} else {
		fmt.Println("wooo")
	}
}

func NewArray(value ArrayValue) *Array {
	proto := NewObject(ObjectValue{})
	result := &Array{
		Value:     value,
		Prototype: proto,
	}

	proto.Set("push", NewNativeFn("push", func(env *Env) Type {
		args := env.GetVar("args").(*Array).Value
		result.Value = append(result.Value, args...)
		return NewBool(true)
	}))

	proto.Set("join", NewNativeFn("push", func(env *Env) Type {
		joiner := env.GetVar("args").(*Array).Value[0].String()
		elms := []string{}
		for _, elm := range result.Value {
			elms = append(elms, elm.String())
		}
		return NewString(strings.Join(elms, joiner))
	}))

	proto.Set("length", NewGetter(func() Type {
		return NewInt(len(result.Value))
	}))

	return result

}
