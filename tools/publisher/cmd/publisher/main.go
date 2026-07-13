package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/avvero/devirium/tools/publisher/internal/gitdelta"
	"github.com/avvero/devirium/tools/publisher/internal/mapper"
	"github.com/avvero/devirium/tools/publisher/internal/openai"
	"github.com/avvero/devirium/tools/publisher/internal/publisher"
	"github.com/avvero/devirium/tools/publisher/internal/resolver"
	"github.com/avvero/devirium/tools/publisher/internal/telegram"
)

const defaultCorrectorPrompt = `Check the note below for grammatical errors. If there are no errors, return the string "Note is correct" without any additional messages. If there are errors, suggest corrections and return only the corrected text of the note. Provide a detailed description of the problems and how you fixed them. Important: do not change the text in [[double square brackets]], as these are links. Also, do not check code snippets. Do not change the language of the note; it must remain in the original language. The notes may be incomplete or structurally imperfect; the task is to check for obvious serious errors, not to nitpick everything. If the note contains only title and tags (#tag) it's considered to be correct. Here is the note:`

func main() {
	var (
		repoRoot     = flag.String("repo", ".", "path to git repo (content root)")
		baseRef      = flag.String("base", "HEAD~1", "base git ref for diff (empty = list all *.md at head)")
		headRef      = flag.String("head", "HEAD", "head git ref for diff")
		dryRun       = flag.Bool("dry-run", false, "log actions instead of calling Telegram/OpenAI")
		deviriumLink = flag.String("devirium-link", envOr("DEVIRIUM_LINK", "https://devirium.com"), "public site base URL")
		tgBase       = flag.String("telegram-base", envOr("TELEGRAM_URI", "https://api.telegram.org"), "telegram API base")
		openaiBase   = flag.String("openai-base", envOr("OPENAI_URI", "https://api.openai.com"), "openai API base")
		correctorModel = flag.String("corrector-model", envOr("CORRECTOR_MODEL", "gpt-4"), "openai model for corrector")
		correctorPrompt = flag.String("corrector-prompt", envOr("CORRECTOR_PROMPT", defaultCorrectorPrompt), "corrector prompt")
	)
	flag.Parse()

	tgToken := os.Getenv("TELEGRAM_TOKEN")
	deviriumChat := os.Getenv("DEVIRIUM_CHAT_ID")
	gardenerChat := os.Getenv("DEVIRIUM_GARDENER_CHAT_ID")
	openaiToken := os.Getenv("OPENAI_TOKEN")

	if !*dryRun {
		for k, v := range map[string]string{
			"TELEGRAM_TOKEN":            tgToken,
			"DEVIRIUM_CHAT_ID":          deviriumChat,
			"DEVIRIUM_GARDENER_CHAT_ID": gardenerChat,
			"OPENAI_TOKEN":              openaiToken,
		} {
			if v == "" {
				log.Fatalf("missing env %s (use --dry-run to skip)", k)
			}
		}
	}
	if *dryRun {
		if deviriumChat == "" {
			deviriumChat = "<DEVIRIUM_CHAT_ID>"
		}
		if gardenerChat == "" {
			gardenerChat = "<DEVIRIUM_GARDENER_CHAT_ID>"
		}
		if tgToken == "" {
			tgToken = "<TELEGRAM_TOKEN>"
		}
	}

	root, err := filepath.Abs(*repoRoot)
	if err != nil {
		log.Fatal(err)
	}

	changed, err := gitdelta.ChangedMarkdown(root, *baseRef, *headRef)
	if err != nil {
		log.Fatalf("git delta: %v", err)
	}
	if len(changed) == 0 {
		log.Printf("no *.md changes between %s..%s", *baseRef, *headRef)
		return
	}
	log.Printf("changed files: %d", len(changed))

	idx, err := resolver.BuildIndex(root)
	if err != nil {
		log.Fatalf("index: %v", err)
	}
	m := mapper.New(*deviriumLink)

	var tg publisher.TelegramSender
	var ai publisher.Corrector
	if *dryRun {
		tg = telegram.NewDryRun(*tgBase, tgToken, os.Stdout)
		ai = openai.NewDryRun(*openaiBase, openaiToken, os.Stdout)
	} else {
		httpc := &http.Client{Timeout: 30 * time.Second}
		tg = telegram.New(*tgBase, tgToken, httpc)
		ai = openai.New(*openaiBase, openaiToken, httpc)
	}

	pub := publisher.New(publisher.Config{
		DeviriumChatID:  deviriumChat,
		GardenerChatID:  gardenerChat,
		CorrectorPrompt: *correctorPrompt,
		CorrectorModel:  *correctorModel,
	}, tg, ai, m, idx)

	failed := 0
	for _, path := range changed {
		body, err := gitdelta.FileAt(root, *headRef, path)
		if err != nil {
			log.Printf("read %s: %v", path, err)
			failed++
			continue
		}
		name := filepath.Base(path)
		if err := pub.PublishNote(name, path, body); err != nil {
			log.Printf("publish %s: %v", path, err)
			failed++
		}
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "%d files failed\n", failed)
		os.Exit(1)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
