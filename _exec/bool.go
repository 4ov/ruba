package exec

import "fmt"

type Bool struct {
	Type
	Value     bool
	Prototype *Object
}

func (v Bool) String() string {
	return fmt.Sprint(v.Value)
}

func (v Bool) TypeOf() int {
	return BOOL
}

func (b Bool) True() bool {
	return b.Value
}

func NewBool(value bool) *Bool {
	result := &Bool{
		Value: value,
	}

	return result
}
