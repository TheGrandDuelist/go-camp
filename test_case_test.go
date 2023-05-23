package main

import "testing"

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		map_alloc()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice_alloc()
	}
}
