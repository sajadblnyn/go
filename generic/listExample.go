package main

import "fmt"

type List[T any] struct {
	Items []T
}

// type List[T interface{}] struct {
// 	Items []T
// }

type ListTypes interface {
	int | float32 | string
}

func main() {
	var li *List[int] = &List[int]{Items: []int{1, 2, 3}}
	li.InsertAt(1, 15)

	fmt.Printf("%v", li.Items)

	InsertAtList(&li.Items, 2, 7)
	fmt.Printf("%v", li.Items)

	var li2 []string = []string{"test", "sj"}
	InsertAtList(&li2, 1, "bl")
	fmt.Printf("%v", li2)

}

func (list *List[T]) InsertAt(index int, value T) {

	list.Items = append(list.Items, value)

	copy(list.Items[index+1:], list.Items[index:])

	list.Items[index] = value

}

// func InsertAtList[T int | float32 | string](list *[]T, index int, value T) {

// 	*list = append(*list, value)

// 	copy((*list)[index+1:], (*list)[index:])

// 	(*list)[index] = value

// }

func InsertAtList[T ListTypes](list *[]T, index int, value T) {

	*list = append(*list, value)

	copy((*list)[index+1:], (*list)[index:])

	(*list)[index] = value

}
