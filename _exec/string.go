package exec

type String struct {
	Type
	Value     string
	Prototype *Object
}

func (v String) String() string {
	return v.Value
}

func (v String) TypeOf() int {
	return STRING
}

func (v String) True() bool {
	return len(v.Value) > 0
}

func (v String) Access(child string) Type {
	value := v.Prototype.Get(child)
	return value
}

func NewString(value string) *String {
	result := &String{
		Value: value,
	}

	proto := NewObject(ObjectValue{})

	proto.Set("length", NewInt(len(result.Value)))

	result.Prototype = proto

	return result
}
