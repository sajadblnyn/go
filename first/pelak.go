package main

import "fmt"

func main() {
	fmt.Println("enter pelak number: ")
	var pelakNum float64
	fmt.Scanln(&pelakNum)

	switch {
	case pelakNum >= 1 && pelakNum <= 10:
		println("tehran")
		fallthrough
	case pelakNum >= 11 && pelakNum <= 20:
		println("yeza")
	default:
		println("no")
	}

}
