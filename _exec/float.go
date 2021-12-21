package exec

import "fmt"

type Float struct {
	Type
	Number
	Value     float64
	Prototype *Object
}

func (i Float) String() string {
	return fmt.Sprint(i.Value)
}

func (i Float) TypeOf() int {
	return NUMBER
}

func (f Float) Chain([]string) Type {
	return NullValue
}

func (f Float) Access(child string) Type {

	return f.Prototype.Get(child)
}

func (f Float) True() bool {
	if f.Value == 0 {
		return false
	} else {
		return true
	}
}

func (i Float) Add(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewFloat(l.Value + r.Value)
	case *Int:
		return NewFloat(l.Value + float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Sub(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewFloat(l.Value - r.Value)
	case *Int:
		return NewFloat(l.Value - float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Div(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewFloat(l.Value / r.Value)
	case *Int:
		return NewFloat(l.Value / float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Mul(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewFloat(l.Value * r.Value)
	case *Int:
		return NewFloat(l.Value * float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Eq(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewBool(l.Value == r.Value)
	case *Int:
		return NewBool(l.Value == float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Gt(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewBool(l.Value > r.Value)
	case *Int:
		return NewBool(l.Value > float64(r.Value))
	default:
		return NewNull()
	}
}

func (i Float) Gte(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Float:
		return NewBool(l.Value >= r.Value)
	case *Int:
		return NewBool(l.Value >= float64(r.Value))
	default:
		return NewNull()
	}
}

func NewFloat(value float64) *Float {
	result := &Float{
		Value: value,
	}

	return result
}
