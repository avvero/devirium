!!! Не работает, testcontainers возвращает ошибку `Unknown DOCKER_HOST scheme ssh, skipping the strategy test`


## Подключиться к удаленной машине

```bash
ssh 10.30.128.223
```

## На удаленной машине

1. Установить Docker. Ниже список команд, которые нужно выполнить последовательно:
```bash
sudo apt update
sudo apt install -y docker.io
sudo systemctl enable docker
sudo systemctl start docker

sudo usermod -aG docker $USER
newgrp docker
```
2. Проверить Docker. Запустить команду:
```bash
docker ps
```
3. (опционально) Запустить dozzle (или другой инструмент) для доступа к логам.
```bash
docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 amir20/dozzle:latest
```

## На локальной машине
1. Проверить доступность удаленного Docker
```bash
docker -H ssh://avvero@10.30.128.223 ps
docker -H ssh://10.30.128.223 ps
```

2. Прописать переменную в нужном месте. Например в `~/.bash_profile`.
```bash
export DOCKER_HOST="ssh://10.30.128.223"
```

## На локальной машине (TLD)
1. Проверить доступность удаленного Docker
```bash
ssh -L 2375:localhost:2375 user@remote-host

docker -H ssh://avvero@10.30.128.223 ps
docker -H ssh://10.30.128.223 ps
```

2. Прописать переменную в нужном месте. Например в `~/.bash_profile`.
```bash
export DOCKER_HOST="ssh://10.30.128.223"
```

3. Проверить Docker. Запустить команду:
```bash
docker ps
```

Если на удаленной машине на 3 шаге вы запускали dozzle, то контейнер будет виден в выводе команды на локальной машине.

#docker