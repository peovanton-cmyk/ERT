package main

import (
	"fmt"
)

func main() {
	i := I{"Hello First tutorial Golang"}
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
