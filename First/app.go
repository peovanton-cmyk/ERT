package first

import "fmt"

type M interface {
	Message()
}

type I struct {
	message string
}

func (i I) Message() {
	fmt.Println(i.message)
}
