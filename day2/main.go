package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROCK = iota
	PAPER
	SCISSORS
)

// ROCK > SCISSORS > PAPER > ROCK...

type Round [2]byte

func NewRound(p1, p2 byte) Round {
	return Round([2]byte{p1, p2})
}

// Part 1 Play Results Table
//
//	AX AY AZ - 3 6 0
//	BX BY BZ - 0 3 6
//	CX CY CZ - 6 0 3
var LookupP1 = map[string]int{
	"AX": 3,
	"AY": 6,
	"AZ": 0,
	"BX": 0,
	"BY": 3,
	"BZ": 6,
	"CX": 6,
	"CY": 0,
	"CZ": 3,
}

// Part 2 Result Table
var LookupP2 = map[string]int{
	"AX": 3,
	"AY": 1,
	"AZ": 2,
	"BX": 1,
	"BY": 2,
	"BZ": 3,
	"CX": 2,
	"CY": 3,
	"CZ": 1,
}

// X, Y, Z indicates which to play: Rock, Paper, or Scissors
var ValueP1 = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

// X, Y, Z indicates loss, draw, or win
var ValueP2 = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

// Play will return the score of the play for part1
func Play(p1, p2 byte, lookup map[string]int, bonus map[string]int) int {
	p := &strings.Builder{}
	p.WriteByte(p1)
	p.WriteByte(p2)

	return lookup[p.String()] + bonus[string(p2)]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)
	sum1 := 0
	sum2 := 0
	for s.Scan() {
		match := s.Text()
		sum1 += Play(match[0], match[2], LookupP1, ValueP1)
		sum2 += Play(match[0], match[2], LookupP2, ValueP2)
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
