# Сохранение изменений
save:
	go run script/check-lowercase.go
	go run script/fix-links.go
	go run script/update-index.go
	go run script/update-index-with-tags.go
	git add .
	git commit -m "update"
	git push

find:
	find . -type f | grep -i "$(name)"

# Публикация изменённых заметок в Telegram (dry-run: печатает curl, ничего не шлёт)
publish-dry:
	cd tools/publisher && go run ./cmd/publisher --repo $(CURDIR) --dry-run

# Публикация изменённых заметок в Telegram (реально шлёт). Требуются env: TELEGRAM_TOKEN, DEVIRIUM_CHAT_ID, DEVIRIUM_GARDENER_CHAT_ID, OPENAI_TOKEN
publish:
	cd tools/publisher && go run ./cmd/publisher --repo $(CURDIR)

# Прогон юнит-тестов publisher (запускать из корня модуля с ./...)
publish-test:
	cd tools/publisher && go test ./...

# Полная проверка publisher: тесты + go vet + сборка
publish-check:
	cd tools/publisher && go test ./... && go vet ./... && go build ./...

# Создание новой заметки
new:
