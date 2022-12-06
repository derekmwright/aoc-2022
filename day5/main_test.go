package main

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name string
		have string
		want []byte
	}{
		{
			name: "check for empty crate in row",
			have: "[A] [B]    [D]\n 1  2  3  4 \n\nmove 1 from 1 to 2\nmove 2 from 2 to 1",
			want: []byte{65, 66, 0, 68},
		},
		// {
		// 	name: "single crate",
		// 	have: "[A]",
		// 	want: []byte{65},
		// },
		// {
		// 	name: "check for empty crate at start and crate at end",
		// 	have: "    [A]",
		// 	want: []byte{0, 65},
		// },
		// {
		// 	name: "check ordering of ints doesnt change",
		// 	have: "[B] [A] [D] [C]",
		// 	want: []byte{66, 65, 68, 67},
		// },
		// {
		// 	name: "check for multiple lines",
		// 	have: "[B] [A]    [D]\n[C] [E] [F] [G]",
		// 	want: []byte{65, 66, 0, 68},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.have)
			got := l.Scan()
			fmt.Printf("%+v\n", got.Execute(false))
		})
	}
}
