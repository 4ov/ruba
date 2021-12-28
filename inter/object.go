package inter

import (
	"fmt"
	"reflect"
)

type ObjectValue = map[string]IType

type Object struct {
	IType
	Value ObjectValue
	Proto *Object
}

func (o Object) TypeOf() int {
	return OBJECT
}

func (o Object) String() string {
	return fmt.Sprintf("Object(%v)", o.Value)
}

func (o *Object) Set(key string, value IType) {
	o.Value[key] = value
}

// func (o *Object) ChainGet(chain []IType) IType {
// 	fmt.Println(chain)
// 	return nil
// }

func (o Object) Equal(other IType) bool {
	switch other := other.(type) {
	case *Object:
		return reflect.DeepEqual(o.Value, other.Value)
	}
	return false
}

func (o Object) Gt(other IType) bool {
	return false
}

func (o Object) Gte(other IType) bool {
	return false
}

func (o Object) Lt(other IType) bool {
	return false
}

func (o Object) Lte(other IType) bool {
	return false
}

func (o Object) Bool() bool {
	return true
}

func (o Object) Access(child IType) IType {

	switch child := child.(type) {
	case *String:
		return o.Get(child.Value)
	default:
		panic("Object.Access: invalid child type")
	}
}

func (o *Object) Assign(chain []IType, value IType) IType {

	if len(chain) == 1 {
		if chain[0].TypeOf() == STRING {
			o.Set(chain[0].String(), value)
		} else {
			panic("Cannot set non-string key")
		}
	} else {
		nestAssign(chain[0], chain[1:], value)
	}
	return nil
}

func (o *Object) Get(key string) IType {
	if v, ok := o.Value[key]; ok {
		return v
	} else {
		return nil
	}
}

func NewObject(v ObjectValue) *Object {
	return &Object{
		Value: v,
	}
}
