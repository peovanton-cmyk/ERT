package First

import (
	"fmt"

	"github.com/peovanton-cmyk/ERT/Hello"
)

func First() {
	var i (I)
	i.message = Hello.Hello("Education Right Tutorial")
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
