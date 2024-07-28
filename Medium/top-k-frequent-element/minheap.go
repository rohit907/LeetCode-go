package main

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

 

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]
 

// Constraints:

// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
 

// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	result := []int{}
	for _, v := range nums {
		if _, ok := freq[v]; !ok {
			freq[v] = 1
		} else {
			freq[v] = freq[v] + 1
		}
	}
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for k, v := range freq {
		pair := Pair{
			freq: v,
			key:  k,
		}
		heap.Push(minHeap, pair)
	}
	for k > 0 {
		element := heap.Pop(minHeap).(Pair)
		result = append(result, element.key)
		k--
	}
	//     temp := make([]int, len(freq))
	//     index:=0
	//     for _,v := range freq{
	//         temp[index] = v
	//     }
	//     sort.Ints(temp)
	//     result := make([]int,k)
	//     for k>0 {

	//         for key,v := range freq{
	//             if v== temp[len(temp)-1]{
	//                 result = append(result, key)
	//                 temp = temp[:len(temp)-1]
	//                 k--
	//             }
	//         }
	//     }
	return result

}

type Pair struct {
	key, freq int
}
type MinHeap []Pair

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)-1
	x := old[n]
	*h = old[:n]
	return x
}

func main() {
	nums := []int{3, 0, 1, 0}
	k := 1
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
