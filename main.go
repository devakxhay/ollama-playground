package main

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/modules"
)

func main() {
	fmt.Println("------------ GENERATE API EXAMPLE ------------")

	modules.GenerateApiExample("what is TLS?")
	modules.GenerateApiExample("Can you elaborate it?")

	fmt.Println("------------ CHAT API EXAMPLE ------------")

	modules.ChatApiExample("what is TLS?")
	modules.ChatApiExample("Can you elaborate it?")
}
