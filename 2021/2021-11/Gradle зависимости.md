# [[gradle]] зависимости

Ставил [[java17]] и нужно было обновить gradle.

api и implementation. api вместо compile если нужна транзитивная зависимость и чтобы либы подтянулись.

чтобы api работало в [[gradle7]] нужно добавить плагин

```groovy
plugins {
    id 'java-library'
}
```

#java #gradle