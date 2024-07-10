# Semaphore Producer - Consumer

Взято из курса [[Java Multithreading, Concurrency & Performance Optimization]]

Conceptually, a semaphore maintains a set of permits. Each acquire() blocks if necessary until a permit is available, and then takes it.  Each release() adds a permit, potentially releasing a blocking acquirer. 

```java
class Semaphore(int permits)
```

```java
Semaphore full = new Semaphore(0);
Semaphore empty = new Semaphore(1);
Item item = null;

Producer:

while(true) {
    empty.esquire();
    item = getItem();
    full.release();
}

Consumer:

while(true) {
    full.esquire();
    consume(item)
    empty.release();
}
```

#java #snippet #pattern
#draft