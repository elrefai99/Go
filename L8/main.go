package main

import "fmt"

func main() {

	randumNumbers := []int{1, 5, 99, 55, 7, 900}
	for i, data := range randumNumbers {
		fmt.Println("index i: ", i, "value: ", data)
	}

	// randumNumbers = append(randumNumbers, 10000)
	sliceLastNumbers := randumNumbers[3:]
	sliceFirstNumbers := randumNumbers[:3]

	sliceFirstNumbers = append(sliceFirstNumbers, 9999)

	fmt.Println("slice First Numbers: ", sliceFirstNumbers, "length: ", len(sliceFirstNumbers), "capacity: ", cap(sliceFirstNumbers))
	fmt.Println("slice Last Numbers: ", sliceLastNumbers, "length: ", len(sliceLastNumbers), "capacity: ", cap(sliceLastNumbers))
}
