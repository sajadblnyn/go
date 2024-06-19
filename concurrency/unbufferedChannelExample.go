package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup = sync.WaitGroup{}

func main() {

	content := make(chan string)

	t1 := time.Now().Unix()

	for i := 0; i < 100; i++ {

		go getUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(i+1), content)
		go getContent(content)
	}
	wg.Wait()

	// for i := 0; i < 10; i++ {
	// 	go getUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(i+1), content)
	// }
	// var response string
	// for i := 0; i < 10; i++ {
	// 	response = <-content
	// 	fmt.Printf("%s\n", response)
	// }

	t2 := time.Now().Unix()

	fmt.Printf("duration: %d\n", t2-t1)

}

func getUrl(url string, content chan<- string) {
	wg.Add(1)

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

	content <- string(body)
}

func getContent(content <-chan string) {
	wg.Add(1)

	defer wg.Done()

	response := <-content
	fmt.Printf("%s\n", response)
}
