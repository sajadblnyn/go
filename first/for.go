package main

type person struct {
	name string
	id   int
}

func main() {

	// var arr []int = []int{23, 12, 10, 5}

	// for i := 0; i < 4; i++ {
	// 	println(arr[i])
	// }

	// for _, v := range arr {
	// 	println(v)
	// }

	var persons []person = []person{{id: 10, name: "sajad"}, {id: 23, name: "er"}}

	for _, v := range persons {
		println(v.name)
		println(v.id)
	}
}
