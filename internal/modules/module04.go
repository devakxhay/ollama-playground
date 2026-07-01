package modules

import (
	"fmt"

	"github.com/devakxhay/ollama-playground/internal/ollama"
)

var ch *ollama.ChatHistory

func InitChat(role string) {
	ch = ollama.NewChatHistory()
	ch.AppendSystemMessage(role)
}

func SystemRoleExample(message string) {
	oc := ollama.OC

	ch.AppendUserMessage(message)
	content, _ := oc.ChatSystemRole(ollama.Model, ch.GetMessages())
	fmt.Println(content)

	ch.AppendAssistantMessage(content)
}
