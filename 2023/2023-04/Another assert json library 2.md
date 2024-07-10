# Another assert json library 

Использую, очень нравится

http://jsonassert.skyscreamer.org/

```java
JSONAssert.assertEquals("""{
    "code": "${data.accountCode}",
    "login": "${data.accountCode}"
}""", response.contentAsString, false) // actual can contain more fields than expected
```

Описание механизма сравнения тел из pact [[JSON body matching rules]]. Может пригодиться.

#test #assert #json #java
#draft