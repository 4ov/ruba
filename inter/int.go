package inter

import (
	"fmt"
	"reflect"
)

type Int struct {
	IType
	Value int
	Proto *Object
}

func (i Int) TypeOf() int {
	return INT
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i Int) Access(child IType) IType {
	fmt.Println(reflect.TypeOf(child).Name())
	// switch child := child.(type) {
	// case string:
	// 	return i.Proto.Get(child)
	// }
	return nil
}

func (i Int) Equal(other IType) bool {
	switch other := other.(type) {
	case *Int:
		return i.Value == other.Value
	case *Float:
		return float64(i.Value) == other.Value
	}
	return false
}

func (i Int) Gt(other IType) bool {
	switch other := other.(type) {
	case *Int:
		return i.Value > other.Value
	case *Float:
		return float64(i.Value) > other.Value
	}
	return false
}

func (i Int) Gte(other IType) bool {
	switch other := other.(type) {
	case *Int:
		return i.Value >= other.Value
	case *Float:
		return float64(i.Value) >= other.Value
	}
	return false
}

func (i Int) Lt(other IType) bool {
	switch other := other.(type) {
	case *Int:
		return i.Value < other.Value
	case *Float:
		return float64(i.Value) < other.Value
	}
	return false
}

func (i Int) Lte(other IType) bool {
	switch other := other.(type) {
	case *Int:
		return i.Value <= other.Value
	case *Float:
		return float64(i.Value) <= other.Value
	}
	return false
}

func (i Int) Bool() bool {
	return i.Value != 0
}

func NewInt(v int) *Int {
	return &Int{Value: v, Proto: NewObject(ObjectValue{})}

}
