# Logic type tangle

Одна из [[Developer dirty technics.md]]

Этот метод вытянул из себя кусок бизнеслогики, я бы отделял бизнес логикику и интеграцию, например
вместо
```java
    private static boolean processForCorrespondingRequestedDocument(VerificationProcessSumSub process, RequestedDocumentEntity requestedDocument) {
        return StringUtils.equals(process.getRequestId(), Objects.toString(requestedDocument.getId()));
    }
    ... 
    !processForCorrespondingRequestedDocument(process, requestedDocument)
```
написал так
```java
    private static boolean equals(Object a, Object b) {
        return StringUtils.equals(Objects.toString(a), Objects.toString(b));
    }
    ... 
    !equals(process.getRequestId(), requestedDocument.getId())
```


#development #practice #fun #design #dirty_code
#draft