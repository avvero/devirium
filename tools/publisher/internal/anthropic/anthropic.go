package anthropic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiVersion = "2023-06-01"

const defaultMaxTokens = 2048

type Client struct {
	baseURL string
	apiKey  string
	http    *http.Client
}

func New(baseURL, apiKey string, hc *http.Client) *Client {
	if hc == nil {
		hc = http.DefaultClient
	}
	return &Client{baseURL: baseURL, apiKey: apiKey, http: hc}
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type messagesReq struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []message `json:"messages"`
}
type contentBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type messagesErr struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
type messagesResp struct {
	Content []contentBlock `json:"content"`
	Error   *messagesErr   `json:"error"`
}

func (c *Client) Complete(model, prompt string) (string, error) {
	body, _ := json.Marshal(messagesReq{Model: model, MaxTokens: defaultMaxTokens, Messages: []message{{Role: "user", Content: prompt}}})
	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/v1/messages", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", apiVersion)
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var parsed messagesResp
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", fmt.Errorf("anthropic: bad response %d: %s", resp.StatusCode, string(raw))
	}
	if parsed.Error != nil {
		return "", fmt.Errorf("anthropic: %s: %s", parsed.Error.Type, parsed.Error.Message)
	}
	if len(parsed.Content) == 0 {
		return "", fmt.Errorf("anthropic: no content")
	}
	return parsed.Content[len(parsed.Content)-1].Text, nil
}
