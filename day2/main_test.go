package main

import "testing"

func BenchmarkMaths(b *testing.B) {
	t := "A Z"
	for i := 0; i < b.N; i++ {
		Score(t, 2)
	}
}

func BenchmarkLookup(b *testing.B) {
	t := "A Z"
	for i := 0; i < b.N; i++ {
		Play(t, LookupP1, ValueP1)
	}
}
