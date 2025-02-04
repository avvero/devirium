Wiremock chain definition builder POC

```java
fun WireMockServer.issueServiceAgentToken(responseDefinitionBuilder: ResponseDefinitionBuilder?): WiredRequestCaptor {
    val scenarioName = "IssueServiceAgentTokenScenario"

    // Первый ответ — CONFLICT
    val completionsMapping: StubMapping = this
        .stubFor(
            WireMock.get(WireMock.urlEqualTo("/tokens"))
                .inScenario(scenarioName)
                .whenScenarioStateIs(STARTED)
                .willReturn(WireMock.aResponse()
                    .withStatus(HttpStatus.INTERNAL_SERVER_ERROR.value())
                )
                .willSetStateTo("AfterConflict")
        )

    // Второй ответ — успешный
    val completionsMapping2: StubMapping = this
        .stubFor(
            WireMock.get(WireMock.urlEqualTo("/tokens"))
                .inScenario(scenarioName)
                .whenScenarioStateIs("AfterConflict")
                .willReturn(WireMock.aResponse()
                    .withStatus(HttpStatus.OK.value())
                    .withBody("""{"accessToken": "token-from-mock"}""")
                )
        )

    return WiredRequestCaptor(this, completionsMapping)
}

```

#poc #draft #wiremock