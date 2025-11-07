package main

import (
	"Hello"

	"fmt"
)

func main() {
	i := I{
		message: Hello.Hello("First Tutorial App"),
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
