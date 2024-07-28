package main

import (
	"container/heap"
	"fmt"
)

// Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

// Implement KthLargest class:

// KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
// int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

// Example 1:

// Input
// ["KthLargest", "add", "add", "add", "add", "add"]
// [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
// Output
// [null, 4, 5, 5, 8, 8]

// Explanation
// KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
// kthLargest.add(3);   // return 4
// kthLargest.add(5);   // return 5
// kthLargest.add(10);  // return 5
// kthLargest.add(9);   // return 8
// kthLargest.add(4);   // return 8

// Constraints:

// 1 <= k <= 104
// 0 <= nums.length <= 104
// -104 <= nums[i] <= 104
// -104 <= val <= 104
// At most 104 calls will be made to add.
// It is guaranteed that there will be at least k elements in the array when you search for the kth element.

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

type KthLargest struct {
	minHeap MinHeap
	kth     int
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeap{}
	heap.Init(&h)
	obj := KthLargest{
		minHeap: h,
		kth:     k,
	}
	for _, v := range nums {
		obj.Add(v)
	}
	return obj
}
func (this *KthLargest) Add(val int) int {
	heap.Push(&this.minHeap, val)
	if this.minHeap.Len() > this.kth {
		heap.Pop(&this.minHeap)
	}
	return (this.minHeap)[0]
}
func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3))  // return 4
	fmt.Println(kthLargest.Add(5))  // return 5
	fmt.Println(kthLargest.Add(10)) // return 5
	fmt.Println(kthLargest.Add(9))  // return 8
	fmt.Println(kthLargest.Add(4))  // return 8

}
