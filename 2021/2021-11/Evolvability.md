Способность меняться, быть гибким, легким, минимизируют издержки на сопровождение?


Все последующие мысли и утверждения продиктованы личным опытом и субъективным мнением в продуктовой (не заказной) разработки. Сервисориентированность. 

Работаю в старой (возрастной?) компании. Есть продукты, которым более 10 лет. 

Правильнее оценивать возраст в кошачьем и это старички. Есть зрелые пятилетние и совсем юные.

Мое наблюдение таково:

-   я видел как сервис называли и потом переименовывали, потому что изначальное имя было большим, чем реальная функциональность охватывала.
-   я видел как сервисы усложнялось в плане текущей логики и добавления новой и перерастали своё прежнее имя
-   сервисы переставали быть потому что ценность их была незначительна и выражалось в модуль или библиотеку 
-   Итерационная разработка, задача держится разработчиком не дольше неделе, это крайняк, лучше 3 дня
-   Разработку могут вести разные разработчики ?

  

И понимание, которое пришло - нельзя загадывать

  

-   Интеграционные ете тесты ([[Testing trophy]])
-   тдд
-   Простота реализации, просто сказать, не просто сделать
-   Отказ от генерализации на ранних этапах
-   Отказ от попыток сделать что-то с заделом на будущее. Сложно предугадать
-   Графтинг, орфанадж
-   Фича по пакету или слой по пакету?

  

Осторожно:

-   публичный [[api]]

  

Как это выражается в отношении работы с базой: простая структура, возможно с плохой нормализацией, избытком данных

  

  

На ранних этапах проще. Чем сложнее был, тем больше сил на миграцию. В коде проще - есть тесты, можно хоть сколько переписывать код

  

Иногда сложно понять, как лучше построить дизайн. Думаешь, думаешь, а хорошо придумать не получается. Оставляем как есть - пазл не складывается потому что нет нужных кусочков. Вся картина не понятна, не понятно в какую сторону загибать. Можно подождать новых задач вокруг и возможно это поможет. Оставить код в таком виде это ок, главное чтобы была возможность его переписать, в этом помогут тесты.

Важность и пользу тестов очень сложно переоценить. Преимущество ете совместнос с тдд в том, что ты не представляешь как будет выглядеть код внутри, но тебе и не нужно. Скорее всего контракт определён либо интеграционным взаимодействиями или прочими требованиями. И от того так просто писать ете в том числе потому что контракт есть, так же потому что не нужно париться с тем, как будет внутри все работать. При написании тестов даже и думать то особо нечего, просто описываешь требования и все, никаких головняков вокруг внутреннего дизайна. И пока ты пишешь тесты, требования все больше и больше откладываются в голове. Пробежаться это одно, а потратить 30 минут на какой-то вариант поведение всяко больше отложит в уме. Может сложится впечатление, что пока опишешь тесты для системы к один год пройдёт. Но в нашем случае у нас идёт речь о небольших функциональностях.

Поможет закрепить знание о требованиях намного лучше чем повторное прочтение.

  

[[grafting]]. Если нужно разработать нечто большое в составе, например 3 сервисом, то почему бы не начать с 1 сервиса с 3 независимыми модулями? 

На моей практике очень редко, а почти никогда, не удаётся описать апи для новой функциональности, интеграции таким образом, чтобы в процессе разработки оно чуток не поменялось по причине ошибки или с целью улучшения, дающего бенефиты.

Как нам спланировать разработку? Вот представьте сквозную функциональность такого рода (проходит через 3 сервиса). Скорее всего придётся выкатывать доработки последовательно начиная с хвоста. А если что-то оказалось не так в апи? Что-то добавить? Например забыли параметр и нужно протаскивать. Чем дальше, тем сложнее что-то менять.

  

Отказаться от апи, не делать разделения на сервисы, сделайте модуль, потом оторвёте. То на сколько просто вынести модуль в отдельный сервис покажет вообще умение писать слабосвязный код.

  

Нет финального состояние, сервис постоянно меняется

  

[[tdd]]

Сама по себе концепция ничего не значит, важны инструменты благодаря которым не возможно удобно реализовывать: спок, тестконтейнеры, бетамакс


## Переход от мокирования к эмуляции.


#evolvability 