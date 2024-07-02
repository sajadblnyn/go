package main

import (
	"fmt"
	"notification/externalServices"
	"notification/services"
)

func main() {
	order := services.CreateOrder(23, 23000, "sajadbalaniyan@gmail.com", externalServices.Email)

	fmt.Printf("%v created successfully.\n", order)

}
