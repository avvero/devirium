# Spring config injection in string

```java
	@RetryableTopic(attempts = "#{__this.config.retryAttempts}",
			backoff = @Backoff(delayExpression = "#{__this.config.retryDelay.toMillis()}"),
			retryTopicSuffix = "#{__this.config.imqTopicSuffix}",
			dltTopicSuffix = "#{__this.config.dlqTopicSuffix}",
			timeout = "#{__this.config.retryTimeout != null ? __this.config.retryTimeout.toMillis() : ''}",
			include = RetryableException.class,
			fixedDelayTopicStrategy = FixedDelayStrategy.SINGLE_TOPIC
	)
```

#java #spring #configuration
#draft