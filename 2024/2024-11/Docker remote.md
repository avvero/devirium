## Дисклеймер

Данная инструкция предполагает использование TCP-соединения для удаленного управления Docker. Убедитесь, что настройка безопасна в вашем окружении, так как открытие TCP-порта Docker без дополнительной защиты (например, TLS) может создать риски для безопасности.

## Подключиться к удаленной машине

```bash
ssh <REMOTE_IP>
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

4. Открыть файл `sudo nano /lib/systemd/system/docker.service` и изменить настройку `ExecStart`.
```bash
ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375
```

5. Перезапустить Docker.
```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```

6. Проверить состояние службы Docker:
```bash
sudo systemctl status docker
```

7. Проверить, что Docker порт используется:
```bash
sudo ss -tuln | grep 2375
```

## На локальной машине

1. Проверить доступность удаленного Docker:
```bash
docker -H tcp://<REMOTE_IP>:2375 ps
```

2. Прописать переменную в нужном месте. Например в `~/.bash_profile`.
```bash
export DOCKER_HOST="tcp://<REMOTE_IP>:2375"
```

#### Alias
```bash
alias dockerx="docker -H=your-remote-server.org:2375"
```

#### For ssh
```bash
export DOCKER_HOST="ssh://username@your-remote-server.org"
```

#docker