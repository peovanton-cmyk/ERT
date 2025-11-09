package main

import (
	"github.com/peovanton-cmyk/ERT/First"
	"github.com/peovanton-cmyk/ERT/Hello"
)

type S struct {
	F First.I
}

func main() {
	var s (S)
	// s.F = First.I()
	Hello.Hello("This first tutorial")
	s.F.Message()
}
