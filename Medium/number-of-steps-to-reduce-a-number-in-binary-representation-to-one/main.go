package main

import (
	"fmt"
)

// addOne adds 1 to the binary string and returns the new string.
func addOne(s []rune) string {
	// Traverse from the rightmost bit
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			// If the current bit is '0', flip it to '1'
			s[i] = '1'
			break
		} else {
			// If the current bit is '1', flip it to '0' and continue the carry
			s[i] = '0'
		}
	}
	// If the first character is '0', prepend a '1' to the string
	if s[0] == '0' {
		s = append([]rune{'1'}, s...)
	}
	return string(s)
}

// numSteps reduces the binary number to 0 and returns the number of steps.
func numSteps(s string) int {
	ans := 0
	// Convert the binary string into a number
	for s != "1" {
		ans++
		// If the number is even (ends in '0'), divide by 2
		if s[len(s)-1] == '0' {
			s = s[:len(s)-1]
		} else {
			// If the number is odd, add 1 to make it even
			s = addOne([]rune(s))
		}
	}
	return ans
}

func main() {
	// Test the function with a binary string
	binaryStr := "1110"              // Example input
	fmt.Println(numSteps(binaryStr)) // Output should be the number of steps to reduce to zero
}
