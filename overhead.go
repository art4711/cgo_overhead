package overhead

// void f(void) { return; }
import "C"

func bar() {
}

const (
	FC = iota
	FGo
)

func Fp(which int) func() {
	switch which {
	case FC:
		return func() { C.f() }
	case FGo:
		return func() { bar() }
	}
	panic("ugh")
}