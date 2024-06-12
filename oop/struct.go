package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	pr := Person{Id: 1, Name: "sajad", Age: 5}

	pr2 := new(Person)
	pr2.Id = 2
	pr2.Name = "ali"
	pr2.Age = 12

	var pr3 *Person = &Person{Id: 5, Name: "mamad", Age: 22}

	pr4 := NewPerson(5, "mehran", 25)

	pr5 := NewPersonOption(Person{Id: 23, Name: "sahar", Age: 26})

	fmt.Println(pr.CalBirthYear())
	fmt.Println(pr2.CalBirthYear())
	fmt.Println(pr3.CalBirthYear())
	fmt.Println(pr4.CalBirthYear())
	fmt.Println(pr5.CalBirthYear())

}

func NewPerson(id int, name string, age int) *Person {
	return &Person{Id: id, Name: name, Age: age}
}

func NewPersonOption(person Person) *Person {
	return &person
}

func (person *Person) CalBirthYear() int {
	return 2023 - person.Age
}
