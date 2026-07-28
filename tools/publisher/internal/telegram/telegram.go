package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

type SendMessageResult struct {
	MessageID int64 `json:"message_id"`
}

type response struct {
	OK          bool              `json:"ok"`
	Description string            `json:"description"`
	Result      SendMessageResult `json:"result"`
}

type sendMessageReq struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type sendPhotoReq struct {
	ChatID    string `json:"chat_id"`
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

func (c *Client) SendMessage(chatID, text, parseMode string) (SendMessageResult, error) {
	body, _ := json.Marshal(sendMessageReq{ChatID: chatID, Text: text, ParseMode: parseMode})
	url := fmt.Sprintf("%s/bot%s/sendMessage?disable_web_page_preview=true", c.baseURL, c.token)
	return c.post(url, body)
}

func (c *Client) SendPhoto(chatID, photo, caption, parseMode string) (SendMessageResult, error) {
	body, _ := json.Marshal(sendPhotoReq{ChatID: chatID, Photo: photo, Caption: caption, ParseMode: parseMode})
	url := fmt.Sprintf("%s/bot%s/sendPhoto", c.baseURL, c.token)
	return c.post(url, body)
}

func (c *Client) SendPhotoFile(chatID, localPath, caption, parseMode string) (SendMessageResult, error) {
	f, err := os.Open(localPath)
	if err != nil {
		return SendMessageResult{}, fmt.Errorf("open photo: %w", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	_ = w.WriteField("chat_id", chatID)
	_ = w.WriteField("caption", caption)
	_ = w.WriteField("parse_mode", parseMode)
	fw, err := w.CreateFormFile("photo", filepath.Base(localPath))
	if err != nil {
		return SendMessageResult{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return SendMessageResult{}, err
	}
	w.Close()

	url := fmt.Sprintf("%s/bot%s/sendPhoto", c.baseURL, c.token)
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return SendMessageResult{}, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := c.http.Do(req)
	if err != nil {
		return SendMessageResult{}, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var parsed response
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return SendMessageResult{}, fmt.Errorf("telegram: bad response %d: %s", resp.StatusCode, string(raw))
	}
	if !parsed.OK {
		return SendMessageResult{}, fmt.Errorf("telegram: %s\ncaption: %s", parsed.Description, caption)
	}
	return parsed.Result, nil
}

func (c *Client) post(url string, body []byte) (SendMessageResult, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return SendMessageResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return SendMessageResult{}, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var parsed response
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return SendMessageResult{}, fmt.Errorf("telegram: bad response %d: %s\nrequest body: %s", resp.StatusCode, string(raw), string(body))
	}
	if !parsed.OK {
		return SendMessageResult{}, fmt.Errorf("telegram: %s\nrequest body: %s", parsed.Description, string(body))
	}
	return parsed.Result, nil
}
