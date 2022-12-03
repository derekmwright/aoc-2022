package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)

	sum := 0
	for s.Scan() {
		sum += Score1(s.Text())
	}

	fmt.Println(sum)
}
