package Hello

import "fmt"

func Hello(str string) string {
	message := fmt.Sprintf("Привет, %v. Добро пожаловать.", str)
	return message
}
