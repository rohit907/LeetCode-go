package main

// 901. Online Stock Span
// Solved
// Medium
// Topics
// premium lock icon
// Companies
// Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current day.

// The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going backward) for which the stock price was less than or equal to the price of that day.

// For example, if the prices of the stock in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from today, the price of the stock was less than or equal 2 for 4 consecutive days.
// Also, if the prices of the stock in the last four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, the price of the stock was less than or equal 8 for 3 consecutive days.
// Implement the StockSpanner class:

// StockSpanner() Initializes the object of the class.
// int next(int price) Returns the span of the stock's price given that today's price is price.
 

// Example 1:

// Input
// ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
// [[], [100], [80], [60], [70], [60], [75], [85]]
// Output
// [null, 1, 1, 1, 2, 1, 4, 6]

// Explanation
// StockSpanner stockSpanner = new StockSpanner();
// stockSpanner.next(100); // return 1
// stockSpanner.next(80);  // return 1
// stockSpanner.next(60);  // return 1
// stockSpanner.next(70);  // return 2
// stockSpanner.next(60);  // return 1
// stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
// stockSpanner.next(85);  // return 6
 

// Constraints:

// 1 <= price <= 105
// At most 104 calls will be made to next.

import "fmt"

// type StockSpanner struct {
// 	stock []int
// }

// func Constructor() StockSpanner {
// 	return StockSpanner{
// 		stock: []int{},
// 	}
// }

// func (s *StockSpanner) Next(price int) int {
// 	if len(s.stock) == 0 {
// 		s.stock = append(s.stock, 0)
// 	}
// 	span := 1
// 	for i := len(s.stock) - 1; i > 0; i-- {
// 		if s.stock[i] <= price {
// 			span++
// 		} else {
// 			break
// 		}
// 	}
// 	s.stock = append(s.stock, price)
// 	return span
// }

type StockSpanner struct {
	stock [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stock: make([][]int, 0),
	}
}

func (s *StockSpanner) Next(price int) int {
	span := 1
	for len(s.stock) > 0 && s.stock[len(s.stock)-1][0] <= price {
		span += s.stock[len(s.stock)-1][1]
		s.stock = s.stock[:len(s.stock)-1]
	}
	s.stock = append(s.stock, []int{price, span})
	return span

}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
