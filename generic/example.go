package main

import "fmt"

type Number interface {
	int | float32
}

func main() {
	nums1 := []int{5, 2, 3, 6}
	nums2 := []float32{5.5, 2.6, 3.7}

	fmt.Println(sumSlice(nums1))
	fmt.Println(sumSlice(nums2))

}

// func sumSlice[T int | float32](nums []T) (sum T) {
// 	for _, v := range nums {
// 		sum = v + sum
// 	}
// 	return sum
// }

func sumSlice[T Number](nums []T) (sum T) {
	for _, v := range nums {
		sum = v + sum
	}
	return sum
}
