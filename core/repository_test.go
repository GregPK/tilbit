package core

import (
	"testing"
)

func Benchmark_AllTilbits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AllTilbits()
	}
}
