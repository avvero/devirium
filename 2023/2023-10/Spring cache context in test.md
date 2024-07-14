```properties
fx.logback.logger.package.org.springframework.test.context.cache=DEBUG
```

```groovy
test {
    jvmArgs "-Dspring.test.context.cache.maxSize=3"
}
```

#spring #test