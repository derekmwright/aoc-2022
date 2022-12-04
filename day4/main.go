package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range [2]int

// NewRange creates a range from a hyphen delimited set of integers.
// Example:
//
//	"2-6" -> Range{2,6}
//
// Panics if there are errors converting str to int.
func NewRange(rng string) Range {
	s := strings.Split(rng, "-")

	start, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}

	end, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}

	return Range{start, end}
}

// contains returns true if the provided range is within the canonoical range.
// Example:
//
//	Range{2,6}.contains(Range{2,4}) -> true
//	Range{1,2}.contains(Range{5,6}) -> false
func (r Range) contains(in Range) bool {
	return in[0] >= r[0] && in[1] <= r[1]
}

// overlaps returns true if the provided range's boundary overlaps the canonical range.
// Example:
//
//	Range{2,4}.overlaps(Range{4,6}) -> true
//	Range{1,4}.overlaps(Range{6,10}) -> false
func (r Range) overlaps(in Range) bool {
	return (in[0] >= r[0] && in[0] <= r[1]) || (in[1] >= r[0] && in[1] <= r[1])
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
		r := strings.Split(s.Text(), ",")
		r1 := NewRange(r[0])
		r2 := NewRange(r[1])
		if r1.contains(r2) || r2.contains(r1) {
			sum1++
		}
		if r1.overlaps(r2) || r2.overlaps(r1) {
			sum2++
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
