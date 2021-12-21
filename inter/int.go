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

func NewInt(v int) *Int {
	return &Int{Value: v, Proto: NewObject(ObjectValue{})}
}
