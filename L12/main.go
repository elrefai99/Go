package main

import (
	"fmt"
	"time"
)

func gorotine(message string, channel chan string) {
	const loop int = 3
	for i := 0; i < loop; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
	channel <- "Done"
	channel <- "Done Done"
}

func main() {
	channel := make(chan string, 1)
	go gorotine("Hi all, my name is Mohammed Mostafa", channel)

	getChannel := <-channel
	getChannel2 := <-channel
	fmt.Println(getChannel)
	fmt.Println(getChannel2)

}
