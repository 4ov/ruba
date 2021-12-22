package inter

const (
	INT = iota
	STRING
	OBJECT
	FUNCTION
	NULL
)

type IType interface {
	TypeOf() int
	String() string
	Access(child IType) IType
}
