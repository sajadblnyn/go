package main

import "fmt"

type Pr struct {
	int
	string
	float32
}

func main() {
	// var person struct {
	// 	id   int
	// 	name string
	// } = struct {
	// 	id   int
	// 	name string
	// }{id: 5, name: "sajad"}
	// person := struct {
	// 	id   int
	// 	name string
	// }{
	// 	id:   5,
	// 	name: "sajad",
	// }

	// fmt.Printf("%v", person)

	pr := Pr{1, "sajad", 25}
	fmt.Printf("%v\n", pr)

}
