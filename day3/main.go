package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// uniq returns only uniq characters found in string.
func uniq(in string) string {
	s := &strings.Builder{}
	for _, sub := range in {
		if !strings.Contains(s.String(), string(sub)) {
			s.WriteRune(sub)
		}
	}
	return s.String()
}

// scoreRune takes a rune and applies the scoring rules for the puzzle to it.
//
//	a-z are scored 1-26
//	A-Z are scored 27-52
//
// However ASCII char values are reversed and `A-Z` is a lower val than `a-z`.
//
//	A-Z are 65-90
//	a-z are 97-122
//
// To get the proper value of `a` we can subtract 96 from is which sets its value back to 1
// To get the proper value of `A` we can offset the ASCII value by 38 which sets it back to 27
func scoreRune(r rune) int {
	if int(r) > 96 {
		return int(r) - 96
	}
	return int(r) - 38
}

// Score1 takes a string and splits it in half.
// Then it iterates over the first string for each rune and checks to see if the 2nd half contains it.
// It returns as soon as a match is found.
// If no match is found it returns 0, this should never occur.
func Score1(sack string) int {
	l := len([]byte(sack)) / 2

	for _, v := range sack[0:l] {
		if strings.Contains(sack[l:], string(v)) {
			return scoreRune(v)
		}
	}
	return 0
}

// Score2 takes a string that is a concatenation of the uniqued 3 rucksacks.
// Then each rune is checked to see if it occurs 3 times.
// As soon as the condition is met, the value is scored and returned.
// A 0 is returned if a condition is never met, this should never occur.
func Score2(group string) int {
	for _, v := range group {
		if strings.Count(group, string(v)) == 3 {
			return scoreRune(v)
		}
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)
	sum1 := 0
	sum2 := 0
	group := 0
	groupSack := make([]string, 0)

	for s.Scan() {
		// Compress strings using uniq func
		groupSack = append(groupSack, uniq(s.Text()))
		group++

		if (group % 3) == 0 {
			// Join the 3 strings into a single string and score it
			j := strings.Join(groupSack, "")
			sum2 += Score2(j)

			// Reset groupSack to hold next 3 iterations
			groupSack = make([]string, 0)
		}

		sum1 += Score1(s.Text())
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
