package main

import (
	"fmt"
	"io"
	"os"
)

type Pr struct {
	id   int
	name string
}

var S []Pr = []Pr{}

func main() {

	// sum, ml := calculator(5, 6, 1, 3, 5)

	// fmt.Printf("sum:%d , mul:%d", sum, ml)

	// PrintLog(func(i int) bool {
	// 	if i >= 2 {
	// 		return true
	// 	}
	// 	return false
	// }, 5, 6, "fqwf", map[string]int{"me": 100}, []int{5, 3, 4})

	// nums := []int{1, 5, 22, 3, 2}

	// sort.Slice(nums, func(i, j int) bool {
	// 	if nums[i] < nums[j] {
	// 		return true
	// 	}
	// 	return false
	// })
	// fmt.Printf("%v", nums)

	// names := []string{"sajad", "ali", "reza", "mamal", "asghar"}

	// for i, v := range names {
	// 	func(a int, s string) {
	// 		go println(a, s+"*")
	// 	}(i, v)
	// }
	// time.Sleep(time.Second * 5)\

	// CopyFile("test2.txt", "test.txt")

	// defer println("test1")
	// defer println("test2")

	S = []Pr{{id: 1, name: "sajad"}, {id: 5, name: "mamal"}, {id: 2, name: "ali"}, {id: 3, name: "reza"}, {id: 4, name: "hamed"}}
	fmt.Printf("%v", findPerson(3))
}

func findPerson(id int) *Pr {
	for _, v := range S {
		if v.id == id {
			return &v
		}
	}
	return nil
}

func calculator(nums ...int) (sum int, mult int) {
	mult = 1
	for _, v := range nums {
		sum += v
		mult *= v
	}
	return sum, mult
}

func PrintLog(greather func(i int) bool, logs ...interface{}) {
	for i, v := range logs {
		if greather(i) {
			fmt.Printf("--- %v ---\n", v)

		}
	}
}

func CopyFile(dest string, source string) (int64, error) {
	s, err := os.Open(source)
	if err != nil {
		return -1, err
	}
	defer s.Close()

	d, err := os.Create(dest)

	if err != nil {
		return -1, err
	}
	defer d.Close()

	return io.Copy(d, s)
}
