package publisher

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/avvero/devirium/tools/publisher/internal/mapper"
	"github.com/avvero/devirium/tools/publisher/internal/resolver"
	"github.com/avvero/devirium/tools/publisher/internal/telegram"
)

type sentMsg struct{ chat, text, mode string }
type sentPhoto struct{ chat, photo, caption, mode string }

type fakeTG struct {
	messages []sentMsg
	photos   []sentPhoto
	err      error
}

func (f *fakeTG) SendMessage(chat, text, mode string) (telegram.SendMessageResult, error) {
	f.messages = append(f.messages, sentMsg{chat, text, mode})
	return telegram.SendMessageResult{MessageID: 1}, f.err
}
func (f *fakeTG) SendPhoto(chat, photo, caption, mode string) (telegram.SendMessageResult, error) {
	f.photos = append(f.photos, sentPhoto{chat, photo, caption, mode})
	return telegram.SendMessageResult{MessageID: 2}, f.err
}

type fakeAI struct {
	reply string
	err   error
}

func (f *fakeAI) Complete(model, prompt string) (string, error) { return f.reply, f.err }

func newFixture(t *testing.T) (*resolver.Index, *mapper.Mapper, string) {
	t.Helper()
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "2024", "Note One.md"), "body")
	writeFile(t, filepath.Join(root, "pic.png"), "x")
	idx, err := resolver.BuildIndex(root)
	if err != nil {
		t.Fatal(err)
	}
	return idx, mapper.New("https://devirium.com"), root
}

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func baseCfg() Config {
	return Config{
		DeviriumChatID:  "chan",
		GardenerChatID:  "garden",
		CorrectorPrompt: "check:",
		CorrectorModel:  "gpt-4",
	}
}

func TestSkipIndex(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("index.md", "index.md", "hi"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 0 {
		t.Errorf("index published: %v", tg.messages)
	}
}

func TestSkipDraftTag(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "text #draft"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 0 {
		t.Errorf("draft published: %v", tg.messages)
	}
}

func TestSkipDraftPath(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "draft/N.md", "text"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 0 {
		t.Errorf("draft-path published: %v", tg.messages)
	}
}

func TestSkipIgnoreTag(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "text #ignore"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 0 {
		t.Errorf("ignore published")
	}
}

func TestPublishesToChannelWhenCorrect(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "hello"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("want 1 message, got %d", len(tg.messages))
	}
	if tg.messages[0].chat != "chan" {
		t.Errorf("chat=%s", tg.messages[0].chat)
	}
	if !strings.Contains(tg.messages[0].text, "*N*") {
		t.Errorf("text=%s", tg.messages[0].text)
	}
	if tg.messages[0].mode != "MarkdownV2" {
		t.Errorf("mode=%s", tg.messages[0].mode)
	}
}

func TestRoutesToGardenerWhenIncorrect(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Suggested fix: ..."}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "hello"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 2 {
		t.Fatalf("want 2 gardener messages, got %d", len(tg.messages))
	}
	for _, s := range tg.messages {
		if s.chat != "garden" {
			t.Errorf("chat=%s", s.chat)
		}
	}
	if !strings.Contains(tg.messages[1].text, "Incorrect text") {
		t.Errorf("second=%s", tg.messages[1].text)
	}
}

func TestPhotoPathWhenImagePresent(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "hi ![](pic.png)"); err != nil {
		t.Fatal(err)
	}
	if len(tg.photos) != 1 {
		t.Fatalf("want 1 photo, got %d msgs=%v", len(tg.photos), tg.messages)
	}
	if tg.photos[0].photo != "https://devirium.com/pic.png" {
		t.Errorf("photo=%s", tg.photos[0].photo)
	}
}

func TestDebugTagRoutesToGardener(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "hi #debug"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 1 || tg.messages[0].chat != "garden" {
		t.Errorf("expected gardener route, got %v", tg.messages)
	}
}

func TestUnresolvedLinkNotifiesGardener(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{reply: "Note is correct"}, m, idx)
	err := p.PublishNote("N.md", "2024/N.md", "see [[Missing]]")
	if err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 1 || tg.messages[0].chat != "garden" {
		t.Errorf("gardener notif missing: %v", tg.messages)
	}
	if !strings.Contains(tg.messages[0].text, "not found") {
		t.Errorf("text=%s", tg.messages[0].text)
	}
}

func TestOpenAIErrorNotifiesGardener(t *testing.T) {
	idx, m, _ := newFixture(t)
	tg := &fakeTG{}
	p := New(baseCfg(), tg, &fakeAI{err: errors.New("boom")}, m, idx)
	if err := p.PublishNote("N.md", "2024/N.md", "hi"); err != nil {
		t.Fatal(err)
	}
	if len(tg.messages) != 1 || tg.messages[0].chat != "garden" {
		t.Errorf("want gardener notif, got %v", tg.messages)
	}
}

