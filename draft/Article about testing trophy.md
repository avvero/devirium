## Сокращение объема регрессионного тестирования.
![alt text](<Article about testing trophy image1.png>)
[[Интеграционное тестирование в проекте]].

## Изолированный интеграционный тест

Если тестируется логика взаимодействия приложения с моком вместо реальных зависимостей, но само приложение работает в условиях, близких к реальному запуску (без изоляции отдельных классов, как в юнит-тестах).
Например, проверка, что сервис правильно формирует HTTP-запросы к базе данных через моки

## Контрактное тестирование

Контрактное тестирование — это методика тестирования точек интеграции, при которой каждое приложение проверяется в изоляции для подтверждения, что сообщения, которые оно отправляет или получает, соответствуют общему пониманию, задокументированному в "контракте". Этот подход позволяет убедиться, что взаимодействие между различными частями системы происходит в соответствии с ожиданиями. [[Повышаем наглядность интеграционных тестов]].

[[Качественна характеристика теста]]
[[Хороший юнит-тест]]
link: Сценарии тестирования в DevP

## Баланс между скоростью и качеством обратной связи у разных видов тестов

[[Интеграционные тесты медленнее unit-тестов?]]

## 

Статья про тестинг трофеи

О чем статья: о том как в разработке сервисов отдаю предпочтение интеграционным тестам уровня сервиса, вижу в этом значительную пользу и с какими трудностями имею дело.

Интеграционная логика и бизнес логика, дать определения и значения

Обзервабилити, тест который медленный, но наглядно показывающий что-то

Тесты важнее кода. Изменения тестов и кода идут в параллельных таймлайнах с нахлестом.

Цель в доставке функциональностей с минимизацией затрат на разработку и поддержку кода.

Задачи:

-   Реализация желаемого бизнес поведения
-   Соблюдение Контрактов интеграционного характера (sla, api)
-   Создания условий, позволяющих минимизировать затраты на поддержку.
-   Рефакторинг (на мой практике гарнизонами планирования функциональностей - квартал, мы рассматриваем инициативы в начале и планируем доставку, мы точно знаем что мы реализуем в ближайший квартал, дальше ничего не знаем, если думаем что знаем - скорее всего будем разочарованы), поэтому невозможно строить долгоиграющий дизайн, тем более если сервис только появился. Важно чтобы сервис был эволвобильным (легко и дёшево меняться) с сохранением контракта. Интеграционные енд - то- енд тесты уровня сервиса помогают значительно.
-   Доступность и актуальность информации о текущей бизнес логике приложения (влияет на оценку времени разработки, в любой момент времени знаем что и как работает - например документация по апи содержит только интеграционную логику)

  

Дисклеймер: 1 Статья для тех, кто пишет код с багами и использует тесты для их нахождения. Или тех, кому интересно?

может содержать баги стремящиеся попасть на прод и кто выбрал тесты, как средство 

что использование тестов сводит к минимуму появление багов в проде 

2 в данной статье я в том числе поделюсь своими мыслями и опытом касательно интеграционных тестов и тдд. Я постараюсь обьяснить, а ни в чем-то убедить. Я не ставлю свой целью убедить , я хочу описать причины.

## Статьи с идеями

https://docs.pact.io/faq/convinceme - тут можно почитать о том, почему Contract tests save development time, суть та же и для TT.

----

Я познакомился давно, решил применять, у нас спринг и это было вызовом - медленно и плавающие баги.

Что такой тестинг трофи 
Дороже чем юнит 
В чем особенность спринга (контекст)
Шаред инстансы баз и Кафки и изоляция или дешевое поднятие сервисов
Что дает
Проблемы - баги спринга

---


Идея для картинки юнит и интеграционные 
Юнит - маленькой линейкой измеряет рост
Интеграции - большой линейкой 

Когда размер имеем значение и глубина

---

Основная мотивация
Перестать пользоваться постманом
Перенести в тесты то что я бы вызвал постманом 

---


Нет ни одной практики программирования, которая бы предложила способ написать такой код или ТАК его написать, который никто не будет менять. Вся затея с software буквально для того, чтобы делать «изменяемые» инструкции. Потому что меняется все. И все практики программирования направлены на то, чтобы писать поддерживаемый код. Т.е нужно писать код, который завтра кто-то перепишет. Так вот, покрытие кода юнитами это стрельнуть в ногу следующему разработчику. Интеграционные - помочь ему.

---

Отлично сочитаются с tdd, пайл техника тоже работает. 
Дарит устойчивость к рефакторингу
[[Моки для проверки внутрисистемных взаимодействий]]
Изоляция

## 

- [The testing pyramid is an outdated economic model](https://www.wiremock.io/post/rethinking-the-testing-pyramid)
- [Testing of Microservices](https://engineering.atspotify.com/2018/01/testing-of-microservices/)
- [On the Diverse And Fantastical Shapes of Testing](https://martinfowler.com/articles/2021-test-shapes.html)

## Комментарии к статье "Пара советов по покрытию тестами проекта на SpringBoot"

[[Комментарии к статье "Пара советов по покрытию тестами проекта на SpringBoot"]]

#testing #testing_trophy #draft