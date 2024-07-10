# docker-compose INTERNAL ERROR: cannot create temporary directory

```bash
#!/bin/bash
# Remove all stopped containers
docker rm $(docker ps -a -q)
# Remove all containers
docker rm -f $(docker ps -a -q)
```

#docker #docker-compose #issue #error
#draft