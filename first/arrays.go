package main

import "fmt"

func main() {
	var names [5]string = [5]string{"sajad", "niam"}
	// addToNames(&names, "wefw")
	// fmt.Printf("%v", names)
	// var namesP *[5]string

	// namesP = &names

	// fmt.Print(namesP)
	var name string = "sajad"

	addToNames(&names, &name)
	fmt.Printf("%v", names)
}

func addToNames(xx *[5]string, name *string) {
	fmt.Printf("%v", *xx)
	// xx[0] = "ali"
	// *name = "sjsj"
}
