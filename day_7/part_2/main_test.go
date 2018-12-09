package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Start()
	}
}
