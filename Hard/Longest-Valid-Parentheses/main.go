package main

import (
	"fmt"
)

func longestValidBrackets(s string) int {
	max := 0
	match := map[byte]byte{
		'(': ')', '[': ']', '{': '}',
	}
	open := map[byte]struct{}{
		'(': {}, '{': {}, '[': {},
	}
	stack := []int{-1} // Sentinel index to calculate lengths correctly

	for i := 0; i < len(s); i++ {
		if _, ok := open[s[i]]; ok { // Opening bracket
			stack = append(stack, i)
		} else { // Closing bracket
			if len(stack) > 1 && match[s[stack[len(stack)-1]]] == s[i] {
				stack = stack[:len(stack)-1]
				currentLen := i - stack[len(stack)-1]
				if currentLen > max {
					max = currentLen
				}
			} else {
				stack = append(stack, i) // Reset base index for unmatched closing
			}
		}
	}
	return max
}

func main() {
	testCases := []struct {
		input string
		want  int
	}{
		// {"{()[{}]}", 8},
		// {"}{}()", 4},
		// {"(([]))[]{}", 10},
		{"([)]", 2},
		{"", 0},
	}

	for _, tc := range testCases {
		got := longestValidBrackets(tc.input)
		fmt.Printf("Input: %q, Longest valid length: %d (Expected: %d)\n", tc.input, got, tc.want)
	}
}
