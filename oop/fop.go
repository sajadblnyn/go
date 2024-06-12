package main

import "fmt"

type Person struct {
	id   int
	name string
}
type PersonOptions func(*Person)

func main() {

	person := NewPerson(SetId(1), SetName("sajad"))
	fmt.Printf("%v", person)

}
func NewPerson(options ...PersonOptions) *Person {
	person := &Person{}

	for _, option := range options {
		option(person)
	}
	return person
}
func SetId(id int) PersonOptions {
	return func(person *Person) {
		person.id = id
	}
}

func SetName(name string) PersonOptions {
	return func(person *Person) {
		person.name = name
	}
}
