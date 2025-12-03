package main

import (
	"fmt"
	"sort"
)

func unique(slice1 []string, slice2 []string) []string {
	unique := make(map[string]struct{})
	result := []string{}
	for _, v := range slice1 {

		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
		}
	}
	for _, v := range slice2 {
		if _, ok := unique[v]; !ok {
			unique[v] = struct{}{}
		}
	}
	for k, _ := range unique {
		result = append(result, k)
	}
	sort.Strings(result)
	return result

}

func main() {
	s1 := []string{"t", "a", "v"}
	s2 := []string{"t", "h", "m", "t"}
	// [t a v h m]
	fmt.Println(unique(s1, s2))
}

