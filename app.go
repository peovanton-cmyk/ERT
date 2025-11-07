package main

import (
	"fmt"

	"github.com/peovanton-cmyk/ERT/Hello"
)

func main() {
	i := I{
		message: "Hello First tutorial Golang",
	}
	i.Message()
	Hello.Hello("Anton")
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
