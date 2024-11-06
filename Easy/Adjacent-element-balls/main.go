package main
// Rearrange the elements(balls) such that adjacent elements have an empty
// Sample - ["b","","","b"] to ["b","","b",""]

import (
	"fmt"
)

func rearrangeBalls(arr []string) []string {
	// Count non-empty elements (balls)
	balls := []string{}
	for _, v := range arr {
		if v != "" {
			balls = append(balls, v)
		}
	}

	n := len(arr)
	m := len(balls)

	// If all elements are empty or all are filled, return the original array
	if m == 0 || m == n {
		return arr
	}

	// Create a new array for the result
	result := make([]string, n)

	// Place balls in every other position
	for i := 0; i < m && i*2 < n; i++ {
		result[i*2] = balls[i]
	}

	return result
}

func main() {
	// Test cases
	testCases := [][]string{
		{"b", "", "", "b"},
		{"b", "b", "b", ""},
		{"", "b", "b", "b"},
		{"b", "b", "b", "b"},
		{"", "", "", ""},
	}

	for _, test := range testCases {
		fmt.Printf("Input:  %v\n", test)
		fmt.Printf("Output: %v\n\n", rearrangeBalls(test))
	}
}