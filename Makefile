# Сохранение изменений
save:
	go run script/check-lowercase.go
	go run script/update-index.go
	go run script/update-index-with-tags.go
	go run script/fix-links.go
	git add .
	git commit -m "update"
	git push

# Создание новой заметки
new:
