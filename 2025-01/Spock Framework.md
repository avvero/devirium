```groovy
def "numbers to the power of two"(int a, int b, int c) {
    expect:
    Math.pow(a, b) == c

    where:
    a | b | c
    1 | 2 | 1
    2 | 2 | 4
    3 | 2 | 9
}
```

#testing #tool