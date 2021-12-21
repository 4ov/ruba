package exec

const (
	NUMBER = iota
	STRING
	BOOL
	FN
	ARRAY
	OBJECT
	NULL
)

var (
	type_map = map[int]string{}

	NullValue  = NewNull()
	TrueValue  = NewBool(true)
	FalseValue = NewBool(false)
)

func init() {
	type_map[NUMBER] = "number"
	type_map[STRING] = "string"
	type_map[BOOL] = "bool"
	type_map[FN] = "function"
	type_map[ARRAY] = "array"
	type_map[NULL] = "null"
}

type Type interface {
	String() string
	TypeOf() int
	Access(string) Type
	True() bool
	// Chain([]string) Type
}

type Number interface {
	Type
	Add(Number) Type
	Sub(Number) Type
	Div(Number) Type
	Mul(Number) Type
	Eq(Number) Type
	Gt(Number) Type
	Gte(Number) Type
	Lt(Number) Type
	Lte(Number) Type
}
