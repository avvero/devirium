# RetryableTopic

org.springframework.kafka.annotation.RetryableTopic

```java
@RetryableTopic(
        dltTopicSuffix = "--dlt",
        retryTopicSuffix = "--retry",
        fixedDelayTopicStrategy = FixedDelayStrategy.SINGLE_TOPIC,
        backoff = @Backoff(3000),
        attempts = "10",
        exclude = {SerializationException.class,
                DeserializationException.class,
                JsonProcessingException.class,
                NumberFormatException.class,
                IllegalArgumentException.class,
                Unrecoverable.class})
```

#kafka #topic #retry
#draft