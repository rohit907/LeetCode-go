package main
// 1418. Display Table of Food Orders in a Restaurant
// Solved
// Medium
// Topics
// Companies
// Hint
// Given the array orders, which represents the orders that customers have done in a restaurant. More specifically orders[i]=[customerNamei,tableNumberi,foodItemi] where customerNamei is the name of the customer, tableNumberi is the table customer sit at, and foodItemi is the item customer orders.

// Return the restaurant's “display table”. The “display table” is a table whose row entries denote how many of each food item each table ordered. The first column is the table number and the remaining columns correspond to each food item in alphabetical order. The first row should be a header whose first column is “Table”, followed by the names of the food items. Note that the customer names are not part of the table. Additionally, the rows should be sorted in numerically increasing order.

// Example 1:

// Input: orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
// Output: [["Table","Beef Burrito","Ceviche","Fried Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
// Explanation:
// The displaying table looks like:
// Table,Beef Burrito,Ceviche,Fried Chicken,Water
// 3    ,0           ,2      ,1            ,0
// 5    ,0           ,1      ,0            ,1
// 10   ,1           ,0      ,0            ,0
// For the table 3: David orders "Ceviche" and "Fried Chicken", and Rous orders "Ceviche".
// For the table 5: Carla orders "Water" and "Ceviche".
// For the table 10: Corina orders "Beef Burrito".
// Example 2:

// Input: orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]]
// Output: [["Table","Canadian Waffles","Fried Chicken"],["1","2","0"],["12","0","3"]]
// Explanation:
// For the table 1: Adam and Brianna order "Canadian Waffles".
// For the table 12: James, Ratesh and Amadeus order "Fried Chicken".
// Example 3:

// Input: orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melissa","2","Soda"]]
// Output: [["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]



import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	// Initialize the header and other necessary variables
	header := []string{}
	header = append(header, "Table")           // Add "Table" as the first header entry
	tables := []int{}                          // Slice to store unique table numbers
	foodIndexMapping := make(map[int][]string) // Map to associate table numbers with food items
	foodIndex := make(map[string]int)          // Map to track food item indices
	foodList := []string{}                     // Slice to store unique food items

	// Iterate through each order
	for _, order := range orders {
		// Convert the table number from string to int
		tableNumber, _ := strconv.Atoi(order[1])

		// Check if the table number already exists in the mapping
		if _, ok := foodIndexMapping[tableNumber]; !ok {
			// If not, create a new slice for the table and append it to the tables slice
			foodIndexMapping[tableNumber] = []string{}
			tables = append(tables, tableNumber)
		}

		// Check if the food item is already indexed
		if _, ok := foodIndex[order[2]]; !ok {
			// If not, add it to the food index and list
			foodIndex[order[2]] = 1
			foodList = append(foodList, order[2])
		}

		// Append the food item to the table's food list
		slice := foodIndexMapping[tableNumber]
		slice = append(slice, order[2])
		foodIndexMapping[tableNumber] = slice
	}

	// Sort the tables and food list
	sort.Ints(tables)
	sort.Strings(foodList)

	// Create an index for food items
	for k, v := range foodList {
		foodIndex[v] = k
	}

	// Append food items to the header
	header = append(header, foodList...)

	// Initialize the result with the header
	result := [][]string{}
	result = append(result, header)

	// Process each table to create the final result
	for i := 0; i < len(tables); i++ {
		// Create a temporary slice to hold counts of food items
		tmp := make([]int, len(foodList)+1)
		tmp[0] = tables[i] // Set the table number as the first element

		// Get the food items for the current table
		foodOntables := foodIndexMapping[tables[i]]

		// Count occurrences of each food item for the current table
		for _, v := range foodOntables {
			tmpIndex := foodIndex[v] + 1 // Get the index of the food item
			tmp[tmpIndex] += 1           // Increment the count
		}

		// Convert the counts to string format for the result
		str := []string{}
		for _, v := range tmp {
			str = append(str, strconv.Itoa(v))
		}

		// Append the counts for the current table to the result
		result = append(result, str)
	}

	return result // Return the final result
}
func main() {
	orders := [][]string{
		{"Order1", "1", "Pizza"},
		{"Order2", "1", "Burger"},
		{"Order3", "2", "Pizza"},
		{"Order4", "2", "Pasta"},
		{"Order5", "1", "Burger"},
	}

	result := displayTable(orders)

	// Print the result
	for _, row := range result {
		fmt.Println(row)
	}
}