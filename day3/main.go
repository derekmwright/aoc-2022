package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func Score1(sack string) int {
	l := len([]byte(sack)) / 2

	for _, v := range sack[0:l] {
		if strings.Contains(sack[l:], string(v)) {
			if int(v) >= 65 && int(v) <= 90 {
				return int(v) - 38
			}
			return int(v) - 96
		}
	}

	return 0
}

func uniq(in string) string {
	s := &strings.Builder{}
	for _, sub := range in {
		if !strings.Contains(s.String(), string(sub)) {
			s.WriteRune(sub)
		}
	}
	return s.String()
}

func Score2(group []string) int {
	for _, v := range group[0] {
		if strings.Contains(group[1], string(v)) && strings.Contains(group[2], string(v)) {
			if int(v) >= 65 && int(v) <= 90 {
				return int(v) - 38
			}
			return int(v) - 96
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
		if group == 3 {
			// Sort the slice by length so we only need to iterate the short ones first
			sort.Slice(groupSack, func(i, j int) bool {
				return len(groupSack[i]) < len(groupSack[j])
			})
			sum2 += Score2(groupSack)
			groupSack = make([]string, 0)
			group = 0
		}
		sum1 += Score1(s.Text())
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
