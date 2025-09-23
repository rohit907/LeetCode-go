package main

import (
	"fmt"
)

func reverseVowels(s string) string {

	tmpChar := []byte(s) // Direct conversion to byte slice
	vowels := make([]byte, 0, len(s))
	for _, ch := range tmpChar {
		switch ch {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			vowels = append(vowels, ch)
		}
	}
	index := len(vowels) - 1
	for i, ch := range tmpChar {
		switch ch {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			tmpChar[i] = vowels[index]
			index--
		}
	}
	return string(tmpChar)
}
func main() {
	s := "IceCreAm"
	s = "leetcode"
	fmt.Printf("Reverse Vowels of a String : %s", reverseVowels(s))
}
