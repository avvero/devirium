package telegram

import (
	"bytes"
	"strings"
	"testing"
)

func TestDryRunSendMessagePrintsCurl(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("https_proxy", "")
	t.Setenv("HTTP_PROXY", "")
	t.Setenv("http_proxy", "")
	var buf bytes.Buffer
	c := NewDryRun("https://api.telegram.org", "TOKEN", &buf)
	res, err := c.SendMessage("chat", "hello 'world'", "MarkdownV2")
	if err != nil {
		t.Fatal(err)
	}
	if res.MessageID != 1 {
		t.Errorf("id=%d", res.MessageID)
	}
	out := buf.String()
	if !strings.Contains(out, "curl -sS -X POST") {
		t.Errorf("missing curl: %s", out)
	}
	if !strings.Contains(out, "https://api.telegram.org/botTOKEN/sendMessage?disable_web_page_preview=true") {
		t.Errorf("missing URL: %s", out)
	}
	if !strings.Contains(out, `"chat_id":"chat"`) {
		t.Errorf("missing chat_id: %s", out)
	}
	if !strings.Contains(out, `'\''world'\''`) {
		t.Errorf("single quotes not escaped: %s", out)
	}
}

func TestDryRunSendPhotoPrintsCurl(t *testing.T) {
	var buf bytes.Buffer
	c := NewDryRun("https://api.telegram.org", "TOKEN", &buf)
	if _, err := c.SendPhoto("chat", "https://x/pic.png", "cap", "MarkdownV2"); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	if !strings.Contains(out, "/botTOKEN/sendPhoto") {
		t.Errorf("missing URL: %s", out)
	}
	if !strings.Contains(out, `"photo":"https://x/pic.png"`) {
		t.Errorf("missing photo: %s", out)
	}
}

func TestDryRunIncludesProxyFlag(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "http://10.0.1.80:8118")
	var buf bytes.Buffer
	c := NewDryRun("https://api.telegram.org", "TOKEN", &buf)
	if _, err := c.SendMessage("c", "hi", "MarkdownV2"); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), "-x 'http://10.0.1.80:8118'") {
		t.Errorf("missing proxy flag: %s", buf.String())
	}
}

func TestDryRunNoProxyFlagWhenUnset(t *testing.T) {
	t.Setenv("HTTPS_PROXY", "")
	t.Setenv("https_proxy", "")
	t.Setenv("HTTP_PROXY", "")
	t.Setenv("http_proxy", "")
	var buf bytes.Buffer
	c := NewDryRun("https://api.telegram.org", "TOKEN", &buf)
	if _, err := c.SendMessage("c", "hi", "MarkdownV2"); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(buf.String(), " -x ") {
		t.Errorf("unexpected proxy flag: %s", buf.String())
	}
}

func TestDryRunSequentialMessageIDs(t *testing.T) {
	var buf bytes.Buffer
	c := NewDryRun("u", "t", &buf)
	r1, _ := c.SendMessage("c", "a", "m")
	r2, _ := c.SendMessage("c", "b", "m")
	if r1.MessageID != 1 || r2.MessageID != 2 {
		t.Errorf("ids: %d %d", r1.MessageID, r2.MessageID)
	}
}
