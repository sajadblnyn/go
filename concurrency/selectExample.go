package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	resource1 := make(chan string)

	resource2 := make(chan string)

	go getUrl("https://api.digikala.com/v1/categories/mobile-phone/search/?page=1", resource1)
	go getUrl("https://api.torob.com/v4/base-product/search/?page=0&sort=popularity&size=24&category=94&category_name=%DA%AF%D9%88%D8%B4%DB%8C-%D9%85%D9%88%D8%A8%D8%A7%DB%8C%D9%84-mobile&_url_referrer=https%3A%2F%2Fwww.google.com%2F&source=next_desktop", resource2)

	// select {
	// case result := <-resource1:
	// 	println(result)
	// case result := <-resource2:
	// 	println(result)
	// case <-time.After(time.Second * 3):
	// 	println("Time out")

	// }

	for true {
		select {
		case result := <-resource1:
			println(result)
			return
		case result := <-resource2:
			println(result)
			return
		case <-time.After(time.Second * 3):
			println("Time out")
			return
		default:
			println("default")

		}
	}
}

func getUrl(url string, content chan<- string) {

	// if url == "https://api.digikala.com/v1/categories/mobile-phone/search/?page=1" {
	// 	time.Sleep(time.Second * 3)
	// }

	// time.Sleep(time.Second * 3)

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

	dest := &bytes.Buffer{}

	err = json.Indent(dest, body, "", " ")
	if err != nil {
		panic(err)
	}
	content <- dest.String()
}
