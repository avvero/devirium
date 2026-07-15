package mapper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var (
	zetLink   = regexp.MustCompile(`\[\[.*?\]\]`)
	mdLink    = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	imageLink = regexp.MustCompile(`!\[.*?\]\(.*?\)\s*`)
)

// Mapper turns a note into a MarkdownV2 Telegram message.
type Mapper struct {
	deviriumLink string
}

func New(deviriumLink string) *Mapper {
	return &Mapper{deviriumLink: deviriumLink}
}

// Map produces the Telegram MarkdownV2 body for the note.
// links: wikilink name -> resolved path (no extension, dashes).
func (m *Mapper) Map(fileName, filePath, content string, links map[string]string, hasPhoto bool) (string, error) {
	messageLimit := 3000
	if hasPhoto {
		messageLimit = 700
	}
	if len([]rune(content)) > messageLimit {
		runes := []rune(content)
		content = string(runes[:messageLimit])
		link := strings.ReplaceAll(strings.TrimSuffix(filePath, ".md"), " ", "-")
		display := strings.TrimSuffix(fileName, ".md")
		content += "...\n\nПолный текст в " + fmt.Sprintf("[%s](%s/%s)", escape(display), m.deviriumLink, link)
	}

	content = imageLink.ReplaceAllString(content, "")

	meta := map[string]string{}
	content = extractMeta(content, meta, zetLink, mdLink)
	content = escape(content)
	for k, v := range meta {
		content = strings.ReplaceAll(content, k, v)
	}

	if links != nil {
		for name, path := range links {
			url := fmt.Sprintf("[%s](%s/%s)", escape(name), m.deviriumLink, path)
			content = strings.ReplaceAll(content, fmt.Sprintf("[[%s]]", name), url)
		}
	}

	if unresolved := zetLink.FindString(content); unresolved != "" {
		return "", fmt.Errorf("Can't resolve link %s", unresolved)
	}

	if fileName != "" {
		return fmt.Sprintf("*%s*\n\n%s", escape(strings.TrimSuffix(fileName, ".md")), content), nil
	}
	return content, nil
}

func (m *Mapper) URLForPhoto(link string) string {
	parts := strings.Split(link, "/")
	for i, p := range parts {
		parts[i] = url.PathEscape(p)
	}
	return fmt.Sprintf("%s/%s", m.deviriumLink, strings.Join(parts, "/"))
}

var escapeChars = []string{"_", "*", "~", "#", "+", "-", "=", "|", "{", "}", ".", "!", "[", "]", "(", ")"}

func escape(value string) string {
	for _, ch := range escapeChars {
		value = strings.ReplaceAll(value, ch, "\\"+ch)
	}
	return escapeBackticks(value)
}

// escapeBackticks: preserve ``` and paired `, escape single unpaired `.
func escapeBackticks(value string) string {
	const marker = "TRIPLE_BACKTICK_MARKER"
	value = strings.ReplaceAll(value, "```", marker)

	var out strings.Builder
	i := 0
	for i < len(value) {
		if strings.HasPrefix(value[i:], marker) {
			out.WriteString("```")
			i += len(marker)
			continue
		}
		if value[i] == '`' {
			next := strings.IndexByte(value[i+1:], '`')
			if next != -1 {
				end := i + 1 + next + 1
				out.WriteString(value[i:end])
				i = end
			} else {
				out.WriteString("\\`")
				i++
			}
		} else {
			out.WriteByte(value[i])
			i++
		}
	}
	return out.String()
}

func extractMeta(content string, meta map[string]string, patterns ...*regexp.Regexp) string {
	for _, p := range patterns {
		for _, match := range p.FindAllString(content, -1) {
			id := randHex()
			meta[id] = match
			content = strings.ReplaceAll(content, match, id)
		}
	}
	return content
}

func randHex() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
