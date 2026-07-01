package main

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/modules"
)

func main() {
	currentModule := 2

	if currentModule == 1 {
		fmt.Println("------------ GENERATE API EXAMPLE ------------")

		modules.GenerateApiExample("what is TLS?")
		modules.GenerateApiExample("Can you elaborate it?")

		fmt.Println("------------ CHAT API EXAMPLE ------------")

		modules.ChatApiExample("what is TLS?")
		modules.ChatApiExample("Can you elaborate it?")
	}

	if currentModule == 2 {
		fmt.Println("------------ CHAT STREAM API EXAMPLE ------------")

		fmt.Println("1/2 -- What is TLS?")
		modules.ChatStreamApiExample("what is TLS?")

		fmt.Println("2/2 -- Can you elaborate it?")
		modules.ChatStreamApiExample("Can you elaborate it?")
	}
}
