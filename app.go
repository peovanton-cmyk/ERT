package main

import (
	"github.com/peovanton-cmyk/ERT/Hello"

	"fmt"
)

func main() {
	var i (I)
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
