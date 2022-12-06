package main

import (
	"fmt"
	"log"
	"os"
)

type Window map[byte]int

func (w Window) add(b byte) {
	if _, ok := w[b]; ok {
		w[b] += 1
		return
	}
	w[b] = 1
}

func (w Window) delete(b byte) {
	if i, ok := w[b]; ok {
		if i-1 <= 0 {
			delete(w, b)
			return
		}
		w[b] -= 1
	}
}

type Scanner struct {
	input    []byte
	position int
	window   Window
}

func NewScanner(b []byte) Scanner {
	return Scanner{b, 0, make(Window, 0)}
}

// Scan finds a series of unique characters within the specified size.
// Returns the index after the unique set is found.
func (s Scanner) Scan(size int) int {
	for s.position < (len(s.input) - size) {
		// set map to hold unique keys
		m := make(map[byte]struct{}, 0)
		for i := 0; i < size; i++ {
			m[s.input[s.position+i]] = struct{}{}
		}

		if len(m) == size {
			return s.position + size
		}

		s.position += 1
	}
	return -1
}

func (s Scanner) ScanNew(size int) int {
	for s.position < (len(s.input) - size) {
		s.window.add(s.input[s.position])
		if s.position >= size {
			s.window.delete(s.input[s.position-size])
			if len(s.window) == size {
				return s.position + 1
			}
		}
		s.position += 1
	}
	return -1
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := NewScanner(f)
	fmt.Println(s.ScanNew(4))
	fmt.Println(s.ScanNew(14))
}
