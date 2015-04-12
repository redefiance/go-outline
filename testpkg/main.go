package testpkg

// IsAnInt has a
// multiline comment
type IsAnInt int

type IsAnArrayOfComplex128 []complex128

type IsAChanOfBool chan bool

type IsAnInterface interface {
	Foo() []byte
}

func NewStruct(a, b, c int) *IsAStruct {
	return &IsAStruct{a, b, "neuneuneu!", nil}
}

var A = make(IsAChanOfBool)

const (
	FOO = iota
	FOOOO
	FOOOOO
)

type IsAPointerToRune *rune

type IsAStruct struct {
	A, B int
	C    string
	D    chan chan IsAnArrayOfComplex128
}

type IsAFunc func(in1, in2 string) (out string, err error)

const notExported = "i am private"
