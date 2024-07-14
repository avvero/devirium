## JMX confluentinc/cp-kafka

https://docs.confluent.io/platform/current/installation/docker/config-reference.html#confluent-ak-configuration   


```bash
docker run -d \
--name=kafka-jmx \
-h kafka-jmx \
-p 9101:9101 \
-e KAFKA_NODE_ID=1 \
-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP='CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT' \
-e KAFKA_ADVERTISED_LISTENERS='PLAINTEXT://kafka-jmx:29092,PLAINTEXT_HOST://localhost:9092' \
-e KAFKA_JMX_PORT=9101 \
-e KAFKA_JMX_HOSTNAME=localhost \
-e KAFKA_PROCESS_ROLES='broker,controller' \
-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
-e KAFKA_CONTROLLER_QUORUM_VOTERS='1@kafka-jmx:29093' \
-e KAFKA_LISTENERS='PLAINTEXT://kafka-jmx:29092,CONTROLLER://kafka-jmx:29093,PLAINTEXT_HOST://0.0.0.0:9092' \
-e KAFKA_INTER_BROKER_LISTENER_NAME='PLAINTEXT' \
-e KAFKA_CONTROLLER_LISTENER_NAMES='CONTROLLER' \
-e CLUSTER_ID='MkU3OEVBNTcwNTJENDM2Qk' \
confluentinc/cp-kafka:7.5.3
```

```bash
docker run -d \
--name=kafka-jmx \
-h kafka-jmx \
-p 9101:9101 \
-p 9092:9092 \
-e KAFKA_AUTO_CREATE_TOPICS_ENABLE=true \
-e KAFKA_NODE_ID=1 \
-e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP='CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT' \
-e KAFKA_ADVERTISED_LISTENERS='PLAINTEXT://kafka-jmx:29092,PLAINTEXT_HOST://localhost:9092' \
-e KAFKA_JMX_PORT=9101 \
-e KAFKA_JMX_HOSTNAME=localhost \
-e KAFKA_PROCESS_ROLES='broker,controller' \
-e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
-e KAFKA_CONTROLLER_QUORUM_VOTERS='1@kafka-jmx:29093' \
-e KAFKA_LISTENERS='PLAINTEXT://kafka-jmx:29092,CONTROLLER://kafka-jmx:29093,PLAINTEXT_HOST://0.0.0.0:9092' \
-e KAFKA_INTER_BROKER_LISTENER_NAME='PLAINTEXT' \
-e KAFKA_CONTROLLER_LISTENER_NAMES='CONTROLLER' \
-e CLUSTER_ID='MkU3OEVBNTcwNTJENDM2Qk' \
confluentinc/cp-kafka:7.5.3
```

## GC EMK

```groovy
bootRun {
    jvmArgs = [
            '-Xloggc:gc.log',
            '-XX:+PrintGCDetails',
    ]
}
```

## GC EMK Native

```bash
./emk-application/build/native/nativeCompile/emk-application -XX:+PrintGC -XX:+VerboseGC
```

#kafka #performance