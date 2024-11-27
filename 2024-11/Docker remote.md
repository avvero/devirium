```bash
sudo apt update
sudo apt install -y docker.io
sudo systemctl enable docker
sudo systemctl start docker

sudo usermod -aG docker $USER
newgrp docker
docker ps

docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8080:8080 amir20/dozzle:latest
```


## On remote server
```bash
sudo nano /lib/systemd/system/docker.service

ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375

sudo systemctl daemon-reload
sudo systemctl restart docker
```

To Check:
```bash
sudo systemctl status docker
sudo ss -tuln | grep 2375
```

## On local
To check:
```bash
docker -H tcp://10.30.128.212:2375 ps
```
To configure:
```bash
export DOCKER_HOST="tcp://10.30.128.212:2375"
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