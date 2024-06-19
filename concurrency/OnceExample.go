package main

import (
	"fmt"
	"sync"
)

type Config struct {
	ConnectionString string
}

var (
	once   sync.Once = sync.Once{}
	config *Config
	mx     sync.Mutex = sync.Mutex{}
)

func main() {
	for i := 0; i < 10; i++ {
		GetConfig()

		fmt.Println(&config)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		fmt.Println("create config")
	})
	fmt.Println("get config")

	return config
}

// func GetConfigBySingleTon() *Config {
// 	if config != nil {
// 		fmt.Println("get config")

// 		return config
// 	}
// 	mx.Lock()
// 	defer mx.Unlock()

// 	config = &Config{}
// 	fmt.Println("create config")

// 	return config
// }
