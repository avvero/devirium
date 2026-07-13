package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// DryRunClient prints the curl command that would have been executed
// instead of calling the Telegram API.
type DryRunClient struct {
	baseURL string
	token   string
	out     io.Writer
	seq     int64
}

func NewDryRun(baseURL, token string, out io.Writer) *DryRunClient {
	return &DryRunClient{baseURL: baseURL, token: token, out: out}
}

func (c *DryRunClient) SendMessage(chatID, text, parseMode string) (SendMessageResult, error) {
	body, _ := json.Marshal(sendMessageReq{ChatID: chatID, Text: text, ParseMode: parseMode})
	url := fmt.Sprintf("%s/%s/sendMessage?disable_web_page_preview=true", c.baseURL, c.token)
	c.printCurl(url, body)
	return c.nextResult(), nil
}

func (c *DryRunClient) SendPhoto(chatID, photo, caption, parseMode string) (SendMessageResult, error) {
	body, _ := json.Marshal(sendPhotoReq{ChatID: chatID, Photo: photo, Caption: caption, ParseMode: parseMode})
	url := fmt.Sprintf("%s/%s/sendPhoto", c.baseURL, c.token)
	c.printCurl(url, body)
	return c.nextResult(), nil
}

func (c *DryRunClient) printCurl(url string, body []byte) {
	fmt.Fprintf(c.out, "[dry-run] curl -sS -X POST %s \\\n  -H 'Content-Type: application/json' \\\n  -d %s\n",
		shellQuote(url), shellQuote(string(body)))
}

func (c *DryRunClient) nextResult() SendMessageResult {
	c.seq++
	return SendMessageResult{MessageID: c.seq}
}

// shellQuote wraps s in single quotes, escaping embedded single quotes.
func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}
