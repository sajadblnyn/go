package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
	Age    int    `json:"age,omitempty"`
	// Age    int    `json:"-"`
}

func main() {
	prsJson := []byte(`[{"name":"sajad","family":"blnyn","age":20},{"name":"ali","family":"mamli","age":23},{"name":"mamad","family":"mamadi","age":30}]`)

	var prs []Person
	err := json.Unmarshal(prsJson, &prs)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", prs)

	pr1Json := []byte(`{"name":"sajad","family":"blnyn","age":20}`)

	var pr Person

	err = json.Unmarshal(pr1Json, &pr)

	if err != nil {
		panic(err)

	}
	fmt.Printf("%v\n", pr)

	fmt.Printf("%v\n", pr.Name)

}
