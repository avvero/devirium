Не должно быть привязки к дизайну, иначе при его изменении все тесты нужно переписать.

Необходимо проверять структуру и состав передаваемых сущностей.

В Тестах мы должны видеть всю картину поведения целиком, а не разрозненными кусками.

## Разбор либы

1 Общие топики на все оригинальные. Это не позволяет изолировать повторы и один забитый топик застопорит обработку всех.
2 spring.kafka.listener.ack-mode=manual почему?
3 батч
4 так не вышло: may contains additional 4th argument with type: ConsumerRecord, if you need raw record (for topic name or message key for example)
5 требует прописать 
```properties
services.esb.sla.rlt-and-dlt.signature.set.enabled=false
services.esb.sla.rlt.signature.verify.enabled=false
services.esb.sla.rlt-and-dlt.signature.public-key=false
services.esb.sla.rlt-and-dlt.signature.private-key=false
```
6 почему то сообщение было разбито при отправке
```groovy
entityBatch = {ArrayList@23671}  size = 20
 1 = ""status":"complete""
 2 = ""created":"Thu Feb 16 10:09:26 UTC 2023""
 4 = ""geography":{"country":"CHN""
 ```

 Payload разбивается на массив строк тут - org.springframework.messaging.handler.invocation.InvocableHandlerMethod#invoke и дальше передается не ConsumerRecord, а список строк.

#test #testing #testing_trophy 