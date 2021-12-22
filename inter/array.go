package inter

import (
	"fmt"
	"reflect"
)

type ArrayValue = []IType

type Array struct {
	IType
	Value ArrayValue
	Proto *Object
}

func (a Array) TypeOf() int {
	return OBJECT
}

func (a Array) String() string {
	return fmt.Sprintf("Array(%v)", a.Value)
}

func (a Array) Equal(other IType) bool {
	switch other := other.(type) {
	case *Array:
		return reflect.DeepEqual(a.Value, other.Value)
	default:
		return false
	}
}

func (a Array) Gt(other IType) bool {
	return false
}

func (a Array) Gte(other IType) bool {
	return false
}

func (a Array) Lt(other IType) bool {
	return false
}

func (a Array) Lte(other IType) bool {
	return false
}

func (a Array) Access(child IType) IType {
	switch child := child.(type) {
	case *Int:
		index := child.Value
		if index < 0 {
			panic("out of range")
		}
		if index >= len(a.Value) {
			panic("out of range")
		}
		return a.Value[index]
	case *String:
		return a.Proto.Access(child)
	default:
		fmt.Println(reflect.TypeOf(child))
		panic("Array.Access: invalid child type")
	}
}

func NewArray(value ArrayValue) *Array {
	result := &Array{
		Value: value,
	}

	proto := NewObject(ObjectValue{
		"push": NewNativeFn("push", func(env *Env, args []IType) IType {
			result.Value = append(result.Value, args...)
			return result
		}),
	})

	result.Proto = proto

	return result

}
