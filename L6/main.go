package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 6, 7, 3, 1}
	sl := arr[:4]
	fmt.Println(sl, len(sl), cap(sl))

	str := []string{"hi", "mohammed"}
	for i := 0; i < 10; i++ {
		fmt.Println(str)
	}
	for _, data := range str {
		fmt.Println(data)
	}
}
