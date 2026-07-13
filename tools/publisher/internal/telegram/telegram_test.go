package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendMessage(t *testing.T) {
	var gotBody map[string]any
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path + "?" + r.URL.RawQuery
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		_, _ = w.Write([]byte(`{"ok":true,"result":{"message_id":42}}`))
	}))
	defer srv.Close()

	c := New(srv.URL, "TOKEN", srv.Client())
	got, err := c.SendMessage("chat", "hello", "MarkdownV2")
	if err != nil {
		t.Fatal(err)
	}
	if got.MessageID != 42 {
		t.Errorf("id=%d", got.MessageID)
	}
	if !strings.Contains(gotPath, "/TOKEN/sendMessage") {
		t.Errorf("path=%s", gotPath)
	}
	if !strings.Contains(gotPath, "disable_web_page_preview=true") {
		t.Errorf("missing preview flag: %s", gotPath)
	}
	if gotBody["chat_id"] != "chat" || gotBody["text"] != "hello" || gotBody["parse_mode"] != "MarkdownV2" {
		t.Errorf("body=%v", gotBody)
	}
}

func TestSendPhoto(t *testing.T) {
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		_, _ = w.Write([]byte(`{"ok":true,"result":{"message_id":7}}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "TOKEN", srv.Client())
	got, err := c.SendPhoto("chat", "https://x/pic.png", "cap", "MarkdownV2")
	if err != nil {
		t.Fatal(err)
	}
	if got.MessageID != 7 {
		t.Errorf("id=%d", got.MessageID)
	}
	if gotBody["photo"] != "https://x/pic.png" || gotBody["caption"] != "cap" {
		t.Errorf("body=%v", gotBody)
	}
}

func TestSendMessageErrorBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"ok":false,"description":"bad markdown"}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "T", srv.Client())
	_, err := c.SendMessage("c", "t", "MarkdownV2")
	if err == nil || !strings.Contains(err.Error(), "bad markdown") {
		t.Fatalf("want bad-markdown error, got %v", err)
	}
}
