package main

import (
	"fmt"
	"strconv"
)

func minPartitions(n string) int {
	max := 0
	for _, v := range n {
		x, _ := strconv.Atoi(string(v))
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	fmt.Println(minPartitions("32"))
}