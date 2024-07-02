package main

import (
	"PersianDate/DateConversions"
	"fmt"
)

type OrderStatus int
type Order struct {
	Id     int
	Status OrderStatus
}

const min_salary float64 = 5600000

const (
	Paid    OrderStatus = 1
	Waiting OrderStatus = 2
)

func init() {
	fmt.Println(DateConversions.ShamsiToMiladi("fqw"))
}
func main() {
	// var order1 Order = Order{Id: 1, Status: Paid}
	// jsonObj, _ := json.Marshal(order1)
	// fmt.Println(string(jsonObj))

	// var a int = 5
	// var b *int = &a
	// fmt.Println(*b)
	// printByRef(&a)
	// printByValue(a)

	// var a string = "سحاد"
	// fmt.Printf("%s,%T,%d", a, a, len(a))

	// var b []rune = []rune("سجاد")
	// b := []rune("سحاد")
	// fmt.Printf("%v,%T,%d", b, b, len(b))
	// fmt.Print(b)

	// fmt.Println(strings.Contains("sajad", "sa1"))

	// fmt.Println(strings.Count("sajad", "a"))
	// fmt.Println(strings.Cut("sajad", "a"))

	// fmt.Println(strings.Split("sajad", "a"))

	// fmt.Println(strings.Join(strings.Split("sajad", "a"), "."))

	// fmt.Println(strings.Repeat("sd", 5))

	// fmt.Println(strings.ReplaceAll("sajad", "a", "e"))

	// fmt.Println(strings.Compare("qa", "sa"))
	// fmt.Println(strings.EqualFold("sajad", "SaJad"))

	// fmt.Println(strings.HasPrefix("qa", "q"))

	// fmt.Println(strings.HasSuffix("qa", "q"))

	// fmt.Println(strings.Index("qa", "a"))

	// fmt.Println(strings.ToLower("SaJad"))

	// fmt.Println(strings.ToUpper("SaJad"))

	// fmt.Println(strings.Title("saJad is goood"))

	// fmt.Println(strings.Trim("__saJad is goood__", "_"))
	// fmt.Println(strings.TrimLeft("__saJad is goood__", "_"))
	// fmt.Println(strings.TrimRight("__saJad is goood__", "_"))

	// var tax float64 = 0
	// var salary float64
	// fmt.Print("enter your salary: ")
	// fmt.Scanln(&salary)

	// if salary <= min_salary {
	// 	tax = 0
	// } else if salary <= (2 * min_salary) {
	// 	tax = 0.10
	// } else if salary <= (4 * min_salary) {
	// 	tax = 0.20
	// } else {
	// 	tax = 0.40
	// }
	// fmt.Printf("your income:%f\n", salary-(salary*tax))

	i := 5

	var a *int
	a = &i

	println(*a)

}

func printByRef(a *int) {
	fmt.Println(a)
}
func printByValue(a int) {
	fmt.Println(&a)
}
