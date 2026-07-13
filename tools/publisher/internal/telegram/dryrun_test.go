package telegram

import (
	"bytes"
	"strings"
	"testing"
)

func TestDryRunSendMessagePrintsCurl(t *testing.T) {
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
	if !strings.Contains(out, "https://api.telegram.org/TOKEN/sendMessage?disable_web_page_preview=true") {
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
	if !strings.Contains(out, "/TOKEN/sendPhoto") {
		t.Errorf("missing URL: %s", out)
	}
	if !strings.Contains(out, `"photo":"https://x/pic.png"`) {
		t.Errorf("missing photo: %s", out)
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
