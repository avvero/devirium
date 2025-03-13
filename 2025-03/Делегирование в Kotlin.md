```java
interface Receiver {
    fun receiveMessage(): String
}

class Lupa : Receiver {
    override fun receiveMessage() = "Сообщение от Лупы!"
}

class Pupa(lupa: Receiver) : Receiver by lupa

fun main() = println(Pupa(Lupa()).receiveMessage()) // "Сообщение от Лупы!"
```

#kotlin #fun