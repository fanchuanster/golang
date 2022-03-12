package main

import (
	"fmt"
	"hello/greetings"
	"log"
)

func main() {
	names := []string{"Mr", "Mrs", "Lady"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
