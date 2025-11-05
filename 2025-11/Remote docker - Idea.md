Если idea не подхватывает окружение, то можно добавить `DOCKER_HOST=tcp://rd:2375`, где `rd` - это хост, на котором запущен docker remote.

```bash
cat /etc/hosts
##
# Host Database
#
# localhost is used to configure the loopback interface
# when the system is booting.  Do not change this entry.
##
127.0.0.1	localhost
255.255.255.255	broadcasthost
::1             localhost
2a12:5511:31:81c::5   rd
```

#docker #idea