Бинарника есть, сборка крива, но можно в docker.

https://foxglove.dev/blog/installing-ros2-on-macos-with-docker

Сборка
```Dockerfile
FROM ros:lyrical-ros-core
RUN apt-get update && apt-get install -y \
ros-lyrical-foxglove-bridge \
ros-lyrical-tf2-ros \
ros-lyrical-ros-base \
ros-lyrical-desktop
```

```bash
docker pull ros:lyrical-ros-core
docker build -t rosco .
```

Проверка с [[foxglove]]
```bash
docker run --rm -it -p 8765:8765 rosco ros2 launch foxglove_bridge foxglove_bridge_launch.xml
```
```bash
docker run --rm -it rosco ros2 run tf2_ros static_transform_publisher --x 0 --y 1 --z 1 --roll 0 --pitch 0 --yaw 0 --frame-id base_link --child-frame-id sensor
```

#ros2 #robot