package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	char := []byte(s)
	stack := []int{}
	for i := 0; i < len(char); i++ {
		if char[i] == '(' {
			stack = append(stack, i)
		}
		if char[i] == ')' {
			if len(stack) == 0 {
				char[i] = 0
			} else {
				stack = stack[:len(stack)-1]
			}

		}
	}
	var result []byte
	for _, idx := range stack {
		char[idx] = 0
	}
	for idx, _ := range char {
		if char[idx] != 0 {
			result = append(result, char[idx])
		}
	}
	return string(result)
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)")) // "lee(t(c)o)de"
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))       // "ab(c)d"
	fmt.Println(minRemoveToMakeValid("))(("))

}
