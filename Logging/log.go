package main

import (
	"log"
	"os"
)

var (
	WarnLogger  log.Logger
	ErrorLogger log.Logger
	InfoLogger  log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln("error in open log file:", err)
	}
	log.SetOutput(file)
	flags := log.Ldate | log.Ltime | log.Lshortfile
	log.SetFlags(flags)

	WarnLogger = *log.New(file, "warning: ", flags)
	ErrorLogger = *log.New(file, "Error: ", flags)
	InfoLogger = *log.New(file, "Info: ", flags)

}

func main() {
	// log.Println("start main func")

	InfoLogger.Println("start main func")

	sum(5, 6)
}

func sum(a, b int) {
	defer func() {
		err := recover()
		if err != nil {
			ErrorLogger.Println("Manual error with panic")
		}

	}()
	WarnLogger.Printf("sum numbers : %d , %d ", a, b)
	InfoLogger.Println("end sum func")

	panic("manual error")

}
