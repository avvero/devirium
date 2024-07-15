Use property POSTGRES_HOST_AUTH_METHOD - https://hub.docker.com/_/postgres

```java
    private final static PostgreSQLContainer<?> POSTGRES = new PostgreSQLContainer<>(DockerImageName.parse("public.ecr.aws/docker/library/postgres:14").asCompatibleSubstituteFor("postgres"))
            .withDatabaseName("application")
            .withEnv("POSTGRES_HOST_AUTH_METHOD", "trust")
            .withCommand("postgres -c max_connections=300")
            .waitingFor(Wait.forListeningPort());
```

В стектрейсе было:
```java
java.lang.invoke.VarHandleByteArrayAsInts$ArrayHandle.index
org.postgresql.shaded.com.ongres.scram.common.util.CryptoUtil.hi
org.postgresql.shaded.com.ongres.scram.common.ScramMechanisms.saltedPassword
```

#test_containers #postgresql #authentication