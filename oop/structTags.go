package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	Name   string `json:"first_name"`
// 	Family string `json:"last_name"`
// 	Age    int    `json:"age"`
// }

// type Person struct {
// 	Name   string `json:"first_name"`
// 	Family string `json:"last_name"`
// 	Age    int    `json:"-"`
// }

type Person struct {
	Name   string `json:"first_name"`
	Family string `json:"last_name"`
	Age    int    `json:"age,omitempty"`
}

func main() {
	var person Person = Person{Name: "sajad", Family: "balanian", Age: 25}

	// jsonPerson, _ := (json.Marshal(person))
	// fmt.Printf("%v\n", string(jsonPerson))

	jsonPerson, _ := (json.MarshalIndent(person, "", " "))
	fmt.Printf("%v\n", string(jsonPerson))

}
