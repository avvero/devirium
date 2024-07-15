```java
@Bean
MockWebServer growthBookMockWebServer() throws IOException {
    MockWebServer server = new MockWebServer();
    server.start();
    server.enqueue(new MockResponse().addHeader("Content-Type", "application/json").setBody(fromFile("growthbook/features.json")));
    return server;
}

@Bean
@Primary
public GrowthBookProviderConfigProperties growthBookProviderConfigPropertiesTest(MockWebServer mockWebServer) {
    GrowthBookProviderConfigProperties properties = new GrowthBookProviderConfigProperties();
    properties.setApiHost(mockWebServer.url("/").toString());
    properties.setClientKey("sdk-hGgov5ik7L90MEZO");
    properties.setRefreshStrategy("STALE_WHILE_REVALIDATE");
    properties.setSwrTtlSeconds(10);
    return properties;
}
```

#mock #http