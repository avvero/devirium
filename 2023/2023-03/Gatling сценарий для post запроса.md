# Gatling сценарий для post запроса

```java
package computerdatabase;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.util.concurrent.ThreadLocalRandom;

import io.gatling.javaapi.core.ScenarioBuilder;
import io.gatling.javaapi.core.Simulation;
import io.gatling.javaapi.http.HttpProtocolBuilder;

import java.time.Duration;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.http;
import static io.gatling.javaapi.http.HttpDsl.status;

public class ComputerDatabaseSimulation extends Simulation {

    HttpProtocolBuilder httpProtocol = http.baseUrl("http://localhost:10031/v3/verificationState")
            .header("Content-Type", "application/json")
            .header("Accept-Encoding", "gzip")
            .check(status().is(200));

    ScenarioBuilder scn = scenario("Root end point calls")
            .exec(http("root end point").post("/").body(StringBody("{\"clientId\":\"19547201\"}")));

    {
        setUp(scn.injectOpen(constantUsersPerSec(100).during(Duration.ofMinutes(15))))
                .protocols(httpProtocol)
                .assertions(global().responseTime().percentile3().lt(100),
                        global().successfulRequests().percent().gt(95.0))
        ;
    }
}
```

#test #testing #loadtesting #gatling
#draft