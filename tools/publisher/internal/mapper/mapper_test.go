package mapper

import (
	"strings"
	"testing"
)

func newMapper() *Mapper { return New("https://devirium.com") }

func TestSimpleWithFileName(t *testing.T) {
	got, err := newMapper().Map("Test Note.md", "notes/Test Note.md", "Simple test content", nil, false)
	if err != nil {
		t.Fatal(err)
	}
	want := "*Test Note*\n\nSimple test content"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNoFileName(t *testing.T) {
	got, err := newMapper().Map("", "", "Content without title", nil, false)
	if err != nil {
		t.Fatal(err)
	}
	if got != "Content without title" {
		t.Errorf("got %q", got)
	}
}

func TestTrimTextLimit(t *testing.T) {
	content := strings.Repeat("A", 3500)
	got, err := newMapper().Map("Long Note.md", "notes/Long Note.md", content, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Полный текст в [Long Note](https://devirium.com/notes/Long-Note)") {
		t.Errorf("missing tail link: %s", got)
	}
	if len(got) >= len(content) {
		t.Errorf("expected trimmed")
	}
}

func TestTrimPhotoLimit(t *testing.T) {
	content := strings.Repeat("B", 800)
	got, err := newMapper().Map("Photo Note.md", "photos/Photo Note.md", content, nil, true)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "Полный текст в [Photo Note](https://devirium.com/photos/Photo-Note)") {
		t.Errorf("missing tail link: %s", got)
	}
}

func TestRemovesImages(t *testing.T) {
	content := "Text before ![alt text](image.jpg) text after ![](another.png)  more text"
	got, err := newMapper().Map("Note.md", "Note.md", content, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "![") || strings.Contains(got, "image.jpg") || strings.Contains(got, "another.png") {
		t.Errorf("image left: %s", got)
	}
	for _, needle := range []string{"Text before", "text after", "more text"} {
		if !strings.Contains(got, needle) {
			t.Errorf("missing %q in %q", needle, got)
		}
	}
}

func TestEscapeSpecialChars(t *testing.T) {
	cases := map[string]string{
		"Simple text":            "Simple text",
		"Text with_underscore":   "Text with\\_underscore",
		"Text with*asterisk":     "Text with\\*asterisk",
		"Text with-dash":         "Text with\\-dash",
		"Text with.period":       "Text with\\.period",
		"Text with[brackets]":    "Text with\\[brackets\\]",
		"Text with(parentheses)": "Text with\\(parentheses\\)",
		"Text with{braces}":      "Text with\\{braces\\}",
		"Text with|pipe":         "Text with\\|pipe",
		"Text with#hash":         "Text with\\#hash",
		"Text with+plus":         "Text with\\+plus",
		"Text with=equals":       "Text with\\=equals",
		"Text with~tilde":        "Text with\\~tilde",
		"Text with!exclamation":  "Text with\\!exclamation",
	}
	for in, want := range cases {
		got, err := newMapper().Map("", "", in, nil, false)
		if err != nil {
			t.Fatalf("in=%q err=%v", in, err)
		}
		if got != want {
			t.Errorf("in=%q got=%q want=%q", in, got, want)
		}
	}
}

func TestBackticks(t *testing.T) {
	cases := map[string]string{
		"Text with`single backtick":             "Text with\\`single backtick",
		"Code `inline code` text":               "Code `inline code` text",
		"Code ```block code``` text":            "Code ```block code``` text",
		"Multiple `code1` and `code2` blocks":   "Multiple `code1` and `code2` blocks",
	}
	for in, want := range cases {
		got, _ := newMapper().Map("", "", in, nil, false)
		if got != want {
			t.Errorf("in=%q got=%q want=%q", in, got, want)
		}
	}
}

func TestReplaceZettelLinks(t *testing.T) {
	content := "Text with [[Note 1]] and [[Note-2]] links"
	links := map[string]string{
		"Note 1":  "2023/note-1",
		"Note-2":  "2023/note-2",
	}
	got, err := newMapper().Map("", "", content, links, false)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "[Note 1](https://devirium.com/2023/note-1)") {
		t.Errorf("missing Note 1: %s", got)
	}
	if !strings.Contains(got, "[Note\\-2](https://devirium.com/2023/note-2)") {
		t.Errorf("missing Note-2: %s", got)
	}
	if strings.Contains(got, "[[") {
		t.Errorf("wikilink left: %s", got)
	}
}

func TestUnresolvedLink(t *testing.T) {
	_, err := newMapper().Map("", "", "Text with [[Unresolved Link]] content", map[string]string{"Other": "x"}, false)
	if err == nil || !strings.Contains(err.Error(), "Can't resolve link [[Unresolved Link]]") {
		t.Fatalf("want unresolved error, got %v", err)
	}
}

func TestPreservesMarkdownLinks(t *testing.T) {
	got, _ := newMapper().Map("", "", "Check out [this link](http://example.com) and [another](https://test.com)", nil, false)
	if !strings.Contains(got, "[this link](http://example.com)") || !strings.Contains(got, "[another](https://test.com)") {
		t.Errorf("md link mangled: %s", got)
	}
}

func TestMixedContent(t *testing.T) {
	content := "Text with [[Link]] and ![](image.jpg) plus `code` and *bold*"
	got, err := newMapper().Map("Mixed.md", "mixed/Mixed.md", content, map[string]string{"Link": "2023/link"}, true)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(got, "*Mixed*\n\n") {
		t.Errorf("bad prefix: %s", got)
	}
	if !strings.Contains(got, "[Link](https://devirium.com/2023/link)") {
		t.Errorf("missing link: %s", got)
	}
	if strings.Contains(got, "![") {
		t.Errorf("image left: %s", got)
	}
	if !strings.Contains(got, "`code`") {
		t.Errorf("code lost: %s", got)
	}
	if !strings.Contains(got, "\\*bold\\*") {
		t.Errorf("bold not escaped: %s", got)
	}
}

func TestEmptyContent(t *testing.T) {
	got, _ := newMapper().Map("Empty.md", "", "", nil, false)
	if got != "*Empty*\n\n" {
		t.Errorf("got %q", got)
	}
}

func TestPhotoURL(t *testing.T) {
	if got := newMapper().URLForPhoto("photos/test-image.jpg"); got != "https://devirium.com/photos/test-image.jpg" {
		t.Errorf("got %q", got)
	}
}

func TestUnicode(t *testing.T) {
	got, _ := newMapper().Map("Unicode Note.md", "", "Русский текст with émojis 🚀 and symbols ñáéíóú", nil, false)
	for _, needle := range []string{"*Unicode Note*", "Русский текст", "🚀", "ñáéíóú"} {
		if !strings.Contains(got, needle) {
			t.Errorf("missing %q in %q", needle, got)
		}
	}
}

func TestEmptyWikilink(t *testing.T) {
	_, err := newMapper().Map("", "", "Text with [[]] empty link", map[string]string{}, false)
	if err == nil || !strings.Contains(err.Error(), "Can't resolve link [[]]") {
		t.Fatalf("want empty-link error, got %v", err)
	}
}
