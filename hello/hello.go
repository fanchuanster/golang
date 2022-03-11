package main

import (
	"fmt"

	"hello/greetings"
)

func main() {
	message := greetings.Hello("Mr")
	fmt.Println(message)
}
