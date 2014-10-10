package smtest

import (
	"testing"
)

func BenchmarkGotoSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runGotoSimple()
	}
}

func BenchmarkGotoFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runGotoFunctions()
	}
}

func BenchmarkFuncFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runFuncFunc()
	}
}

func BenchmarkStructFuncFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runStructFuncFunc()
	}
}
