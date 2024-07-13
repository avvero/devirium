```bash
#!/bin/bash
# Remove all stopped containers
docker rm $(docker ps -a -q)
# Remove all containers
docker rm -f $(docker ps -a -q)
```

#docker #docker_compose #issue #error