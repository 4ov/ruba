package exec

import "fmt"

type Int struct {
	Type
	Number
	Value     int
	IsNaN     bool
	Prototype *Object
}

func (i Int) String() string {
	if !i.IsNaN {
		return fmt.Sprint(i.Value)
	} else {
		return "NaN"
	}
}

func (i Int) TypeOf() int {
	return NUMBER
}

func (i Int) Chain([]string) Type {
	return NullValue
}

func (i Int) Access(child string) Type {
	return i.Prototype.Get(child)
}

func (i Int) True() bool {
	if i.Value == 0 {
		return false
	} else {
		return true
	}
}

//Number related
func (i Int) Add(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewInt(l.Value + r.Value)
	case *Float:
		return NewFloat(float64(l.Value) + r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Sub(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewInt(l.Value - r.Value)
	case *Float:
		return NewFloat(float64(l.Value) - r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Div(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewInt(l.Value / r.Value)
	case *Float:
		return NewFloat(float64(l.Value) / r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Mul(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewInt(l.Value * r.Value)
	case *Float:
		return NewFloat(float64(l.Value) * r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Eq(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewBool(l.Value == r.Value)
	case *Float:
		return NewBool(float64(l.Value) == r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Gt(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewBool(l.Value > r.Value)
	case *Float:
		return NewBool(float64(l.Value) > r.Value)
	default:
		return NewNull()
	}
}

func (i Int) Gte(to Number) Type {
	l := i
	r := to
	switch r := r.(type) {
	case *Int:
		return NewBool(l.Value >= r.Value)
	case *Float:
		return NewBool(float64(l.Value) >= r.Value)
	default:
		return NewNull()
	}
}

func NewInt(value int) *Int {
	proto := NewObject(ObjectValue{})

	result := &Int{
		Value:     value,
		IsNaN:     false,
		Prototype: proto,
	}

	return result
}

func NewNaN() *Int {
	result := &Int{
		Value: 0,
		IsNaN: true,
	}

	return result
}
