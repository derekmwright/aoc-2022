package main

import (
	"testing"
)

func TestScoreRune(t *testing.T) {
	var tests = []struct {
		have rune
		want int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}

	for _, tt := range tests {
		t.Run("test that score rune returns the proper value", func(t *testing.T) {
			got := scoreRune(tt.have)
			if got != tt.want {
				t.Errorf("scoreRune(%s) got %v , want %v", string(tt.have), got, tt.want)
			}
		})
	}
}
