package main

import "fmt"

type Person struct {
	name   string
	family string
	age    int
}

type PersonBuilder interface {
	SetName(name string) *Person
	SetFamily(family string) *Person
	SetAge(age int) *Person
	GetName() string
	GetFamily() string
	GetAge() int
}

func main() {
	var pr PersonBuilder = &Person{}
	pr.SetName("sajad").SetFamily("balanian").SetAge(23)

	fmt.Println(pr)
}

func (person *Person) SetName(name string) *Person {
	person.name = name
	return person

}
func (person *Person) SetFamily(family string) *Person {
	person.family = family
	return person

}
func (person *Person) SetAge(age int) *Person {
	person.age = age
	return person

}
func (person *Person) GetName() string {
	return person.name
}
func (person *Person) GetFamily() string {
	return person.family
}
func (person *Person) GetAge() int {
	return person.age

}
