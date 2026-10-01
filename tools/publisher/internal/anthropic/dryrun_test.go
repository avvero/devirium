package anthropic

import (
	"bytes"
	"strings"
	"testing"
)

func TestDryRunPrintsCurlAndReturnsCorrect(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("https_proxy", "")
	t.Setenv("HTTP_PROXY", "")
	t.Setenv("http_proxy", "")
	var buf bytes.Buffer
	c := NewDryRun("https://api.anthropic.com", "sk-ant-secret-token", &buf)
	got, err := c.Complete("claude-sonnet-5", "prompt body")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Note is correct" {
		t.Errorf("got %q", got)
	}
	out := buf.String()
	if !strings.Contains(out, "curl -sS -X POST") {
		t.Errorf("missing curl: %s", out)
	}
	if !strings.Contains(out, "https://api.anthropic.com/v1/messages") {
		t.Errorf("missing URL: %s", out)
	}
	if !strings.Contains(out, "x-api-key: sk-***") {
		t.Errorf("token not masked: %s", out)
	}
	if strings.Contains(out, "sk-ant-secret-token") {
		t.Errorf("token leaked: %s", out)
	}
	if !strings.Contains(out, `"model":"claude-sonnet-5"`) {
		t.Errorf("missing model: %s", out)
	}
}

func TestDryRunIncludesProxyFlag(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "http://10.0.1.80:8118")
	var buf bytes.Buffer
	c := NewDryRun("https://api.anthropic.com", "sk-ant-x", &buf)
	if _, err := c.Complete("claude-sonnet-5", "p"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "-x 'http://10.0.1.80:8118'") {
		t.Errorf("missing proxy flag: %s", buf.String())
	}
}

func TestDryRunEmptyToken(t *testing.T) {
	var buf bytes.Buffer
	c := NewDryRun("https://api.anthropic.com", "", &buf)
	if _, err := c.Complete("claude-sonnet-5", "x"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "<ANTHROPIC_API_KEY>") {
		t.Errorf("missing placeholder: %s", buf.String())
	}
}
