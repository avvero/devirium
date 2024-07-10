# Обход блокировки Docker Hub

В связи с тем, что Docker Hub больше не работает (https://t.me/EchelonEyes/2811) в России, пользователи ищут варианты обойти запрет.

Важно: обходные пути чреваты рисками получить бэкдор, поэтому особый смысл приобретает проверка образов с помощью инструмента Cosign (https://edu.chainguard.dev/open-source/sigstore/cosign/an-introduction-to-cosign/) от Google, который предоставляет возможность подписывать и проверять подписи контейнерных образов. Российского аналога Cosign пока нет, и единственная возможность проверять образы – западный сервис Sigstore (https://www.sigstore.dev/).

Также в качестве зеркала работает Yandex:

```
$ docker pull cr.yandex/mirror/alpine

Using default tag: latest
latest: Pulling from mirror/alpine
59bf1c3509f3: Pull complete 
Digest: sha256:e7d88de73db3d3fd9b2d63aa7f447a10fd0220b7cbf39803c803f2af9ba256b3
Status: Downloaded newer image for cr.yandex/mirror/alpine:latest
cr.yandex/mirror/alpine:latest
```

https://huecker.io/

#docker #workaround
#draft