package exec

type Null struct {
	Type
}

func (v Null) String() string {
	return "null"
}

func (v Null) TypeOf() int {
	return NULL
}

func (n Null) True() bool {
	return false
}

func NewNull() *Null {
	result := &Null{}

	return result
}
