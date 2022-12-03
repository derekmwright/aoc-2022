package main

import "testing"

func BenchmarkScore(b *testing.B) {
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

func BenchmarkScore2(b *testing.B) {
	t := "A Z"
	for i := 0; i < b.N; i++ {
		Score2(t, 2)
	}
}

func BenchmarkLookup2(b *testing.B) {
	t := "A Z"
	for i := 0; i < b.N; i++ {
		Play(t, LookupP2, ValueP2)
	}
}
