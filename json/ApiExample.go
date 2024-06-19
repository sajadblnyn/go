package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	content := make(chan []byte)
	go getUrl("https://jsonplaceholder.typicode.com/todos", content)
	var posts []Post
	err := json.Unmarshal(<-content, &posts)

	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", posts)

	for _, v := range posts {
		fmt.Printf("%v\n", v.Title)

	}

}

func getUrl(url string, content chan<- []byte) {

	client := http.Client{}

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(httpReq)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content <- body
}
