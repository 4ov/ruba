package exec

func increaseValue(value Type) Type {

	switch {
	case value.TypeOf() == NUMBER:
		return value.(Number).Add(NewInt(1))
	// case value.TypeOf() == STRING:
	default:
		return NewInt(1)
	}

	// return value
}
