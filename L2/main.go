package main

import "fmt"

func main() {
	const x string = "Hi Mohammed"
	const y uint8 = 255
	z := 5
	ton_number := 9.9652424523
	fmt.Println(x, y, z)
	fmt.Printf("type of x is %T", x)  // type
	fmt.Printf("\n %v", y)            // value
	fmt.Printf("\n %e", ton_number)   // ton of different decimal number
	fmt.Printf("\n %b", z)            // binary
	fmt.Printf("\n %.2f", ton_number) // get 2 number of float
}
