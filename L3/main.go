package main

import (
	"fmt"
	"math"

	"strconv"
)

func main() {
	x := uint8(8)
	y := uint8(7)
	i := 10
	z := y + x + uint8(i)
	fmt.Println(z)

	// math package
	fmt.Println(math.Min(1, 4))
	fmt.Println(math.Max(1, 4))
	fmt.Println(math.Round(4.58))
	fmt.Println(math.Sqrt(4))
	fmt.Println(math.Pow(1, 4))

	// convert
	numberTest1 := "12345Hadi"
	numberTest2 := "12345"
	number1, err := strconv.Atoi(numberTest1) // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	number2, err := strconv.Atoi(numberTest2) // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	fmt.Println(number1, err)
	fmt.Println(number2, err)
}
