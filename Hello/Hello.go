package Hello

import "fmt"

func Hello(str string) string {
	message := fmt.Sprintf("Hello, %v. Welcome.", str)
	return message
}
