package main

import (
	"fmt"
)

// func canPlaceFlowers(flowerbed []int, n int) bool {

//		for i := 0; i < len(flowerbed); i++ {
//			if flowerbed[i] == 0 {
//				emptyPrev := (i == 0) || (flowerbed[i-1] == 0)
//				emptyNext := (i == len(flowerbed)-1) || (flowerbed[i+1] == 0)
//				if emptyPrev && emptyNext {
//					flowerbed[i] = 1
//					n = n - 1
//				}
//			}
//		}
//		return n == 0
//	}
func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}
		if i-1 >= 0 && i+1 < len(flowerbed) {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else {
			if i == len(flowerbed)-1 && i-1 >= 0 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			} else if i == 0 && i+1 < len(flowerbed) && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}

	}

	return n == 0
}
func main() {
	flowerbed := []int{0, 0, 1, 0, 1}
	n := 2
	fmt.Printf("Can Place Flower: %b", canPlaceFlowers(flowerbed, n))
}
