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
#gradle #test