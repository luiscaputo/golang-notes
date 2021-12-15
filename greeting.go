package Hello

import "fmt"

func Hello(string name) string {
	message := fmt.Spintf("Hi, %v. Welcome!", name)
	return message
}