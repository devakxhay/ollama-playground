package ollama

type ChatHistory struct {
	messages []Message
}

func NewChatHistory() *ChatHistory {
	return &ChatHistory{
		messages: []Message{},
	}
}

func (c *ChatHistory) addMessage(m Message) {
	c.messages = append(c.messages, m)
}

func (c *ChatHistory) AppendUserMessage(message string) {
	c.addMessage(Message{
		Role:    "user",
		Content: message,
	})
}

func (c *ChatHistory) AppendAssistantMessage(message string) {
	c.addMessage(Message{
		Role:    "assistant",
		Content: message,
	})
}

func (c *ChatHistory) GetMessages() []Message {
	return c.messages
}
