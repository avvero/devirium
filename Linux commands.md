# Linux commands

Получить ответ
```bash
wget -O - http://wiremock:8080/getLifecycleState
```

Запустить, выполнить, убить
```bash
docker run --rm alpine:3.17 wget -O - https://ya.ru
```


wget --server-response -O - http://wiremock:8080/getLifecycleState 2>&1 | grep 'HTTP/' 
wget --server-response -O - http://wiremock:8080/getLifecycleState
#draft