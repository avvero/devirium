# Cliffhanger

Одна из [[Developer dirty technics]]

Когда от метода отрезается важный для понимания кусок.

```java
public Response method1(Request request) {
    return request
        .attributes()
        .filter()
        .map()
        .forEach(this::updateAttribute)
}

public Value updateAttribute(Attribute attribute) {
    return attribute * 2;
}
```

#development #practice #humor #design #dirty_code #todo
#draft