package vcl

type IObject interface {
	Instance() uintptr
	ClassName() string
	Free()
	GetHashCode() int32
	Equals(IObject) bool
	IsValid() bool
}
