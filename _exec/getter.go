package exec

type Reflective = func() Type

type Getter struct {
	Type
	Value Reflective
}

func (g Getter) String() string {
	return g.Value().String()
}

func (g Getter) TypeOf() int {
	return g.Value().TypeOf()
}

func NewGetter(value Reflective) *Getter {

	result := &Getter{
		Value: value,
	}

	return result
}
