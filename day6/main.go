package main

import (
	"fmt"
	"log"
	"os"
)

type Scanner struct {
	input    []byte
	position int
}

func NewScanner(b []byte) Scanner {
	return Scanner{
		input:    b,
		position: 0,
	}
}

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

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := NewScanner(f)
	fmt.Println(s.Scan(4))
	fmt.Println(s.Scan(14))
}
