# vscode rest client

Дока: https://github.com/Huachao/vscode-restclient

Расширение для vscode для вызова http прямо из текста (для запроса нужно выделить текст)
```http
POST https://example.com/comments HTTP/1.1
Content-Type: application/json

{
    "name": "sample",
    "time": "Wed, 21 Oct 2015 18:27:50 GMT"
}
```

можно так, но для запроса нужно выделить текст
POST http://kyc-profile-service-fxb2-priv.qa-env.com/v3/getProfile
Content-Type: application/json
Authorization: token xxx

{
    "clientId":"100118793"
}

или из файла с расширением !example.http

Есть поддержка авторизации

GET https://httpbin.org/basic-auth/user/passwd HTTP/1.1
Authorization: Basic user:passwd

#http #vscode #rest