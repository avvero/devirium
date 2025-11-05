Cloud-init для [[Docker remote]].

Dozzle для доступа к логам - http://host:8080/

```yaml
#cloud-config
package_update: true
package_upgrade: false

packages:
  - docker.io

groups:
  - docker

system_info:
  default_user:
    groups: [docker]

write_files:
  - path: /etc/systemd/system/docker.service.d/override.conf
    content: |
      [Service]
      ExecStart=
      ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock -H tcp://0.0.0.0:2375

runcmd:
  - systemctl daemon-reload
  - systemctl enable docker
  - systemctl restart docker
  - |
    docker run -d \
      -v /var/run/docker.sock:/var/run/docker.sock \
      -p 8080:8080 \
      amir20/dozzle:latest

```

#docker