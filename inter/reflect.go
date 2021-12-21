package inter

type ReflectiveValue = func() IType

type Reflective struct {
	IType
	Value ReflectiveValue
}

func (r Reflective) String() string {
	return r.Value().String()
}

func (r Reflective) TypeOf() int {
	return r.Value().TypeOf()
}
