package main

import "fmt"

var maxd int = 100000

func main() {

	var n int
	var q int

	fmt.Scan(&n)
	fmt.Scan(&q)

	numbers := []int{}

	var input int
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		numbers = append(numbers, input)
	}

	questions := []int{}

	for i := 0; i < q; i++ {
		fmt.Scan(&input)
		questions = append(questions, input)
	}

	dp := [100000]int{}

	for _, v := range numbers {
		dp[v] += 1
	}

	for i := (maxd - 1); i > 0; i-- {
		dp[i-1] += dp[i]
	}

	for _, v := range questions {
		fmt.Println(n - dp[v])
	}

}
