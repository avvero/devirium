# Another assert json library 

Использую, очень нравится

http://jsonassert.skyscreamer.org/

```java
JSONAssert.assertEquals("""{
    "code": "${data.accountCode}",
    "login": "${data.accountCode}"
}""", response.contentAsString, false) // actual can contain more fields than expected
```

#test #assert #json #java
#draft