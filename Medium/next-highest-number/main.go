package main

import (
	"fmt"
	"sort"
	"strconv"
)

func nextHigherNumber(num int) int {
	numStr := []rune(strconv.Itoa(num))
	n := len(numStr)

	// Step 1: Find the pivot
	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if numStr[i] < numStr[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		return -1 // No higher permutation is possible
	}

	// Step 2: Find the successor
	successor := -1
	for i := n - 1; i > pivot; i-- {
		if numStr[i] > numStr[pivot] {
			successor = i
			break
		}
	}

	// Step 3: Swap the pivot and successor
	numStr[pivot], numStr[successor] = numStr[successor], numStr[pivot]

	// Step 4: Sort the suffix and remove duplicates
	suffix := numStr[pivot+1:]
	sort.Slice(suffix, func(i, j int) bool {
		return suffix[i] < suffix[j]
	})

	// Create a new slice for unique digits
	uniqueSuffix := []rune{}
	seen := make(map[rune]bool)

	for _, char := range suffix {
		if !seen[char] {
			seen[char] = true
			uniqueSuffix = append(uniqueSuffix, char)
		}
	}

	// Combine the results
	result := append(numStr[:pivot+1], uniqueSuffix...)

	// Convert back to integer
	resultNum, err := strconv.Atoi(string(result))
	if err != nil {
		return -1
	}

	return resultNum
}

func main() {
	// fmt.Println(nextHigherNumber(1234)) // Output: 1243
	// fmt.Println(nextHigherNumber(1223)) // Output: 1232
	// fmt.Println(nextHigherNumber(321))  // Output: -1
	fmt.Println(nextHigherNumber(534978)) // Output: 536479
}
