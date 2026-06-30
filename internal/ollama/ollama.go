package ollama

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type GenerateResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ChatResponse struct {
	Message Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaClient struct {
	URL string `json:"url"`
}

func NewOllamaClient(ollamaUrl string) *OllamaClient {
	return &OllamaClient{
		URL: ollamaUrl,
	}
}

func (o *OllamaClient) Generate(model string, prompt string) (string, error) {
	reqBody := GenerateRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jb, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Post(o.URL, "application/json", bytes.NewReader(jb))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		rb, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ollama returned status: %s / body: %s", resp.Status, string(rb))
	}

	rb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var generateRes GenerateResponse

	if err := json.Unmarshal(rb, &generateRes); err != nil {
		return "", fmt.Errorf("Error in decoding response. Error: %w / Raw: %s", err, string(rb))
	}

	return generateRes.Response, nil
}

func (o *OllamaClient) Chat(model string, messages []Message) (string, error) {
	reqBody := ChatRequest{
		Model:    model,
		Messages: messages,
		Stream:   false,
	}

	jb, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Post(o.URL, "application/json", bytes.NewReader(jb))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		rb, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ollama returned status: %s / body: %s", resp.Status, string(rb))
	}

	rb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatResponse ChatResponse

	if err := json.Unmarshal(rb, &chatResponse); err != nil {
		return "", fmt.Errorf("Error in decoding response. Error: %w / Raw: %s", err, string(rb))
	}

	return chatResponse.Message.Content, nil
}
