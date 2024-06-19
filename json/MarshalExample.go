package main

import "encoding/json"

type Person struct {
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
	Age    int    `json:"age,omitempty"`
	// Age    int    `json:"-"`
}

func main() {
	pr1 := Person{Name: "sajad", Family: "blnyn", Age: 20}
	pr2 := Person{Name: "ali", Family: "mamli", Age: 23}
	pr3 := Person{Name: "mamad", Family: "mamadi", Age: 30}

	prs := []Person{pr1, pr2, pr3}
	pr1Json, err := json.Marshal(pr1)

	if err != nil {
		panic(err)
	}
	println(string(pr1Json))

	prsJson, err := json.Marshal(prs)

	if err != nil {
		panic(err)

	}
	println(string(prsJson))

}
