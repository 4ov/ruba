package exec

type ObjectValue = map[string]Type

type Object struct {
	Type
	Value     ObjectValue
	ProtoType *Object
}

func (o Object) String() string {
	return "{}"
}

func (o Object) TypeOf() int {
	return OBJECT
}

func (o Object) True() bool {
	return true
}

func (o *Object) Set(key string, val Type) {
	o.Value[key] = val
}

func (o *Object) Has(key string) bool {
	if _, ok := o.Value[key]; ok {
		return true
	} else {
		return false
	}
}

func (o *Object) Get(key string) Type {
	value := o.Value[key]
	// fmt.Println(reflect.TypeOf(value))
	if value != nil {
		return value
	} else {
		return nil
	}
}

func (o *Object) Access(key string) Type {
	return o.Get(key)
}

func NewObject(value ObjectValue) *Object {

	result := &Object{
		Value: value,
	}

	return result
}
