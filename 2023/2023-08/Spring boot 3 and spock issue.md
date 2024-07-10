# Spring boot 3 and spock issue

Reverences: 
- https://github.com/spockframework/spock/issues/1748
- https://github.com/spockframework/spock/issues/1539

Такие зависимости работают верно
```
plugins {
    id 'java'
    id 'groovy'
    id 'application'
    id 'org.springframework.boot' version '3.1.2'
    id 'io.spring.dependency-management' version '1.1.2'
    id 'org.graalvm.buildtools.native' version '0.9.23'
}


dependencies {
    compileOnly 'org.projectlombok:lombok:1.18.24'
    annotationProcessor 'org.projectlombok:lombok:1.18.24'

    implementation 'org.springframework.boot:spring-boot-starter-web'

    testCompileOnly 'org.projectlombok:lombok:1.18.24'
    testAnnotationProcessor 'org.projectlombok:lombok:1.18.24'

    testImplementation 'org.springframework.boot:spring-boot-starter-test'
    testImplementation platform('org.apache.groovy:groovy-bom:4.0.5')
    testImplementation 'org.apache.groovy:groovy'
    testImplementation 'org.apache.groovy:groovy-json'
    testImplementation platform("org.spockframework:spock-bom:2.4-M1-groovy-4.0")
    testImplementation "org.spockframework:spock-core"
    testImplementation "org.spockframework:spock-spring"

    testImplementation "com.github.tomakehurst:wiremock-jre8-standalone:2.35.0"
    testImplementation "org.assertj:assertj-core:3.11.1"
    testImplementation 'org.awaitility:awaitility:4.0.2'
}

test {
    useJUnitPlatform()
}
```

#spock #spring #java
#draft