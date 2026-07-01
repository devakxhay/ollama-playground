package ollama

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

func postChat(URL string, model string, messages []Message, options map[string]any, stream bool) (*http.Response, error) {
	reqBody := ChatRequest{
		Model:    model,
		Messages: messages,
		Options:  options,
		Stream:   stream,
	}

	jb, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	return httpClient.Post(URL, "application/json", bytes.NewReader(jb))
}
