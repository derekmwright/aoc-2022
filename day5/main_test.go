package main

import "testing"

func TestLexer(t *testing.T) {
	input := "[A] [B]    [D]"
	ScanCrates(input)
}
