```bash
#!/bin/bash

# Ждем появления контейнера
while true; do
    CONTAINER_NAME=$(docker ps --format "{{.Names}}" | grep "postgres")
    if [ -n "$CONTAINER_NAME" ]; then
        echo "$(date) - Найден контейнер PostgreSQL: $CONTAINER_NAME"
        break
    fi
    echo "$(date) - Не вижу PostgreSQL, жду..."
    sleep 1
done

# Получаем порт контейнера
PORT=$(docker ps --format "{{.Names}}\t{{.Ports}}" | grep "$CONTAINER_NAME" | awk -F '[:>-]' '{for (i=1; i<=NF; i++) if ($i ~ /^[0-9]+$/) print $i; exit}')
if [ -z "$PORT" ]; then
    echo "$(date) - Не удалось определить порт PostgreSQL"
    exit 1
fi

echo "$(date) - PostgreSQL слушает на порту $PORT"

# Проверяем соединение каждую секунду
while true; do
    PSQL_OUTPUT=$(PGPASSWORD=postgres psql -U postgres -h localhost -p "$PORT" -t -A -c "SELECT 1;" 2>&1)
    if [[ "$PSQL_OUTPUT" == "1" ]]; then
        echo "$(date) - Соединение успешно"
    else
        echo "$(date) - Ошибка соединения: $PSQL_OUTPUT"
    fi
    sleep 1
done
```

#bash #postgresql