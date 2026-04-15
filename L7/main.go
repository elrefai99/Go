package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func (p Person) setName(payload string) {
	p.name = payload
	fmt.Println(p)
}

func main() {

	str := Person{
		name: "Mohamed mostafa",
		age:  27,
	}
	str.setName("na")
	fmt.Println(str)
}
