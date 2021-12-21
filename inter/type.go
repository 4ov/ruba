package inter

const (
	INT = iota
	STRING
	OBJECT
	FUNCTION
)

type IType interface {
	TypeOf() int
	String() string
	Access(child IType) IType
}
