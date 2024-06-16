package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	count := 0
	for a := 1; a <= n/3; a++ {
		minX := max(0, (n/2)-2*a+1)
		maxX := (n - 3*a) / 2

		fmt.Printf("minx:%d , maxx:%d\n", minX, maxX)
		if minX <= maxX {
			count += maxX - minX + 1
		}
	}

	fmt.Println(count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
