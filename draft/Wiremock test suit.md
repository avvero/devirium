Каптор описан тут:
+ [[Разносим по полочкам этапы тестирования http запросов в Spring]], [[Ordering Chaos: Arranging HTTP Request Testing in Spring]]
+ [[Повышаем наглядность интеграционных тестов]], [[Enhancing the Visibility of Integration Tests]]

Об использовании в проекте [[Интеграционное тестирование в проекте]]

### Идея для названия 

Stunt double
butt double
back double
mockPerformer

## Введение

Я давно использую wiremock для мокирования внепроцессных зависимостей в тестах. Хотел бы поделиться с вами тем, как я это делаю. Будут разобраны популярные сценарии тестирования и приведены примеры тестов для сценаривев.

Эта статья не является вводной в wiremock, а может быть второй после вводной.

## Инструменты

Spring, spock. На котлине благодаря сахору проще описывать dsl для мокирования.

## Arrange Act Assert и изоляция

Не указываю явно, просто проверяю url.

Необходимо, чтобы приложение перешло в состояние [[Settled state]].

## Приложение / объект тестирования 

Условный телеграм-бот, который перенаправляет запросы к OpenAI API и отправляет ответы пользователям. Контракты взаимодействия с сервисами описаны в упрощенном виде, чтобы подчеркнуть основную логику работы. Ниже приведена диаграмма последовательностей, демонстрирующая архитектуру приложения. Понимаю, что дизайн может вызвать вопросы с точки зрения системной архитектуры, но прошу отнестись к этому с пониманием — главная цель здесь продемонстрировать подход к повышению наглядности в тестах.

## Перехват исходящих запросов и DSL

Для того, чтобы следовать паттерну Arrange Act Assert я использую [объекты специального класса для перехвата исходящих запросов](https://habr.com/ru/articles/781812/). Подобный подход позволяет получить доступ к данным запросов на этапе Assert, где сконцентрирована вся проверка результатов работы в рамках теста.  Я так же использую DSL-обёртки. Это позволяет [скрыть обслуживающий мокирование код и предоставить простой интерфейс для работы со спецификацией](https://habr.com/ru/articles/804673/). Важно подчеркнуть, что предлагается не конкретный DSL, а общий подход, который он реализует. 
И в итоге мой код для мокирования из такого:
```groovy
StubMapping completionsMapping = wireMockServer.stubFor(WireMock.post(urlEqualTo("/v1/chat/completions"))
        .willReturn(aResponse()
                .withBody("""{...}""")
                .withStatus(200)
                .withHeader("Content-Type", "application/json")))
def openaiRequestCaptor = new WiredRequestCaptor(wireMockServer, completionsMapping)
```
превращается в такой
```groovy
def openaiRequestCaptor = restExpectation.openai.completions(withSuccess("""{...}"""))
```

## Тестовые сценарии

Итак перейдем к тестовым сценариям. При написании тестов важно придерживаться изоляции. Тесты должны быть изолированными. Это не означает, что они не могут разделять ресурсы, например общую базу данных. Однако работа одного теста не должна влиять на выполнение другого. В большинстве случаев изоляция может быть достигнута за счет уникальных идентификаторов сущностей (спорное утверждение). Выбор способа изоляции зависит от практической необходимости. Если невозможно достичь изоляции в рамках одного 
> контекста приложения, тесты стоит вынести в отдельный контекст с использованием `@DirtiesContext`.

В качестве тестовых сценариев предлагаются к рассмотрению следующие:
- Запрос-ответ
- Обработка ошибок от openai
- Доступность openai
- Повторы к openai
- Цепочки запросов
- Колбеки

Полный код тестов - pw.avvero.spring.sandbox.bot.wiremock.FeatureWiremockGTestsStep1.

## Ошибки приложения

В случае, если openai на запрос отвечает 404 NOT_FOUND или 403 FORBIDDEN, метод отвечает 500, в телеграм запросов не идет.

```groovy
def "Processing fails if access to openai is not available"() {
    setup:
    def openaiRequestCaptor = restExpectation.openai.completions(openaiResponse)
    def telegramRequestCaptor = restExpectation.telegram.sendMessage(withSuccess("{}"))
    when:
    mockMvc.perform(post("/telegram/webhook")
            .contentType(APPLICATION_JSON_VALUE)
            .content("""...""".toString())
            .accept(APPLICATION_JSON_VALUE))
            .andExpect(status().is(500))
    then:
    openaiRequestCaptor.times == 1
    telegramRequestCaptor.times == 0
    where:
    openaiResponse        | _
    withStatus(NOT_FOUND) | _
    withStatus(FORBIDDEN) | _
}
```
## Ошибки соединения

В случае, если openai на запрос отвечает 404 NOT_FOUND или 403 FORBIDDEN, метод отвечает 500, в телеграм запросов не идет.

```groovy
def "Processing fails if access to openai is not available"() {
    setup:
    def openaiRequestCaptor = restExpectation.openai.completions(openaiResponse)
    def telegramRequestCaptor = restExpectation.telegram.sendMessage(withSuccess("{}"))
    when:
    mockMvc.perform(post("/telegram/webhook")
            .contentType(APPLICATION_JSON_VALUE)
            .content("""...""".toString())
            .accept(APPLICATION_JSON_VALUE))
            .andExpect(status().is(500))
    then:
    openaiRequestCaptor.times == 1
    telegramRequestCaptor.times == 0
    where:
    openaiResponse        | _
    withStatus(NOT_FOUND) | _
    withStatus(FORBIDDEN) | _
}
```

## Delay

Если нужно установить задержку ответа, может быть полезно для тестирование таймаутов для вызовов. Сценарий: если openai не отвечает дольше 1 секунды, то сервис должен вернуть ошибку.

Пример с задержкой:
```kotlin
val methodCaptor = serverMock.method(
    withDelay(Duration.parse(delay), withSuccess("""{}"""))
)
```

## Chain 

Если нужно проверить повторы: если openai отвечает ошибкой 500 и мы делаем повтор. Всего бы делаем 3 повтора и дальше пробрасываем ошибку.

Пример с цепочкой ответов:
```kotlin
val methodCaptor = serverMock.method(
    withConnectionReset(),
    withSuccess("""{}""")
)
```

## Callbacks

Если нужно описать эммитирование ответа от системы, типа уведомление или встречный запрос. 

Как можно проверить, что мы ответили?

## Заключение

#testing #wiremock #article #draft