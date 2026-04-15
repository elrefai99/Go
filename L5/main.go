package main

import (
	"fmt"
)

func main() {
	str := "hi All"

	fmt.Println(str[0])
	for _, char := range str { // _ is index array
		fmt.Println(char)
	}
}
