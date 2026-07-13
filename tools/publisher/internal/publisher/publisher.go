package publisher

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/avvero/devirium/tools/publisher/internal/mapper"
	"github.com/avvero/devirium/tools/publisher/internal/resolver"
	"github.com/avvero/devirium/tools/publisher/internal/telegram"
)

// TelegramSender is the subset of telegram.Client the publisher needs.
type TelegramSender interface {
	SendMessage(chatID, text, parseMode string) (telegram.SendMessageResult, error)
	SendPhoto(chatID, photo, caption, parseMode string) (telegram.SendMessageResult, error)
}

// Corrector reviews the note; returns text; if it contains "note is correct" the note publishes.
type Corrector interface {
	Complete(model, prompt string) (string, error)
}

type Config struct {
	DeviriumChatID  string
	GardenerChatID  string
	CorrectorPrompt string
	CorrectorModel  string // usually "gpt-4"
}

type Publisher struct {
	cfg      Config
	tg       TelegramSender
	openai   Corrector
	mapper   *mapper.Mapper
	resolver *resolver.Index
}

func New(cfg Config, tg TelegramSender, ai Corrector, m *mapper.Mapper, idx *resolver.Index) *Publisher {
	return &Publisher{cfg: cfg, tg: tg, openai: ai, mapper: m, resolver: idx}
}

// PublishNote replicates PublicationService.publishNote.
// name = base file name, path = repo-relative path, content = note body.
func (p *Publisher) PublishNote(name, path, content string) error {
	log.Printf("process: %s (path=%s, %d bytes)", name, path, len(content))
	if name == "index.md" {
		log.Printf("skip: %s (index)", name)
		return nil
	}
	if strings.Contains(content, "#draft") || (path != "" && strings.Contains(path, "draft")) {
		log.Printf("skip: %s (draft)", name)
		return nil
	}
	if strings.Contains(content, "#limbo") || (path != "" && strings.Contains(path, "limbo")) {
		log.Printf("skip: %s (limbo)", name)
		return nil
	}
	for _, tag := range []string{"#person", "#book", "#cv", "#aboutme", "#ignore"} {
		if strings.Contains(content, tag) {
			log.Printf("skip: %s (%s)", name, tag)
			return nil
		}
	}

	links, err := p.resolver.ResolveLinks(content)
	if err != nil {
		return p.notifyGardener(name, err)
	}
	images, err := p.resolver.ResolveImages(content)
	if err != nil {
		return p.notifyGardener(name, err)
	}
	hasPhoto := len(images) > 0
	log.Printf("resolved: %s (links=%d, images=%d, hasPhoto=%v)", name, len(links), len(images), hasPhoto)

	body, err := p.mapper.Map(name, path, content, links, hasPhoto)
	if err != nil {
		return p.notifyGardener(name, err)
	}

	log.Printf("corrector: %s (model=%s) requesting review", name, p.cfg.CorrectorModel)
	corrected, err := p.openai.Complete(p.cfg.CorrectorModel, p.cfg.CorrectorPrompt+"\n"+content)
	if err != nil {
		return p.notifyGardener(name, err)
	}
	verdict := "OK"
	if !strings.Contains(strings.ToLower(corrected), "note is correct") {
		verdict = "REVIEW"
	}
	log.Printf("corrector: %s verdict=%s (reply=%q)", name, verdict, truncate(corrected, 120))

	target := p.cfg.DeviriumChatID
	if strings.Contains(body, "#debug") {
		target = p.cfg.GardenerChatID
		log.Printf("route: %s -> gardener (#debug tag)", name)
	}

	if verdict == "REVIEW" {
		log.Printf("route: %s -> gardener for review (chat=%s)", name, p.cfg.GardenerChatID)
		res, err := p.tg.SendMessage(p.cfg.GardenerChatID, body, "MarkdownV2")
		if err != nil {
			return err
		}
		log.Printf("sent: %s draft to gardener (message_id=%d)", name, res.MessageID)
		msg := fmt.Sprintf("Can't process %s: Incorrect text, proposal:\n%s", name, corrected)
		res2, err := p.tg.SendMessage(p.cfg.GardenerChatID, msg, "markdown")
		if err != nil {
			return err
		}
		log.Printf("sent: %s proposal to gardener (message_id=%d)", name, res2.MessageID)
		return nil
	}

	if hasPhoto {
		var firstImage string
		for _, v := range images {
			firstImage = v
			break
		}
		photoURL := p.mapper.URLForPhoto(firstImage)
		log.Printf("publish: %s -> sendPhoto (chat=%s, photo=%s)", name, target, photoURL)
		res, err := p.tg.SendPhoto(target, photoURL, body, "MarkdownV2")
		if err != nil {
			return err
		}
		log.Printf("PUBLISHED: %s to chat=%s message_id=%d (photo)", name, target, res.MessageID)
		return nil
	}
	log.Printf("publish: %s -> sendMessage (chat=%s, %d bytes)", name, target, len(body))
	res, err := p.tg.SendMessage(target, body, "MarkdownV2")
	if err != nil {
		return err
	}
	log.Printf("PUBLISHED: %s to chat=%s message_id=%d", name, target, res.MessageID)
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func (p *Publisher) notifyGardener(name string, cause error) error {
	log.Printf("error processing %s: %v -> notifying gardener (chat=%s)", name, cause, p.cfg.GardenerChatID)
	msg := fmt.Sprintf("Can't process %s: %s", name, cause.Error())
	res, err := p.tg.SendMessage(p.cfg.GardenerChatID, msg, "markdown")
	if err != nil {
		log.Printf("gardener notify failed: %v", err)
		return err
	}
	log.Printf("sent: %s error notice to gardener (message_id=%d)", name, res.MessageID)
	return nil
}

// FileNameFromPath returns basename with .md preserved.
func FileNameFromPath(path string) string {
	return filepath.Base(path)
}
