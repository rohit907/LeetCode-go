package main

import "fmt"

func squareIsWhite(coordinates string) bool {

	return ((coordinates[0]-'a')%2 == 0 || coordinates[1]-'1'%2 != 0) && ((coordinates[0]-'a')%2 != 0 || coordinates[1]-'1'%2 == 0)
}

func main() {
	fmt.Println(squareIsWhite("a1"))
}