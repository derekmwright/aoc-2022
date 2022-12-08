package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"text/scanner"
)

const (
	COMMAND = "$"
	CHDIR   = "cd"
	LIST    = "ls"
	DIR     = "dir"
)

type fs struct {
	directories map[string]int
	current     string
}

var sizeRegex = regexp.MustCompile(`^\d+$`)

func PartOne(input io.Reader) int {
	var s scanner.Scanner
	s.Init(input)

	sum := 0
	f := &fs{}
	f.directories = make(map[string]int)

	_ = regexp.MustCompile(`\d+`)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch s.TokenText() {
		case COMMAND:
			s.Scan()
			switch s.TokenText() {
			case CHDIR:
				s.Scan()
				if s.TokenText() == "." && s.Peek() == '.' {
					s.Scan()
					lastDir := f.current
					f.current = path.Join(f.current, "..")
					f.directories[f.current] += f.directories[lastDir]
				} else {
					f.current = path.Join(f.current, s.TokenText())
				}
			case LIST:
				cancel := false
				for !cancel {
					switch s.TokenText() {
					case DIR:
						s.Scan()
						f.directories[f.current] = 0
					default:
						cancel = true
					}
				}
			}
		case sizeRegex.FindString(s.TokenText()):
			size, err := strconv.Atoi(s.TokenText())
			if err != nil {
				log.Fatalln("invalid size value")
			}
			f.directories[f.current] += size
		default:
		}
	}

	for _, dir := range f.directories {
		if dir <= 100000 {
			sum += dir
		}
	}

	return sum
}

func main() {
	fname := "input.txt"
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(PartOne(f))
}
