package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	MaxHeap := &MaxHeap{}
	for _, v := range stones {
		heap.Push(MaxHeap, v)
	}

	for MaxHeap.Len() > 1 {
		x := heap.Pop(MaxHeap).(int)
		y := heap.Pop(MaxHeap).(int)
		if x != y {
			heap.Push(MaxHeap, x-y)
		}
	}
	res := 0
	if MaxHeap.Len() == 1 {
		res = (*MaxHeap)[0]
	}
	return res
}

func main() {
	stones := []int{12,7,4,1,8,1}
	fmt.Println(lastStoneWeight(stones))
}
