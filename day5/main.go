package main

import "fmt"

// [A]     [B]
// [C] [D] [E]
//  1   2   3
//
// If read from top to bottom
//
// stacks = [['A','C'],['','D'],['B','E']]
//
// from 1 move 1 to 2
//
// stacks = [['C'],['A','D']['B','E']]

const (
	SCRATE = '['
	ECRATE = ']'
	SPACE  = ' '
)

func ScanCrates(in string) {
	collect := []byte{}
	for _, r := range in {
		switch r {
		case SCRATE:
		}
	}
	fmt.Println(in)
}

func main() {

}
