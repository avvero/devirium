Речь про [[Pact]].

## Pact работает через непрямое тестирование.

Мы указываем серию контрактов. Разделяем их между собой за счет разных параметров запроса. Но это не верно называть контактом, это предопределенные тестовые сценарии, которые работают в рамках контракта системы. Контракт можно разглядеть через эти сценарии, но по факту это демо сценарии с демо данными, разделенные за счет разных параметров запроса. Например если пользователь Иван, то баланс 100. Если Сергей, то баланс 200. Что за Иван и Сергей вообще такие? Бизнес так не описывает требования. Обычно это АС, и там так:
1. Клиент зарегистрировался
2. Клиент пополнился на 100
3. Баланс клиента 100

Вот это описание контракта.

И вывод - можно ли набор моков назвать контрактом? Ну скорее нет. 

Кроме того, подход с матчингом запросов это непрямое тестирование.

Это как раз [[Direct and Indirect Assessment]]
-
#testing #pact