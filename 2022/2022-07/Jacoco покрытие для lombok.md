Luckily, JaCoCo honors Lombok’s `@Generated` annotation by ignoring methods annotated with it. We simply have to tell Lombok to add this annotation **by creating a file `lombok.config`** in the main folder of our project with the following content:

```text
lombok.addLombokGeneratedAnnotation = true
```

#jacoco #lambok
#draft