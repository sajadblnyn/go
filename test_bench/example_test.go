package main

import (
	"testing"
)

func BenchmarkConcatStringsWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsWithBuilder("test1", "test2")
	}
}

func BenchmarkConcatStringsWithoutBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringsWithoutBuilder("test1", "test2")
	}
}

func BenchmarkAddItemsToSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddItemsToSlice(5, []int{})
	}
}

func BenchmarkAddItemsToMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddItemsToMap(5, make(map[int]int))
	}
}
