package inter

func Add(l IType, r IType) IType {
	switch {
	case l.TypeOf() == INT && r.TypeOf() == INT:
		return NewInt(l.(*Int).Value + r.(*Int).Value)
	default:
		panic("unsupported operation")
	}
}

func Mul(l IType, r IType) IType {

	switch {
	case l.TypeOf() == INT && r.TypeOf() == INT:
		return NewInt(l.(*Int).Value * r.(*Int).Value)
	default:
		panic("unsupported operation")
	}
}
