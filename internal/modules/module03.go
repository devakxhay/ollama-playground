package modules

import (
	"fmt"
	"strings"

	"github.com/devakxhay/ollama-playground/internal/ollama"
)

func GenerateJsonResponseExample(message string) {

	out, errCh, err := ollama.OC.GenerateJsonResponse(ollama.Model, message)
	if err != nil {
		fmt.Printf("Error starting stream: %v\n", err)
		return
	}

	var sb strings.Builder
	for token := range out {
		fmt.Print(token)
		sb.WriteString(token)
	}

	if err := <-errCh; err != nil {
		fmt.Printf("\nStream error: %v\n", err)
		return
	}
	fmt.Println()
}
