package main

import (
	"Project1/sort"
	"flag"
	"fmt"
)

func main() {
	y := []int{1, 4, 55, 3, 66, 7, 442, 24, 546, 7, 864}

	BubbleSort := flag.Bool("bubble1", false, "run bubble sort 1")
	BubbleSort_2 := flag.Bool("bubble2", false, "run bubble sort 2")

	flag.Parse()

	if *BubbleSort {
		sort.BubbleSort(y)
		fmt.Println("BubbleSort1:", y)
	}

	if *BubbleSort_2 {
		sort.BubbleSort_2(y)
		fmt.Println("BubbleSort2:", y)
	}
}
