# Тестирование сообщений в Kafka

>Kafka Message Testing
>Kafka Testing: Verifying Message

Данная статья предлагает подход по написанию интеграционных тестов для приложений с Kafka, позволяющий сосредоточиться на спецификации взаимодействия, делая тесты более читаемыми и легкими для поддержки. Представленный подход не только повышает эффективность тестирования, но и способствует лучшему пониманию интеграционных процессов в приложении.

Данная статья опирается на три концепции?, изложенные в соответствующих статьях: [написание тестов с четким разделением на отдельные этапы Arrange-Act-Assert](https://habr.com/ru/articles/781812/), [Изоляция в тестах с Kafka](https://habr.com/ru/articles/797049/) и [использование инструментов повешения наглядности тестов](https://habr.com/ru/articles/804673/). Рекомендую ознакомиться с указанными перед погружением в материал данной.

## Реализация

Предлагаю рассмотреть на примере.

### Демонстрационный сценарий

Возьмем в качестве примера телеграм-бот, который перенаправляет запросы к OpenAI API и возвращает результат пользователю. В случае если запрос в open ai нарушает правила безопасности системы, клиент получит об этом сообщение. Дополнительно будет отправлено сообщение в kafka для системы поведенческого контроля, чтобы товарищ менеджер мог связаться с пользователем и рассказать, что его запрос оказался слишком пикантным даже для нашего робота и попросить его пересмотреть свои предпочтения.

Контракты взаимодействия с сервисами описаны в упрощенном виде, чтобы подчеркнуть основную логику работы. Ниже приведена диаграмма последовательностей, демонстрирующая архитектуру приложения. Понимаю, что дизайн может вызвать вопросы с точки зрения системной архитектуры, но прошу отнестись к этому с пониманием — главная цель здесь продемонстрировать подход к написанию тестов.

```plantuml
@startuml
autonumber

participant "Telegram" as tg
box "Application" #lightgreen
     participant "ServiceA" as ServiceA
     participant "ServiceB" as ServiceB
end box
box "Kafka" #lightblue
participant "topicA" as topicA
participant "topicB" as topicB
participant "topicC" as topicC
end box
participant "OpenAI" as openai

tg -> ServiceA++: request
ServiceA -> topicA: message
ServiceA -> tg--: ok
...
topicA -> ServiceB++: message
ServiceB -> openai++: request
openai -> ServiceB--: response
alt content policy violation error
ServiceB -> topicC++: policy violation event
topicC--
end
ServiceB -> topicB--: message
...
topicB -> ServiceA++: message
ServiceA -> tg--: request

@enduml
```

### Объект захвата сообщений

В качестве основного инструмента тестирования будет использоваться объект захвата сообщений - `RecordCaptor`, его код описан тут - https://github.com/avvero/kafka-test-support/blob/sb3/kafka-support/src/main/java/pw/avvero/test/kafka/RecordCaptor.java.

Захват сообщений будет выполняться через стандартный консумер Kafka. Список топиков нужно указать явно.

```java
@KafkaListener(id = "recordCaptor", topics = "#{'${test.record-captor.topics}'.split(',')}", groupId = "test")
public void eventCaptorListener(ConsumerRecord<Object, Object> record,
                                @Headers Map<String, Object> boundedHeaders) {
    RecordSnapshot recordSnapshot = mapper.recordToSnapshot(record, boundedHeaders);
    recordCaptor.capture(recordSnapshot);
}
```
Образы / snapshots сообщений будут накапливаться в `RecordCaptor`.

Использование подобного подхода требует соблюдение [изоляции в тестах с Kafka](https://habr.com/ru/articles/797049/). Ожидание подтверждения смещения перед проверкой результатов теста необходимо осуществить с помощью метода `pw.avvero.emk.KafkaSupport#waitForPartitionOffsetCommit`.

### Пример теста

Ниже представлен код теста для описанного сценария.

```java
def "User Message Processing with OpenAI"() {
    setup:
    KafkaSupport.waitForPartitionAssignment(applicationContext)                           // 1
    and:                                                                                  // 2
    def openaiRequestCaptor = restExpectation.openai.completions(withBadRequest().contentType(APPLICATION_JSON)
            .body("""{
                "error": {
                "code": "content_policy_violation",
                "message": "Your request was rejected as a result of our safety system."
                }
            }"""))
    def telegramRequestCaptor = restExpectation.telegram.sendMessage(withSuccess('{}', APPLICATION_JSON))
    when:
    mockMvc.perform(post("/telegram/webhook")                                             // 3
            .contentType(APPLICATION_JSON_VALUE)
            .content("""{
                "message": {
                "from": {
                    "id": 10000000
                },
                "chat": {
                    "id": 20000000
                },
                "text": "Hello!"
                }
            }""".toString())
            .accept(APPLICATION_JSON_VALUE))
            .andExpect(status().isOk())
    KafkaSupport.waitForPartitionOffsetCommit(applicationContext)                         // 4
    then:
    openaiRequestCaptor.times == 1                                                        // 5
    JSONAssert.assertEquals("""{
        "content": "Hello!"
    }""", openaiRequestCaptor.bodyString, false)
    and:
    telegramRequestCaptor.times == 1
    JSONAssert.assertEquals("""{
        "chatId": "20000000",
        "text": "Your request was rejected as a result of our safety system."
    }""", telegramRequestCaptor.bodyString, false)
    when:                                                                                 // 6
    def message = recordCaptor.getRecords("topicC", "20000000").last
    then:
    message != null
    JSONAssert.assertEquals("""{
        "webhookMessage": {
        "message": {
            "chat": {
            "id": "20000000"
            },
            "text": "Hello!"
        }
        },
        "error": {
        "code": "content_policy_violation",
        "message": "Your request was rejected as a result of our safety system."
        }
    }""", message.value as String, false)
}
```

Ключевые шаги:
1. Ожидание назначения партиций до начала тестового сценария
2. Мокирование запросов к OpenAI и Telegram
3. Выполнение тестового сценария
4. Ожидание подтверждения смещения
5. Проверки запросов к OpenAI и Telegram  
6. Проверки сообщения в Kafka

Использование `JSONAssert.assertEquals` в данном случае позволяет обеспечить консистентность представления данных на разных уровнях присутствия - в Kafka сообщениях и тестах. Это значительно упрощает процесс тестирования, обеспечивая гибкость сравнения и точность диагностики ошибок.

В статье представлен пример с json форматом сообщений, другие форматы представлений не рассматриваются, но подход описанный не накладывает ограничений.

### Как найти свое сообщение в `RecordCaptor`

Распределение сообщений в `RecordCaptor` осуществляется по названию топика и ключу. В предложенном тесте в качестве ключа используется ключ сообщения в Kafka. При отправке мы явно указываем его:

```java
sendMessage("topicC", chatId, ...);
...
private void sendMessage(String topic, String key, Object payload) {
    Message message = MessageBuilder
            .withPayload(objectMapper.writeValueAsString(payload))
            .setHeader(KafkaHeaders.TOPIC, topic)
            .setHeader(KafkaHeaders.KEY, key)                          <-- указываем ключ
            .build();
    kafkaTemplate.send(message).get();
}
```

Поиск по ключу сообщения в топике:

```java
when:                                                                                
def message = recordCaptor.getRecords("topicC", "20000000").last       <-- используем ключ
```

В случае, если подобный вариант не подходит, необходимо описать ключи для сообщений

### Подключение RecordCaptor

```java
@TestConfiguration(proxyBeanMethods = false)
public class RecordCaptorConfiguration {
    @Bean
    RecordCaptor recordCaptor() {
        return new RecordCaptor();
    }

    @Bean
    RecordCaptorConsumer recordCaptorConsumer(RecordCaptor recordCaptor) {
        return new RecordCaptorConsumer(recordCaptor, new RecordSnapshotMapper());
    }
}
```

## OffsetSnapshotFrame

## Заключение

Ссылка на репозиторий проект с демонстрацией тестов - [sandbox/bot](https://github.com/avvero/kafka-test-support/tree/sb3/example-testcontainers).

Спасибо за внимание к статье, и удачи в вашем стремлении к написанию эффективных и наглядных тестов!

## TODO

- Картинка
- Проверить в [[Главред]]

## Идеи для картинки 

#article #test #kafka #draft