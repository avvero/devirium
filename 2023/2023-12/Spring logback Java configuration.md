From https://logback.qos.ch/manual/configuration.html
> Given that Groovy is a full-fledged language, we have dropped support for logback.groovy in order to protect the innocent.

Logback 1.2.9 dropped the support for Groovy

LogBack java configuration for spring boot 3.

Repo: [spring-logback-java-configuration](https://github.com/avvero/spring-logback-java-configuration)


```java
Logging system failed to initialize using configuration from 'null'
java.lang.IllegalStateException: Could not initialize Logback logging from classpath:logback.groovy
	at org.springframework.boot.logging.logback.LogbackLoggingSystem.lambda$loadConfiguration$1(LogbackLoggingSystem.java:252)
	at org.springframework.boot.logging.logback.LogbackLoggingSystem.withLoggingSuppressed(LogbackLoggingSystem.java:467)
	at org.springframework.boot.logging.logback.LogbackLoggingSystem.loadConfiguration(LogbackLoggingSystem.java:244)
```

Заставим спринг работать 

Почему не попросим? Не получается 

#article #spring #logging #draft