package main

import (
	"Hello"

	"fmt"
)

func main() {
	//TODO #1 exchange type I not use parametr message use change in variable where I isset type
	var i (I)
	// var i (I) = I{
	// 	message: Hello.Hello("This first tutorial"),
	// }
	i.message = Hello.Hello("This first tutorial")

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
