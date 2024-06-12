package main

import "fmt"

func main() {
	// var n int
	// fmt.Scan(&n)
	// var nums []int
	// var input int
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&input)
	// 	nums = append(nums, input)
	// }
	// var sums []int

	// sums = nums

	// var sum int

	// for i := 0; i < n-1; i++ {
	// 	sum = nums[i] + nums[i+1]
	// 	sums = append(sums, sum)

	// 	for j := i + 1; j < n-1; j++ {

	// 		sum = sum + nums[j+1]
	// 		sums = append(sums, sum)

	// 	}
	// }

	// var max int = -100000

	// for _, v := range sums {
	// 	if v > max {
	// 		max = v
	// 	}
	// }

	// fmt.Print(max)

	// var n int
	// fmt.Scan(&n)

	// nums := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&nums[i])
	// }

	// maxCurrent := nums[0]
	// maxGlobal := nums[0]

	// for i := 1; i < n; i++ {
	// 	if nums[i] > maxCurrent+nums[i] {
	// 		maxCurrent = nums[i]
	// 	} else {
	// 		maxCurrent += nums[i]
	// 	}

	// 	if maxCurrent > maxGlobal {
	// 		maxGlobal = maxCurrent
	// 	}
	// }

	// fmt.Println(maxGlobal)

	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var max_current int = nums[0]

	var max_global int = nums[0]

	for i := 0; i < n-1; i++ {
		if nums[i+1] < max_current+nums[i+1] {
			max_current = max_current + nums[i+1]
		} else {
			max_current = nums[i+1]
		}

		if max_current > max_global {
			max_global = max_current
		}
	}
	fmt.Print(max_global)

}
