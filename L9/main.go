package main

import "fmt"

func setToNinety(x *int) {
	*x = 90
}

func setToEighty(x *int) {
	*x = 80
}

func main() {
	var x int
	fmt.Scanln(&x)

	fmt.Println("Before:", x)
	setToNinety(&x)
	fmt.Println("After Test:", x)

	setToEighty(&x)
	fmt.Println("After TwoPointer:", x)
}
