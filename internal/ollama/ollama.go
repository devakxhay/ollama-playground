package ollama

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GenerateRequest struct {
	Model   string         `json:"model"`
	Prompt  string         `json:"prompt"`
	Stream  bool           `json:"stream"`
	Format  string         `json:"format,omitempty"`
	Options map[string]any `json:"options,omitempty"`
}

type GenerateResponse struct {
	Model    string `json:"model"`
	Response string `json:"response"`
}

type ChatRequest struct {
	Model    string         `json:"model"`
	Messages []Message      `json:"messages"`
	Options  map[string]any `json:"options,omitempty"`
	Stream   bool           `json:"stream"`
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

	resp, err := httpClient.Post(o.URL+"/api/generate", "application/json", bytes.NewReader(jb))
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

func (o *OllamaClient) Chat(model string, messages []Message, options map[string]any) (string, error) {
	resp, err := postChat(o.URL+"/api/chat", model, messages, options, false)
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

func (o *OllamaClient) ChatStream(model string, messages []Message) (<-chan string, <-chan error, error) {
	resp, err := postChat(o.URL+"/api/chat", model, messages, nil, true)
	if err != nil {
		return nil, nil, err
	}

	out := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer resp.Body.Close()
		defer close(out)
		defer close(errCh)
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			var chunk ChatResponse
			if err := json.Unmarshal(scanner.Bytes(), &chunk); err != nil {
				continue
			}
			out <- chunk.Message.Content
		}
		if err := scanner.Err(); err != nil {
			errCh <- err
		}
	}()

	return out, errCh, nil
}

func (o *OllamaClient) GenerateJsonResponse(model string, prompt string) (<-chan string, <-chan error, error) {
	reqBody := GenerateRequest{
		Model:  model,
		Prompt: prompt,
		Stream: true,
		Format: "json",
		Options: map[string]any{
			"temperature": 0.0,
		},
	}

	jb, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	resp, err := httpClient.Post(o.URL+"/api/generate", "application/json", bytes.NewReader(jb))
	if err != nil {
		return nil, nil, err
	}

	out := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer resp.Body.Close()
		defer close(out)
		defer close(errCh)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			var chunk GenerateResponse
			if err := json.Unmarshal(scanner.Bytes(), &chunk); err != nil {
				continue
			}
			out <- chunk.Response
		}
		if err := scanner.Err(); err != nil {
			errCh <- err
		}
	}()

	return out, errCh, nil
}

func (o *OllamaClient) ChatSystemRole(model string, messages []Message) (string, error) {
	return o.Chat(model, messages, map[string]any{
		"temperature": 0.0,
	})
}
