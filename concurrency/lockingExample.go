package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	Id     int
	Salary int64
}

var Employees []Employee

func main() {
	var money int64 = 51000000

	wg := sync.WaitGroup{}

	mx := sync.Mutex{}
	// mx := sync.RWMutex{} //when you want to just lock read and write be unlocked
	for i := 1; i <= 5000; i++ {
		Employees = append(Employees, Employee{Id: i, Salary: 10000})
	}

	wg.Add(5000)
	for _, v := range Employees {
		go func() {
			defer wg.Done()

			// atomic.AddInt64(&money, -v.Salary) //use cpu feature and have better performance for math calculations use cases

			mx.Lock()
			money -= (v.Salary)
			mx.Unlock()

		}()
	}
	wg.Wait()

	fmt.Printf("remaining money: %d\n", money)
}
