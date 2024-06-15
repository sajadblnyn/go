package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	var div int = n / 2

	var a, b, c int

	arr := map[int][]int{}

	for i := 1; i <= div; i++ {
		a = i
		b = i
		c = n - (a + b)

		if (a+b) > c && (c+b) > a && (a+c) > b {
			arr[a*b*c] = []int{a, b, c}
		}

		a = i
		b = i + 1
		c = n - (a + b)

		if (a+b) > c && (c+b) > a && (a+c) > b {
			arr[a*b*c] = []int{a, b, c}
		}
	}

	fmt.Println(len(arr))
}
