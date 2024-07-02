package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go task1()

	go task2()

	go task3()

	count := 0

	go func() {
		count++
	}()
	go func() {
		count++
	}()
	go func() {
		count++
	}()

	time.Sleep(time.Second)
	fmt.Println(count)

}

func task1() {
	fmt.Println("task1")
}
func task2() {
	fmt.Println("task2")
}
func task3() {
	fmt.Println("task3")
}
