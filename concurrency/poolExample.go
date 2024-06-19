package main

import (
	"sync"
)

type DbConnectionConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

var DbConnectionPool sync.Pool = sync.Pool{New: func() interface{} {
	return &DbConnectionConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "laskdofwqfqw",
	}
}}

func main() {

	var dbConnection *DbConnectionConfig
	//wrong method with bad performance and wasting resource
	for i := 0; i < 10; i++ {
		dbConnection = &DbConnectionConfig{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "laskdofwqfqw",
		}
		dbConnection.Username = "root"
	}

	//correct method with good performance and saving resource
	for i := 0; i < 10; i++ {
		dbConnection = DbConnectionPool.Get().(*DbConnectionConfig)

		dbConnection.Username = "root"
		DbConnectionPool.Put(dbConnection)

	}

}
