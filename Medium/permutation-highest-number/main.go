package main

import (
	"fmt"
	"sort"
	"strconv"
)

func getNextHighestNumber(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	fmt.Println(digits)
	i := 0
	combo := make(map[string]int)
	for i < len(digits) {
		j := 0
		num := strconv.Itoa(digits[i])
		for j < len(digits) {
			if i != j {
				num += strconv.Itoa(digits[j])
			}
			fmt.Println(j, i)
			j++
			
		}
		fmt.Println(num)
		if _,ok:=combo[num];!ok{
			combo[num]=1
		}
		i++
	}
	list :=[]int{}
	for k, _:= range combo{
		number,_ := strconv.Atoi(k)
		list = append(list, number)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	fmt.Println(list)
	return list[len(list)-1]
}

func main(){
	fmt.Println(getNextHighestNumber(134))
}