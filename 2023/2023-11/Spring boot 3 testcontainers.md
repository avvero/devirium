# Spring boot 3 testcontainers

https://spring.io/blog/2023/06/23/improved-testcontainers-support-in-spring-boot-3-1

## Launch locally with testcontainers

```java
@TestConfiguration(proxyBeanMethods = false)
public class ContainersConfiguration {

    @Bean
    @ServiceConnection
    PostgreSQLContainer<?> postgreSQLContainer() {
        return new PostgreSQLContainer<>("postgres:15-alpine");
    }

}

@TestConfiguration(proxyBeanMethods = false)
public class TestSandboxApplication {

	public static void main(String[] args) {
		SpringApplication.from(SandboxApplication::main)
				.with(ContainersConfiguration.class)
				.run(args);
	}

}
```

#spring/boot/3 #test-containers
#draft