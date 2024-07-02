package main

import "fmt"

func main() {
	type person struct {
		id   int
		name string
	}

	// var persons map[string]person = map[string]person{"dqwdqw": {id: 33666, name: "asdqwdwq"}}

	persons := map[string]person{"dqwdqw": {id: 33666, name: "asdqwdwq"}}

	// persons["1223553"] = person{id: 2335, name: "asdasd"}
	// persons["dfwewe"] = person{id: 3665, name: "ssaa"}
	// persons["5464964"] = person{id: 12545, name: "csaccs"}
	persons["6568889"] = person{id: 55489, name: "scswqwfw"}

	// p1, exists := persons["dqwdqw"]

	// fmt.Printf("%v\n", exists)

	// fmt.Printf("%v\n", p1)
	// delete(persons, "dqwdqw")

	fmt.Printf("%v\n", persons)

}
