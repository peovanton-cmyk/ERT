package main

import (
	"github.com/peovanton-cmyk/ERT/Hello"

	"fmt"
)

func main() {
	Hello.Hello("Anton")
	i := I{
		message: "Hello First tutorial Golang",
	}
	i.Message()
}

type M interface {
	Message()
}

type I struct {
	message string
}

func (i I) Message() {
	fmt.Println(i.message)
}
