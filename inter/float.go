package inter

import "fmt"

type Float struct {
	IType
	Value float64
	Proto *Object
}

func (f Float) TypeOf() int {
	return FLOAT
}

func (f Float) String() string {
	return fmt.Sprintf("%v", f.Value)
}

func (f Float) Access(child IType) IType {
	switch child := child.(type) {
	case *Int:
		return NewFloat(float64(child.Value))
	case *String:
		return f.Proto.Access(child)
	default:
		panic("Float.Access: invalid child type")
	}
}

func (f Float) Equal(other IType) bool {
	switch other := other.(type) {
	case *Float:
		return f.Value == other.Value
	case *Int:
		return f.Value == float64(other.Value)
	default:
		return false
	}
}

func (f Float) Gt(other IType) bool {
	switch other := other.(type) {
	case *Float:
		return f.Value > other.Value
	case *Int:
		return f.Value > float64(other.Value)
	default:
		return false
	}
}

func (f Float) Gte(other IType) bool {
	switch other := other.(type) {
	case *Float:
		return f.Value >= other.Value
	case *Int:
		return f.Value >= float64(other.Value)
	default:
		return false
	}
}

func (f Float) Lt(other IType) bool {
	switch other := other.(type) {
	case *Float:
		return f.Value < other.Value
	case *Int:
		return f.Value < float64(other.Value)
	default:
		return false
	}
}

func (f Float) Lte(other IType) bool {
	switch other := other.(type) {
	case *Float:
		return f.Value <= other.Value
	case *Int:
		return f.Value <= float64(other.Value)
	default:
		return false
	}
}

func (f Float) Bool() bool {
	return f.Value != 0
}

func NewFloat(v float64) *Float {
	return &Float{Value: v, Proto: NewObject(ObjectValue{})}
}
