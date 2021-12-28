package inter

import "fmt"

type Bool struct {
	IType
	Value bool
	Proto *Object
}

func (b Bool) TypeOf() int {
	return BOOL
}

func (b Bool) String() string {
	return fmt.Sprintf("%v", b.Value)
}

func (b Bool) Access(child IType) IType {
	return b.Proto.Access(child)
}

func (b Bool) Equal(other IType) bool {
	switch other := other.(type) {
	case *Bool:
		return b.Value == other.Value
	case *Int:
		return other.Value == 0 && !b.Value
	case *Float:
		return other.Value == 0 && !b.Value
	default:
		return false
	}
}

func (b Bool) Gt(other IType) bool {
	return false
}
func (b Bool) Gte(other IType) bool {
	return false
}
func (b Bool) Lte(other IType) bool {
	return false
}

func (b Bool) Lt(other IType) bool {
	return false
}
func (b Bool) Bool() bool {
	return b.Value
}

func NewBool(v bool) *Bool {
	return &Bool{Value: v, Proto: NewObject(ObjectValue{})}
}
