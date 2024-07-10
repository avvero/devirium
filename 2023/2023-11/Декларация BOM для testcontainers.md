# Декларация BOM для testcontainers

## 1 Вариант

Источник - https://java.testcontainers.org

```groovy
    implementation platform('org.testcontainers:testcontainers-bom:1.19.3') //import bom
    testImplementation('org.testcontainers:mysql') //no version specified
```

## 2 Вариант

Источник - https://github.com/testcontainers/testcontainers-java-spring-boot-quickstart

```groovy
ext {
    set('testcontainersVersion', "1.19.0")
}

dependencies {
    ...
    ...
    testImplementation 'org.springframework.boot:spring-boot-starter-test'
    testImplementation 'org.springframework.boot:spring-boot-testcontainers'
    testImplementation 'org.testcontainers:junit-jupiter'
    testImplementation 'org.testcontainers:postgresql'
    testImplementation 'io.rest-assured:rest-assured'
}

dependencyManagement {
    imports {
        mavenBom "org.testcontainers:testcontainers-bom:${testcontainersVersion}"
    }
}
```


#test-containers #bom #spring/boot/3
#draft