package main

import "strings"

func ConcatStringsWithBuilder(a string, b string) (s string) {
	var sb strings.Builder
	for i := 0; i < 1000; i++ {
		sb.WriteString(a + b)
	}
	return sb.String()
}

func ConcatStringsWithoutBuilder(a string, b string) (s string) {
	for i := 0; i < 1000; i++ {
		s = a + b
	}
	return s
}

func AddItemsToSlice(a int, s []int) []int {
	for i := 0; i < 1000; i++ {
		s = append(s, a)
	}
	return s
}
func AddItemsToMap(a int, m map[int]int) map[int]int {
	for i := 0; i < 1000; i++ {
		m[i] = a
	}
	return m
}
