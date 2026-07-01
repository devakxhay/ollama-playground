package modules

import (
	"fmt"
	"strings"

	"github.com/devakxhay/ollama-playground/internal/ollama"
)

var history = []ollama.Message{}

func ChatStreamApiExample(message string) {
	history = append(history, ollama.Message{
		Role:    "user",
		Content: message,
	})

	out, errCh, err := ollama.OC.ChatStream(ollama.Model, history)
	if err != nil {
		fmt.Printf("Error starting stream: %v\n", err)
		return
	}

	var sb strings.Builder
	for token := range out {
		fmt.Print(token)
		sb.WriteString(token)
	}
	fullResponse := sb.String()

	if err := <-errCh; err != nil {
		fmt.Printf("\nStream error: %v\n", err)
		return
	}
	fmt.Println()

	history = append(history, ollama.Message{
		Role:    "assistant",
		Content: fullResponse,
	})
}
