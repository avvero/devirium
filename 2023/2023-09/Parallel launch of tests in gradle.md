# Parallel launch of tests in gradle

```groovy
test {
    maxHeapSize '2048m'
    maxParallelForks = 1
}
testlogger {
    theme 'standard-parallel'
    showStandardStreams true
}
```
#draft