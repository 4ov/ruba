package inter

type Null struct {
	IType
}

func (n Null) TypeOf() int {
	return NULL
}

func (n Null) String() string {
	return "null"
}

func (n Null) Access(child IType) IType {
	return NewNull()
}

func NewNull() *Null {
	return &Null{}
}
