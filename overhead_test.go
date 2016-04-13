package cgo_overhead

import (
	"testing"
)

func BenchmarkCFuncall(b *testing.B) {
	f := Fp(FC)
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkGoFuncall(b *testing.B) {
	f := Fp(FGo)
	for i := 0; i < b.N; i++ {
		f()
	}
}
