package main

import (
	"fmt"
	"os"
)

func main() {

	var count_numbers int
	var count_queries int

	// fmt.Println("enter count numbers: ")
	fmt.Scan(&count_numbers)

	// fmt.Println("enter count queries: ")
	fmt.Scan(&count_queries)

	if count_numbers < 1 {
		println("count_numbers should be greater than one")
		os.Exit(0)
	}

	if count_queries > 1000 || count_queries < 1 {
		println("count_numbers should be less than 1000 and greater than 1")
		os.Exit(0)
	}

	var numbers []int = []int{}

	var input int
	for i := 0; i < count_numbers; i++ {
		// fmt.Printf("enter number %d :", i)
		fmt.Scan(&input)
		numbers = append(numbers, input)
		if numbers[i] > 1000000 || numbers[i] < 1 {
			println("number should be between 1 to 1000000")
			os.Exit(0)
		}
	}

	var query_numbers []int

	for i := 0; i < count_queries; i++ {
		// fmt.Printf("enter query number %d :", i)
		fmt.Scan(&input)
		query_numbers = append(query_numbers, input)

		if query_numbers[i] > 1000000 || query_numbers[i] < 1 {
			println("query number should be between 1 to 1000000")
			os.Exit(0)
		}
	}

	for _, v := range query_numbers {
		fmt.Println(findSmallerNumbersCountOfNumber(&v, &numbers))
	}

}

func findSmallerNumbersCountOfNumber(number *int, numbers *[]int) int {
	var smaller_numbers_count int = 0
	for _, v := range *numbers {
		if v < *number {
			smaller_numbers_count = smaller_numbers_count + 1
		}
	}
	return smaller_numbers_count
}

func twoSum(nums []int, target int) []int {
	var diff int
	for i, v := range nums {
		diff = target - v
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == diff {
				return []int{i, j}
			}
		}

	}
	return []int{0, 0}
}

func twoSumByMap(nums map[int]int, target int) []int {
	var diff int
	var exists bool
	var value int

	for i, v := range nums {
		diff = target - i
		value, exists = nums[diff]
		if exists && value != v {
			return []int{value, v}
		}
	}
	return []int{0, 0}
}

func twoSumNew(numbers []int, target int) []int {
	nums := map[int]int{}
	var diff int
	var exists bool
	var value int

	for i, v := range numbers {
		diff = target - v
		value, exists = nums[diff]
		if exists {
			return []int{value, i}
		}
		nums[v] = i
	}
	return []int{0, 0}
}
