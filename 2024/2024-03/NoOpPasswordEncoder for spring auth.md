# NoOpPasswordEncoder for spring auth

If one doesn't need to encode basic auth password

```java
@Bean
public PasswordEncoder passwordEncoder() {
    return NoOpPasswordEncoder.getInstance();
}
```

#spring #security #fix
#draft