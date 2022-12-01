package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

type Inventory []int

func (inv Inventory) Sum() int {
	s := 0
	for _, v := range inv {
		s += v
	}

	return s
}

type Elves []Inventory

func (e *Elves) Add(inv Inventory) {
	*e = append(*e, inv)
}

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
