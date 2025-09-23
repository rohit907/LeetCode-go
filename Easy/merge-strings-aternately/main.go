package main

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

// Return the merged string.

 

// Example 1:

// Input: word1 = "abc", word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r
// Example 2:

// Input: word1 = "ab", word2 = "pqrs"
// Output: "apbqrs"
// Explanation: Notice that as word2 is longer, "rs" is appended to the end.
// word1:  a   b 
// word2:    p   q   r   s
// merged: a p b q   r   s
// Example 3:

// Input: word1 = "abcd", word2 = "pq"
// Output: "apbqcd"
// Explanation: Notice that as word1 is longer, "cd" is appended to the end.
// word1:  a   b   c   d
// word2:    p   q 
// merged: a p b q c   d
 

// Constraints:

// 1 <= word1.length, word2.length <= 100
// word1 and word2 consist of lowercase English letters.

import (
	"fmt"
)

// func mergeAlternately(word1 string, word2 string) string {
// 	if word1 == "" && word2 != "" {
// 		return word2
// 	}
// 	if word2 == "" && word1 != "" {
// 		return word1
// 	}
// 	var result []byte
// 	var len1, len2 int
// 	for i, j := 0, 0; i < len(word1) && j < len(word2); {
// 		if (i+j)%2 == 0 {
// 			result = append(result, word1[i])
// 			i = i + 1
// 			len1 = i
// 		} else {
// 			result = append(result, word2[j])
// 			j = j + 1
// 			len2 = j
// 		}

// 	}
// 	if len1 < len(word1) {
// 		result = append(result, word1[len1:]...)
// 	}
// 	if len2 < len(word2) {
// 		result = append(result, word2[len2:]...)
// 	}
// 	return string(result)

	
// }

func mergeAlternately(word1 string, word2 string) string {
	if word1 == "" && word2 != "" {
		return word2
	}
	if word2 == "" && word1 != "" {
		return word1
	}
	var result []byte
	var len1, len2 int
	for i, j := 0, 0; i < len(word1) && j < len(word2); {
		if (i+j)%2 == 0 {
			result = append(result, word1[i])
			i = i + 1
			len1 = i
		} else {
			result = append(result, word2[j])
			j = j + 1
			len2 = j
		}

	}
	if len1 < len(word1) {
		result = append(result, word1[len1:]...)
	}
	if len2 < len(word2) {
		result = append(result, word2[len2:]...)
	}
	return string(result)
}
func main() {
	// word1, word2 := "abc", "pqr"
	// word1, word2 := "ab", "pqrs"
	word1, word2 := "abcd", "pq"
	fmt.Printf("Result : %s", mergeAlternately(word1, word2))
}
