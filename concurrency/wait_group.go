package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var toDoList []string
var mx sync.Mutex = sync.Mutex{}

func main() {

	t := time.Now().Unix()

	wg := sync.WaitGroup{}

	// wg.Add(10)
	for i := 1; i <= 50; i++ {
		// getToDo(i, &wg)
		wg.Add(1)
		go getToDo(i, &wg)
	}
	wg.Wait()
	t2 := time.Now().Unix()

	fmt.Printf("%v\n", toDoList)

	fmt.Printf("\n duration: %d \n", t2-t)
}

func getToDo(id int, wg *sync.WaitGroup) {
	getUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func getUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	mx.Lock()
	toDoList = append(toDoList, string(body))
	mx.Unlock()

	// when function has error and  you dont want to done wait group item
	// wg.Done()

}
