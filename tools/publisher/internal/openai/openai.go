package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL string
	token   string
	http    *http.Client
}

func New(baseURL, token string, hc *http.Client) *Client {
	if hc == nil {
		hc = http.DefaultClient
	}
	return &Client{baseURL: baseURL, token: token, http: hc}
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type completionReq struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}
type completionChoice struct {
	Message message `json:"message"`
}
type completionErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type completionResp struct {
	Choices []completionChoice `json:"choices"`
	Error   *completionErr     `json:"error"`
}

func (c *Client) Complete(model, prompt string) (string, error) {
	body, _ := json.Marshal(completionReq{Model: model, Messages: []message{{Role: "user", Content: prompt}}})
	req, err := http.NewRequest(http.MethodPost, c.baseURL+"/v1/chat/completions", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var parsed completionResp
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", fmt.Errorf("openai: bad response %d: %s", resp.StatusCode, string(raw))
	}
	if parsed.Error != nil {
		return "", fmt.Errorf("openai: %s: %s", parsed.Error.Code, parsed.Error.Message)
	}
	if len(parsed.Choices) == 0 {
		return "", fmt.Errorf("openai: no choices")
	}
	return parsed.Choices[len(parsed.Choices)-1].Message.Content, nil
}
