package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return MyError{Message: "There was an error!!", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error()) //=> There was an error!!
}
