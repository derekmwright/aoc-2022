package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

// Advent of Code Day 1
//
// Things I like:
// - How the data is parsed and the data structures setup.
// Things I'd like to improve upon:
// - I think I could do the first part with a little better summing concurrency
// - Not too proud of the 2nd part as I haven't thought up a good algo for retaining the top 3 values. Had to just finish it so I could start work, but may research options/methods for handling this scenario w/ better memory utilization.

// Inventory represents the calories carried by an individual Elf.
type Inventory []int

// Sum returns the total caloric value of the elf's inventory.
func (inv Inventory) Sum() int {
	s := 0
	for _, v := range inv {
		s += v
	}

	return s
}

// Elves represents a group of elves who each have an inventory of caloric values.
type Elves []Inventory

// Add will append an elf to the group with its inventory of calories.
func (e *Elves) Add(inv Inventory) {
	*e = append(*e, inv)
}

// ScanInv takes raw caloric values and converts it into individual Elves, modifies the group of elves in place.
func (e *Elves) ScanInv(r io.Reader) {
	s := bufio.NewScanner(r)

	elf := make(Inventory, 0)

	for s.Scan() {
		if s.Text() == "" {
			e.Add(elf)
			elf = make(Inventory, 0)
			continue
		}

		cal, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		elf = append(elf, cal)
	}
}

// Solves for part 1
func Part1(elves Elves) int {
	sums := make(chan int, len(elves))
	go func() {
		for _, elf := range elves {
			sums <- elf.Sum()
		}
		close(sums)
	}()

	high := 0
	for sum := range sums {
		if sum > high {
			high = sum
		}
	}

	return high
}

// Solves for part 2
func Part2(elves Elves) int {
	rankings := make([]int, 0)

	for _, elf := range elves {
		rankings = append(rankings, elf.Sum())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(rankings)))

	sum := 0
	for _, v := range rankings[0:3] {
		sum += v
	}

	return sum
}

func main() {
	elves := make(Elves, 0)
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	elves.ScanInv(r)

	fmt.Printf("Part 1: %d\n", Part1(elves))
	fmt.Printf("Part 2: %d\n", Part2(elves))
}
