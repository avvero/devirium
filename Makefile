# Сохранение изменений
save:
	go run script/update-index.go
	git add .
	git commit -m "update"
	git push

# Создание новой заметки
new:
