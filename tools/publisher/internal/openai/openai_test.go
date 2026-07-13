package openai

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
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		_, _ = w.Write([]byte(`{"choices":[{"message":{"content":"Note is correct"}}]}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "sk-test", srv.Client())
	got, err := c.Complete("gpt-4", "prompt body")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Note is correct" {
		t.Errorf("got %q", got)
	}
	if gotAuth != "Bearer sk-test" {
		t.Errorf("auth=%q", gotAuth)
	}
	if gotBody["model"] != "gpt-4" {
		t.Errorf("model=%v", gotBody["model"])
	}
}

func TestCompleteError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":{"code":"bad_request","message":"nope"}}`))
	}))
	defer srv.Close()
	c := New(srv.URL, "t", srv.Client())
	_, err := c.Complete("gpt-4", "x")
	if err == nil || !strings.Contains(err.Error(), "nope") {
		t.Fatalf("want error, got %v", err)
	}
}
