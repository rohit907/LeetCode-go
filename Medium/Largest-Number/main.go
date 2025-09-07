package main

// 179. Largest Number
// Solved
// Medium
// Topics
// premium lock iconCompanies

// Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

// Since the result may be very large, so you need to return a string instead of an integer.

// Example 1:

// Input: nums = [10,2]
// Output: "210"

// Example 2:

// Input: nums = [3,30,34,5,9]
// Output: "9534330"

// Constraints:

//     1 <= nums.length <= 100
//     0 <= nums[i] <= 109

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	tmp := make([]string, len(nums))
	for index, value := range nums {
		tmp[index] = strconv.Itoa(value)
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i]+tmp[j] > tmp[j]+tmp[i] //This ensures that 9 + 34 = "934" is greater than 34 + 9 = "349" — and sorts accordingly.
	})
	if tmp[0] == "0" {
		return "0"
	}
	return strings.Join(tmp, "")
}
func main() {
	nums := []int{3, 30, 34, 5, 9}

	fmt.Printf("largest number is : %s", largestNumber(nums))
	nums1 := []int{0, 0}

	fmt.Printf("largest number is : %s", largestNumber(nums1))
}
