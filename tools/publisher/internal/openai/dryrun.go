package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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
	proxyFlag := ""
	if p := firstNonEmpty(os.Getenv("HTTPS_PROXY"), os.Getenv("https_proxy"), os.Getenv("HTTP_PROXY"), os.Getenv("http_proxy")); p != "" {
		proxyFlag = fmt.Sprintf(" -x %s", shellQuote(p))
	}
	fmt.Fprintf(c.out,
		"[dry-run] curl -sS%s -X POST %s \\\n  -H 'Content-Type: application/json' \\\n  -H 'Authorization: Bearer %s' \\\n  -d %s\n",
		proxyFlag, shellQuote(c.baseURL+"/v1/chat/completions"), maskedToken(c.token), shellQuote(string(body)))
	return "Note is correct", nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
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
