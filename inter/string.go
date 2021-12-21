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
	}
	return nil
}

func NewString(v string) *String {
	return &String{Value: v, Proto: NewObject(ObjectValue{
		"len": NewInt(len(v)),
	})}
}
