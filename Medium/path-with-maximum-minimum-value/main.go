package main

import (
	"container/heap"
	"fmt"
)
type pos struct{
    i,j int
}
type key struct{
    p pos
    min int
}

type MaxHeap []key

func(h MaxHeap)Len()int{ return len(h)}
func(h MaxHeap)Swap(i, j int){ h[i],h[j]= h[j],h[i]}
func(h MaxHeap)Less(i, j int)bool{ return h[i].min> h[j].min}

func(h *MaxHeap)Push(x interface{}){
    *h= append((*h), x.(key))
}

func(h *MaxHeap)Pop() interface{}{
    old := *h
    n := len(old)
    x:= old[n-1]
    *h = (*h)[:n-1]
    return x
}

func maximumMinimumPath(grid [][]int) int {
    h := &MaxHeap{ key{pos{0,0},grid[0][0]}}
    heap.Init(h)
     n := len(grid)
    m := len(grid[0])
    visited := make(map[pos]bool)
    helper:= func(i, j ,min int) {
        if i>=0 && j>=0 && i< n && j<m && !visited[pos{i,j}]{
            if grid[i][j]< min{
                min = grid[i][j]
            }
            heap.Push(h,key{pos{i,j},min})
            visited[pos{i,j}] = true
        }
    }
    for h.Len()!=0{
        pop := heap.Pop(h).(key)
        if pop.p.i == n-1 && pop.p.j == m-1{
            return pop.min
        }
        helper(pop.p.i-1,pop.p.j,pop.min)
        helper(pop.p.i+1,pop.p.j,pop.min)
        helper(pop.p.i,pop.p.j-1,pop.min)
        helper(pop.p.i,pop.p.j+1,pop.min)
    }
    return -1
}
func main(){
    grid := [][]int{{5,4,5},{1,2,6},{7,4,6}}
    value :=maximumMinimumPath(grid)
    fmt.Println(value)
}