package main

import (
	"fmt"
	"sort"
	"strconv"
)
func transformDigit( myMap map[int]int) (int, error){
	tmp :=[]int{}
	for k,_ := range myMap{
		tmp = append(tmp, k)
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i]> tmp[j]
	})
	str :=""
	for _, v:= range tmp{
		digit := strconv.Itoa(v)
		str +=digit
	}
	result, err:=strconv.Atoi(str)
	fmt.Println("result",result)
	return result, err
}
func main() {
	x := 143
	myMap := make(map[int]int)
	newNum:= 0
	for x > 0 {
		digit := x % 10
		if _, ok := myMap[digit]; !ok {
			myMap[digit] = 1
		}
		x = x/10
	}
	fmt.Println("result d", myMap)
	newNum,_=transformDigit(myMap)
	fmt.Println("newNum", newNum)
	fmt.Println("helloe")
}