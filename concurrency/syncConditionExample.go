package main

import (
	"fmt"
	"sync"
)

type Profile struct {
	Id int
}

var Profiles []Profile

var Done bool = false

func main() {
	condition := sync.Cond{L: &sync.Mutex{}}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go ShowProfiles(&condition, &wg)

	wg.Add(1)
	go ShowProfiles(&condition, &wg)

	for i := 1; i < 51; i++ {
		wg.Add(1)

		go CrawlProfile(i, &condition, &wg)
	}
	wg.Wait()
}

func CrawlProfile(id int, cond *sync.Cond, wg *sync.WaitGroup) {
	defer cond.L.Unlock()
	cond.L.Lock()
	Profiles = append(Profiles, Profile{Id: id})
	wg.Done()
	if len(Profiles) == 50 {
		fmt.Println("Crawling 50 Profiles Done")
		Done = true
		cond.Broadcast()
	}

}

func ShowProfiles(cond *sync.Cond, wg *sync.WaitGroup) {
	defer cond.L.Unlock()
	cond.L.Lock()
	for !Done {
		cond.Wait()
	}

	fmt.Printf("%v\n", Profiles)
	wg.Done()

}
