https://www.baeldung.com/spring-tests

```java
@JsonTest: Registers JSON relevant components
@DataJpaTest: Registers JPA beans, including the ORM available
@JdbcTest: Useful for raw JDBC tests, takes care of the data source and in memory DBs without ORM frills
@DataMongoTest: Tries to provide an in-memory mongo testing setup
@WebMvcTest: A mock MVC testing slice without the rest of the app
… (we can check the source to find them all) - https://github.com/spring-projects/spring-boot/tree/master/spring-boot-project/spring-boot-test-autoconfigure/src/main/java/org/springframework/boot/test/autoconfigure
```

#test #spring