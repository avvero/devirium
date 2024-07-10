Ссылка [](https://github.com/passy/build-time-tracker-plugin)

```
plugins {
    id "net.rdrei.android.buildtimetracker" version "0.11.0"
}

buildtimetracker {
    reporters {
        csv {
            output "build/times.csv"
            append true
            header false
        }

        summary {
            ordered false
            threshold 50
            barstyle "unicode"
        }

        csvSummary {
            csv "build/times.csv"
        }
    }
}
```

Отчет в таком виде 

```
SUCCESS: Executed 377 tests in 2m 50s

== CSV Build Time Summary ==
Build time today: 3:10.145
Total build time: 3:10.145
(measured since только что)
== Build Time Summary ==
                               0% :clean (0:00.305)
                               1% :payments_api:compileJava (0:02.032)
                               0% :payments_api_…k:compileJava (0:00.376)
                               0% :payments_api_fxbank:jar (0:00.061)
                               2% :compileJava (0:04.455)
                               0% :processResources (0:00.416)
                               1% :compileTestJava (0:02.144)
                            ▇  4% :compileTestGroovy (0:08.137)
                               0% :processTestResources (0:00.087)
▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇▇ 90% :test (2:52.001)

== BUILD SUCCESSFUL ==
```

#gradle 
#draft