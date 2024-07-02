package main

import (
	"fmt"
	"time"
)

func main() {

	numChannel := make(chan int)

	go SendDataToChannel(numChannel)

	go ReceiveDataToChannel(numChannel)

	time.Sleep(time.Second)

}

func SendDataToChannel(numChannel chan int) {

	numChannel <- 5

}

func ReceiveDataToChannel(numChannel chan int) {

	num := <-numChannel

	fmt.Println(num)

}
