package main

import (
	"fmt"
	"log"

	"github.com/ambilykartha/learn-go/greetings"
)

func main() {
	message, err := greetings.Hello("Ambily")
	if err != nil {
		log.Fatal("Empty name")
	}
	fmt.Println(message)
}
