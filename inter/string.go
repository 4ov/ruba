package inter

type String struct {
	IType
	Value string
	Proto *Object
}

func (s String) TypeOf() int {
	return STRING
}

func (s String) String() string {
	return s.Value
}

func (s String) Access(child IType) IType {
	switch child := child.(type) {
	case *String:
		return s.Proto.Get(child.Value)
	case *Int:
		return NewString(string(s.Value[child.Value]))
	}
	return nil
}

func (s String) Equal(other IType) bool {
	switch other := other.(type) {
	case *String:
		return s.Value == other.Value
	}
	return false
}

func (s String) Gt(other IType) bool {
	switch other := other.(type) {
	case *String:
		return s.Value > other.Value
	}
	return false
}

func (s String) Gte(other IType) bool {
	switch other := other.(type) {
	case *String:
		return s.Value >= other.Value
	}
	return false
}

func (s String) Lt(other IType) bool {
	switch other := other.(type) {
	case *String:
		return s.Value < other.Value
	}
	return false
}

func (s String) Lte(other IType) bool {
	switch other := other.(type) {
	case *String:
		return s.Value <= other.Value
	}
	return false
}

func NewString(v string) *String {
	return &String{Value: v, Proto: NewObject(ObjectValue{
		"len": NewInt(len(v)),
	})}
}
