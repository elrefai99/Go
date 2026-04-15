package main

import (
	"Project1/sort"
	"fmt"
)

func main() {
	y := []int{1, 4, 55, 3, 66, 7, 442, 24, 546, 7, 864}

	BubbleSort := sort.BubbleSort(y)
	BubbleSort_2 := sort.BubbleSort_2(y)

	fmt.Println("BubbleSort", BubbleSort)
	fmt.Println("BubbleSort2", BubbleSort_2)
}
