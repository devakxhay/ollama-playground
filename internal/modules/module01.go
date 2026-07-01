package modules

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/ollama"
)

func GenerateApiExample(prompt string) {

	content, _ := ollama.OC.Generate(ollama.Model, prompt)

	fmt.Printf("Response From Ollama: %s\n", content)
}

var messages = []ollama.Message{}

func ChatApiExample(message string) {

	messages = append(messages, ollama.Message{
		Role:    "user",
		Content: message,
	})

	content, _ := ollama.OC.Chat(ollama.Model, messages, nil)

	fmt.Printf("Response From Ollama: %s\n", content)

	messages = append(messages, ollama.Message{
		Role:    "assistant",
		Content: content,
	})
}
