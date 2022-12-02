// First pass early in the morning used lookup tables but I figured there was a better way using maths.
// In my free time later in the day I started working on the forumulas for pure math implementation and then added some benchmarks to see the differences.
// Even though both had 0 allocations, the math version outperformed the lookups significantly.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part 1 Play Results Table
//
//	AX AY AZ - 3 6 0
//	BX BY BZ - 0 3 6
//	CX CY CZ - 6 0 3
var LookupP1 = map[string]int{
	"A X": 3,
	"A Y": 6,
	"A Z": 0,
	"B X": 0,
	"B Y": 3,
	"B Z": 6,
	"C X": 6,
	"C Y": 0,
	"C Z": 3,
}

// Part 2 Result Table
var LookupP2 = map[string]int{
	"A X": 3,
	"A Y": 1,
	"A Z": 2,
	"B X": 1,
	"B Y": 2,
	"B Z": 3,
	"C X": 2,
	"C Y": 3,
	"C Z": 1,
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

// Play will return the score of the play
func Play(match string, lookup map[string]int, bonus map[string]int) int {
	return lookup[match] + bonus[string(match[2])]
}

// Score calculates the score of the match according to the rules set forth in part one.
// result is the raw string in format of:
//
//	Their play: A, B or C for Rock, Paper, Scissors in the first byte
//	A space at 2nd byte
//	Our play: X, Y or Z for Rock, Paper, Scissors
//
// The result of a match can be calculated by using the difference between the int values of the chars used in the match.
// An offset is used to align the two different match notations between their play and our own.
// For example, knowing their are only 3 possible moves, their notation is "ABC", our notation is "XYZ".
//
// Offset can be calculated by taking the difference of the character values of the first byte of the notation and subtracting the lesser one from the greater and taking the modulo sum(number choices).
//
//	Example: X (88) - A (65) = 23 % 3 = 2
//
// The offset should be the same for the entire set of matches and only needs to be calculated once.
// After the offset is known, it can be used to align our values for comparison.
// In the traditional Rock, Paper, Scissors game; the defeatable item is +2 away, and the item it would be defeated by is +1 slot away.
// If rock is a value of 0, then rocks defeatable item is +2 away, skip paper (value 1), and scissors (value 2) would be defeated.
// This wraps around after choices have been exhausted, example: Paper (value 1), is defeated by Scissors (value 2), but defeats Rock (value 0).
// A modulo numChoices can be used to wrap values back around so that we can calculate if we won by using the following formula:
//
//	((me + offset) - them) % 3 // "A Z" is Them: Rock(65), Us: Scissors (90); ((90 + 2) - 65) = 27 % 3 = 0 // Loss
//
// This returns the scoring with the following values:
//
//	0 - loss
//	1 - tie
//	2 - win
//
// Scoring rules dictate that we get 0 for loss, 3 for tie, and 6 for a win, so all we need to do is multiply the result of our forumla by 3.
// Finally we just need to get the bonus points for the item we used and add that to the sum.
// We find the bonus amount by using the char value plus the offset then modulo 3 and add 1 (before one values are 0, 1, 2 and we need 1, 2, 3 per scoring rules).
func Score(result string, offset int) int {
	them := int(result[0])
	me := int(result[2])

	return (((me + offset) - them) % 3 * 3) + (((me + offset) % 3) + 1)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)

	offset := (int("X"[0]-"A"[0]) % 3)

	sum1 := 0
	sum2 := 0
	sum3 := 0

	for s.Scan() {
		match := s.Text()
		sum1 += Play(match, LookupP1, ValueP1)
		sum2 += Play(match, LookupP2, ValueP2)
		sum3 += Score(match, offset)
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
	fmt.Printf("Maths: %d\n", sum3)
}
