package main

import (
	"fmt"
)

func main() {
	x := 4
	y := 10

	fmt.Println(x > y)  // false
	fmt.Println(x >= y) // false
	fmt.Println(x < y)  // true
	fmt.Println(x <= y) // true
	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x != y) // true

	// if condition
	if x < y {
		text := "Yes!!"
		fmt.Println(text) // false
	}
}
