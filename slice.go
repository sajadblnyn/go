package main

import "fmt"

func main() {
	s := []int{5, 10, 20, 30, 40}

	// var d []int = s[:]

	// var x []int = d

	// d = append(d, 30)
	// s[4] = 50

	// fmt.Printf("%v\n", x)
	// fmt.Printf("%v\n", d)
	// fmt.Printf("%v\n", s)

	// for i, _ := range s {
	// 	s[i] = s[i] + 1
	// }

	// fmt.Printf("%v", s)

	// addItem(50, &s)
	// fmt.Printf("%v", s)

	// ns := []int{1, 1, 1, 1, 1}

	// cou:=copy(ns, s)

	// fmt.Printf("%v", ns)
	// fmt.Printf("%d", co)

	// s = append(s, ns...)
	// fmt.Printf("%v", s)

	// s = append(s, 0)

	// copy(s[3:], s[2:])

	// s[2] = 50
	// fmt.Printf("%v", s)

	// s = append(s[:2], s[3:]...)
	// fmt.Printf("%v", s)

	// s = s[:0]
	// fmt.Printf("%v", s)

	s = nil
	fmt.Printf("%v", s)

}

func addItem(i int, a *[]int) {
	*a = append(*a, i)
}
