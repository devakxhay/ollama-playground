package modules

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/ollama"
)

const ollamaUrl = "http://localhost:11434"
const model = "gemma3:270m"

func GenerateApiExample(prompt string) {
	oc := ollama.NewOllamaClient(ollamaUrl + "/api/generate")

	content, _ := oc.Generate(model, prompt)

	fmt.Printf("Response From Ollama: %s\n", content)
}

var messages = []ollama.Message{}

func ChatApiExample(message string) {
	oc := ollama.NewOllamaClient(ollamaUrl + "/api/chat")

	messages = append(messages, ollama.Message{
		Role:    "user",
		Content: message,
	})

	content, _ := oc.Chat(model, messages)

	fmt.Printf("Response From Ollama: %s\n", content)

	messages = append(messages, ollama.Message{
		Role:    "assistant",
		Content: content,
	})
}
