# publisher

CLI that publishes changed `.md` notes to a Telegram channel.

Replaces the old always-on `devirium-bot` Spring service: runs on-demand from
GitHub Actions (or your laptop) using the git commit delta.

## Usage

All commands run from inside `tools/publisher` (this is the Go module root; there
are no `.go` files at the repo root — running `go test` there fails with
`no Go files in .../tools/publisher`).

From the repo root the shortcuts live in the top-level `Makefile`:

- `make publish-dry`   — dry-run against `HEAD~1..HEAD` (prints curl, no network).
- `make publish`       — real publish (requires env vars, see below).
- `make publish-test`  — `go test ./...` inside the module.
- `make publish-check` — tests + `go vet` + build.

Manual invocation:

```
cd tools/publisher

# Dry-run: prints the exact curl command for every OpenAI corrector call and
# every Telegram sendMessage/sendPhoto call. Secrets in headers/URLs are masked.
# No env vars required.
go run ./cmd/publisher --repo "$(git rev-parse --show-toplevel)" --dry-run

# Real publish (env required).
TELEGRAM_TOKEN=... \
DEVIRIUM_CHAT_ID=... \
DEVIRIUM_GARDENER_CHAT_ID=... \
OPENAI_TOKEN=... \
go run ./cmd/publisher --repo "$(git rev-parse --show-toplevel)"
```

Flags: `--repo`, `--base`, `--head`, `--dry-run`, `--devirium-link`,
`--telegram-base`, `--openai-base`, `--corrector-model`, `--corrector-prompt`.

Env: `TELEGRAM_TOKEN`, `DEVIRIUM_CHAT_ID`, `DEVIRIUM_GARDENER_CHAT_ID`,
`OPENAI_TOKEN`, plus optional `DEVIRIUM_LINK`, `TELEGRAM_URI`, `OPENAI_URI`,
`CORRECTOR_MODEL`, `CORRECTOR_PROMPT`.

## Proxy

Every HTTP call (OpenAI + Telegram) honours `HTTPS_PROXY` / `HTTP_PROXY` /
`NO_PROXY` from the environment. In `--dry-run` the emitted curl commands
include the matching `-x <proxy>` flag so the printout is reproducible.

```
HTTPS_PROXY=http://10.0.1.80:8118 \
TELEGRAM_TOKEN=... DEVIRIUM_CHAT_ID=... DEVIRIUM_GARDENER_CHAT_ID=... OPENAI_TOKEN=... \
go run ./cmd/publisher --repo "$(git rev-parse --show-toplevel)"
```

## Layout

- `cmd/publisher` — main
- `internal/gitdelta` — `git diff` / `git show`
- `internal/resolver` — wikilink + image resolution against repo files
- `internal/mapper` — MarkdownV2 escape + wikilink URL replacement
- `internal/telegram` — `sendMessage`, `sendPhoto`
- `internal/openai` — `/v1/chat/completions`
- `internal/publisher` — skip rules, corrector gate, channel/gardener routing

## Tests

Tests must run from the module root with `./...`. Running plain `go test` from
the module root (no packages given) fails because `cmd/publisher/main.go` and
all real code live in subpackages.

```
cd tools/publisher
go test ./...          # all packages
go test ./internal/mapper/...   # single package
go vet ./...
go build ./...
```

## GitHub Actions

`.github/workflows/send_notes.yml` runs `go run ./cmd/publisher` after the
Quartz deploy succeeds. Required secrets: `TELEGRAM_TOKEN`, `DEVIRIUM_CHAT_ID`,
`DEVIRIUM_GARDENER_CHAT_ID`, `OPENAI_TOKEN`, optional `DEVIRIUM_LINK`.

#digital_garden #ignore
