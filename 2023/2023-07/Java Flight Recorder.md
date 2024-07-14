To enable set in gradle 
```groovy
test {
    jvmArgs '-XX:+FlightRecorder', '-XX:StartFlightRecording=filename=myrecording.jfr'
}
```

```groovy
bootRun {
    jvmArgs '-XX:+FlightRecorder', '-XX:StartFlightRecording=filename=myrecording.jfr'
}
```

#java #jfr #jvm