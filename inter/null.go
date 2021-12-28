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

func (n Null) Equal(other IType) bool {
	return other.TypeOf() == NULL
}

func (n Null) Gt(other IType) bool {
	return false
}

func (n Null) Gte(other IType) bool {
	return false
}

func (n Null) Lt(other IType) bool {
	return false
}

func (n Null) Lte(other IType) bool {
	return false
}

func (n Null) Bool() bool {
	return false
}

func NewNull() *Null {
	return &Null{}
}
