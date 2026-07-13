package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// DryRunClient prints the curl command instead of calling OpenAI.
// It always answers "Note is correct" so the publisher takes the happy path.
type DryRunClient struct {
	baseURL string
	token   string
	out     io.Writer
}

func NewDryRun(baseURL, token string, out io.Writer) *DryRunClient {
	return &DryRunClient{baseURL: baseURL, token: token, out: out}
}

func (c *DryRunClient) Complete(model, prompt string) (string, error) {
	body, _ := json.Marshal(completionReq{Model: model, Messages: []message{{Role: "user", Content: prompt}}})
	fmt.Fprintf(c.out,
		"[dry-run] curl -sS -X POST %s \\\n  -H 'Content-Type: application/json' \\\n  -H 'Authorization: Bearer %s' \\\n  -d %s\n",
		shellQuote(c.baseURL+"/v1/chat/completions"), maskedToken(c.token), shellQuote(string(body)))
	return "Note is correct", nil
}

func shellQuote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", `'\''`) + "'"
}

func maskedToken(t string) string {
	if t == "" {
		return "<OPENAI_TOKEN>"
	}
	if len(t) <= 4 {
		return "***"
	}
	return t[:3] + "***"
}
