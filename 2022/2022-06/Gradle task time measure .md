# Gradle task time measure 

```groovy
import java.util.concurrent.TimeUnit
// Log timings per task.
class TimingsListener implements TaskExecutionListener, BuildListener {
    private long startTime
    private timings = []

    @Override
    void beforeExecute(Task task) {
        startTime = System.nanoTime()
    }

    @Override
    void afterExecute(Task task, TaskState taskState) {
        def ms = TimeUnit.MILLISECONDS.convert(System.nanoTime() - startTime, TimeUnit.NANOSECONDS);
        timings.add([ms, task.path])
        task.project.logger.warn "${task.path} took ${ms}ms"
    }

    @Override
    void buildFinished(BuildResult result) {
        println "Task timings:"
        for (timing in timings) {
            if (timing[0] >= 50) {
                printf "%7sms  %s\n", timing
            }
        }
    }

    @Override
    void projectsEvaluated(Gradle gradle) {}

    @Override
    void projectsLoaded(Gradle gradle) {}

    @Override
    void buildStarted(Gradle gradle) {

    }

    @Override
    void settingsEvaluated(Settings settings) {}
}

gradle.addListener new TimingsListener()
```

#gradle #time