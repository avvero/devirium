https://docs.gitlab.com/ee/ci/variables/predefined_variables.html

```groovy
if (System.getenv("CI") == "true") {
    // Настройки для CI

} else {
    // Настройки для локального билда
}
```

#gradle #build #ci