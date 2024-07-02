package main

import (
	"errors"
	"fmt"
)

type HttpError struct {
	StatusCode int
	Message    string
	Err        error
}

func main() {
	res, err := getRequest("")
	if err != nil {
		fmt.Println(err)
		fmt.Println(errors.Unwrap(err))

		return
	}
	fmt.Println(res)
}

// func getRequest(url string) (response string, err error) {
// 	if len(url) == 0 {
// 		return "", errors.New("status code:400 message: url is empty")
// 	}
// 	return "successful", nil

// }

// func getRequest(url string) (response string, err error) {
// 	if len(url) == 0 {
// 		return "", fmt.Errorf("status code:%d message: %s", 400, "url is empty")
// 	}
// 	return "successful", nil

// }

func getRequest(url string) (response string, err error) {
	if len(url) == 0 {
		return "", &HttpError{StatusCode: 400, Message: "url is empty", Err: errors.New("getRequest function has error")}
	}
	return "successful", nil

}

func (err *HttpError) Error() string {
	return fmt.Sprintf("status code:%d message:%s main error:%s", err.StatusCode, err.Message, err.Err.Error())

}

func (err *HttpError) Unwrap() error {
	return err.Err

}
