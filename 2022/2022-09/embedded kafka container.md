# Embedded kafka container


File `EmbeddedKafkaApplication`
```java
package pw.avvero.embeddedkafka;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.kafka.test.EmbeddedKafkaBroker;

@Slf4j
@SpringBootApplication
public class EmbeddedKafkaApplication {

    public static void main(String[] args) {
        SpringApplication.run(EmbeddedKafkaApplication.class, args);
    }

    public static final int KAFKA_PORT = 9093;
    public static final int ZK_PORT = 2181;

    public static class It {

    }

    @Bean
    public It embeddedKafkaBroker(@Value("${app.kafka.advertised.listeners}") String advertisedListeners) {
        long start = System.currentTimeMillis();
        log.info("[KT] Kafka from testcontainers is going to start");
        String[] topics = new String[]{"topic1"};

        EmbeddedKafkaBroker broker = new EmbeddedKafkaBroker(1, true, 1, topics)
                .zkPort(ZK_PORT)
                .kafkaPorts(KAFKA_PORT)
                .brokerProperty("listeners", "PLAINTEXT://0.0.0.0:" + KAFKA_PORT + ",BROKER://0.0.0.0:9092")
                .brokerProperty("listener.security.protocol.map", "BROKER:PLAINTEXT,PLAINTEXT:PLAINTEXT")
                .brokerProperty("inter.broker.listener.name", "BROKER")
                .brokerProperty("advertised.listeners", advertisedListeners)
                .zkConnectionTimeout(EmbeddedKafkaBroker.DEFAULT_ZK_CONNECTION_TIMEOUT)
                .zkSessionTimeout(EmbeddedKafkaBroker.DEFAULT_ZK_SESSION_TIMEOUT);

        broker.afterPropertiesSet();
//        System.setProperty("spring.kafka.bootstrap-servers", broker.getBrokersAsString());
        long finish = System.currentTimeMillis() - start;
        log.info("[KT] Kafka from testcontainers is started on: {} (zookeeper: {}, advertised.listeners: {}) in {} millis",
                broker.getBrokersAsString(), broker.getZookeeperConnectionString(), advertisedListeners, finish);
        return new It();
    }

}
```

```properties
app.kafka.advertised.listeners=PLAINTEXT://localhost:9093,BROKER://localhost:9092
```

Dockerfile
```dockerfile
FROM openjdk:17

COPY build/install/embedded-kafka-boot embedded-kafka-boot

RUN ls -al

EXPOSE 55900

ENTRYPOINT ["./embedded-kafka-boot/bin/embedded-kafka"]
```

KafkaEmbeddedContainer
```java
package com.fxclub.test.spock;

import org.testcontainers.containers.GenericContainer;
import org.testcontainers.utility.DockerImageName;

/**
 * This container wraps Confluent Kafka and Zookeeper (optionally)
 */
public class KafkaEmbeddedContainer extends GenericContainer<KafkaEmbeddedContainer> {

    public static final int KAFKA_PORT = 9093;
    public static final int ZOOKEEPER_PORT = 2181;

    public KafkaEmbeddedContainer() {
        super(DockerImageName.parse("embedded-kafka_emk:latest"));
        addFixedExposedPort(KAFKA_PORT, KAFKA_PORT);
        addFixedExposedPort(ZOOKEEPER_PORT, ZOOKEEPER_PORT);
        String host = getNetwork() != null ? getNetworkAliases().get(0) : "localhost";
        withEnv("app.kafka.advertised.listeners", "PLAINTEXT://" + host + ":" + KAFKA_PORT + ",BROKER://localhost:9092");
    }

    public String getBootstrapServers() {
        return String.format("PLAINTEXT://%s:%s", getHost(), getMappedPort(KAFKA_PORT));
    }
}
```

#test-containers #kafka #test
#draft