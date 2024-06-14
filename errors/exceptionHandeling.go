package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	defer mainDefer()
	list := []int{1, 2, 3, 4}

	//this remove three
	// removeFromListByRef(&list, 1)

	// removeFromListByRef(&list, 5)

	removeFromListByRef(&list, 3)

	fmt.Printf("%v\n", list)

	panic("manual panic")
}

func getFromList(l *[]int, i int) (res int) {
	res = -1
	defer func() {
		fmt.Printf("get from list defer  \n")

		err := recover()
		if err != nil {
			fmt.Printf("error message: %s\n", err)
			debug.PrintStack()

		}
	}()
	return (*l)[i]
}

func removeFromListByRef(l *[]int, index int) {

	defer func() {
		fmt.Printf("removeFromListByRef defer  \n")

		err := recover()
		if err != nil {
			fmt.Printf("error message: %s\n", err)
			// s := debug.Stack()

			debug.PrintStack()
		}
	}()

	item := getFromList(l, index)

	*l = append((*l)[:item], (*l)[item+1:]...)
}

func mainDefer() {
	fmt.Printf("main defer  \n")

	err := recover()
	if err != nil {
		fmt.Printf("error message: %s\n", err)

	}
}
