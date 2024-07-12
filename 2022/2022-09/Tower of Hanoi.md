Хорошее описание решения - https://www.digitalocean.com/community/tutorials/tower-of-hanoi

```java
    private void move(Integer n, Stack<T> source, Stack<T> target, Stack<T> middle) {
        if (n == 1) {
            target.push(source.pop());
            return;
        }
        move(n - 1, source, middle, target);
        target.push(source.pop());
        move(n - 1, middle, target, source);
    }
```

#java #algorithm