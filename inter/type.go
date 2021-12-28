package inter

const (
	INT = iota
	FLOAT
	BOOL
	STRING
	OBJECT
	FUNCTION
	NULL
)

type IType interface {
	TypeOf() int
	String() string
	Access(child IType) IType
	Equal(other IType) bool
	Gt(other IType) bool
	Gte(other IType) bool
	Lt(other IType) bool
	Lte(other IType) bool
	Bool() bool
}
