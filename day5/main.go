package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	SCRATE  = '['
	SPACE   = ' '
	NEWLINE = '\n'
	INSTART = 'm' // instructions start with a lowercase m
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

//
// Stack tracking
//

type Stack []byte

func NewStack() Stack {
	return make(Stack, 0)
}

func (s *Stack) addCrate(ident byte) {
	*s = append(*s, ident)
}

type Stacks []Stack

func (s Stacks) move(chunk bool, num, src, dst int) {
	if chunk {
		s[dst] = append(s[dst], s[src][(len(s[src])-num):len(s[src])]...)
		s[src] = s[src][:len(s[src])-num]
		return
	}
	for i := 0; i < num; i++ {
		s[dst] = append(s[dst], s[src][len(s[src])-1])
		s[src] = s[src][:len(s[src])-1]
	}
}

type Instruction struct {
	NumItems    int
	Source      int
	Destination int
}

func NewInstruction(num, src, dst int) *Instruction {
	return &Instruction{num, src, dst}
}

type Inventory struct {
	Stacks       Stacks
	Instructions []Instruction
}

func (i *Inventory) Execute(chunk bool) string {
	// Reverse the stacks so that we always move from the top (end of the slice), since we read from top to bottom.
	for _, s := range i.Stacks {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	for _, inst := range i.Instructions {
		i.Stacks.move(chunk, inst.NumItems, inst.Source, inst.Destination)
	}

	result := &strings.Builder{}

	for _, s := range i.Stacks {
		if len(s) > 0 {
			result.WriteByte(s[len(s)-1])
		}
	}

	return result.String()
}

// ScanLine scans a line out of the lexer and returns the stack that was parsed.
func (l *Lexer) Scan() *Inventory {
	s := make(Stacks, 10) // feels a lil cheaty but looking at the input we dont have more than 9 rows
	st := 0
	eof := false
	inst := make([]Instruction, 0)

	for {
		switch l.ch {
		case SCRATE:
			l.readChar()
			s[st].addCrate(l.ch)
			l.readChar()
			st += 1
		case SPACE:
			if l.peekChar() == SPACE {
				st += 1
				l.readPosition += 3
			}

			// skip junk line with stack numbers
			if st == 0 {
				if l.peekChar() >= 48 && l.peekChar() <= 57 {
					l.readPosition += strings.IndexByte(l.input[l.position:], byte('\n'))
				}
			}
		case NEWLINE:
			st = 0
		case INSTART:
			var inbuf *bytes.Buffer
			readN := strings.IndexByte(l.input[l.position:], byte('\n'))
			if readN == -1 {
				inbuf = bytes.NewBuffer([]byte(l.input[l.position:]))
				l.readPosition = len(l.input)
			} else {
				inbuf = bytes.NewBuffer([]byte(l.input[l.position : l.position+readN]))
				l.readPosition += readN
			}

			n, s, d := 0, 0, 0
			fmt.Fscanf(inbuf, "move %d from %d to %d\n", &n, &s, &d)
			inst = append(inst, *NewInstruction(n, s-1, d-1))
		case 0:
			if len(s[st]) != 0 {
				st = 0
			}
			eof = true
		}

		if eof {
			break
		}

		l.readChar()
	}

	return &Inventory{
		Stacks:       s,
		Instructions: inst,
	}
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Part One - Uses serial move
	l := NewLexer(string(f))
	inv := l.Scan()
	fmt.Println(inv.Execute(false))

	// Part Two - Uses chunked move
	l = NewLexer(string(f))
	inv = l.Scan()
	fmt.Println(inv.Execute(true))
}
