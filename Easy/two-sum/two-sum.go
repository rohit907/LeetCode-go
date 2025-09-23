package main

import "fmt"

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, value := range nums {
		if index, ok := myMap[target-nums[i]]; ok {
			return []int{index, i}
		}
		myMap[value] = i
	}
	return nil

}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Two Sum : %d", twoSum(nums, target))
}
