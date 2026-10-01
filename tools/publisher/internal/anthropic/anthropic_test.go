package anthropic

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestComplete(t *testing.T) {
	var gotBody map[string]any
	var gotKey, gotVersion string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey = r.Header.Get("x-api-key")
		gotVersion = r.Header.Get("anthropic-version")
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		_, _ = w.Write([]byte(`{"content":[{"type":"text","text":"Note is correct"}]}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "sk-ant-test", srv.Client())
	got, err := c.Complete("claude-sonnet-5", "prompt body")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Note is correct" {
		t.Errorf("got %q", got)
	}
	if gotKey != "sk-ant-test" {
		t.Errorf("x-api-key=%q", gotKey)
	}
	if gotVersion != apiVersion {
		t.Errorf("anthropic-version=%q", gotVersion)
	}
	if gotBody["model"] != "claude-sonnet-5" {
		t.Errorf("model=%v", gotBody["model"])
	}
}

func TestCompleteError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"type":"invalid_request_error","message":"nope"}}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "t", srv.Client())
	_, err := c.Complete("claude-sonnet-5", "x")
	if err == nil || !strings.Contains(err.Error(), "nope") {
		t.Fatalf("want error, got %v", err)
	}
}
