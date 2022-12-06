package main

import (
	"os"
	"testing"
)

func BenchmarkScan4(b *testing.B) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		b.Error(err)
	}

	s := NewScanner(f)

	for i := 0; i < b.N; i++ {
		s.Scan(4)
	}
}

func BenchmarkScanNew4(b *testing.B) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		b.Error(err)
	}

	s := NewScanner(f)

	for i := 0; i < b.N; i++ {
		s.ScanNew(4)
	}
}

func BenchmarkScan14(b *testing.B) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		b.Error(err)
	}

	s := NewScanner(f)

	for i := 0; i < b.N; i++ {
		s.Scan(14)
	}
}

func BenchmarkScanNew14(b *testing.B) {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		b.Error(err)
	}

	s := NewScanner(f)

	for i := 0; i < b.N; i++ {
		s.ScanNew(14)
	}
}
