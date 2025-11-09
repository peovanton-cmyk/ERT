package main

import (
	"github.com/peovanton-cmyk/ERT/Hello"
)

func main() {
	var i = First
	i.message = Hello.Hello("This first tutorial")
	i.Message()
}
