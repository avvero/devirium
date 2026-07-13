package openai

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
	c := NewDryRun("https://api.openai.com", "sk-secret-token", &buf)
	got, err := c.Complete("gpt-4", "prompt body")
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
	if !strings.Contains(out, "https://api.openai.com/v1/chat/completions") {
		t.Errorf("missing URL: %s", out)
	}
	if !strings.Contains(out, "Bearer sk-***") {
		t.Errorf("token not masked: %s", out)
	}
	if strings.Contains(out, "sk-secret-token") {
		t.Errorf("token leaked: %s", out)
	}
	if !strings.Contains(out, `"model":"gpt-4"`) {
		t.Errorf("missing model: %s", out)
	}
}

func TestDryRunOpenaiIncludesProxyFlag(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "http://10.0.1.80:8118")
	var buf bytes.Buffer
	c := NewDryRun("https://api.openai.com", "sk-x", &buf)
	if _, err := c.Complete("gpt-4", "p"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "-x 'http://10.0.1.80:8118'") {
		t.Errorf("missing proxy flag: %s", buf.String())
	}
}

func TestDryRunEmptyToken(t *testing.T) {
	var buf bytes.Buffer
	c := NewDryRun("https://api.openai.com", "", &buf)
	if _, err := c.Complete("gpt-4", "x"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "<OPENAI_TOKEN>") {
		t.Errorf("missing placeholder: %s", buf.String())
	}
}
