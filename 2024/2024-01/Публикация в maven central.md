## Рабочая


1 Регаем https://central.sonatype.com/publishing/namespaces и берем креды

```
ossrhUsername=
ossrhPassword=
```

2 Делаем pom - https://central.sonatype.org/publish/publish-portal-maven/#adding-a-webhook

https://bitbucket.org/simpligility/ossrh-pipeline-demo/src/master/

## Issue "gpg: signing failed: Inappropriate ioctl for device" on MacOS with Maven

https://stackoverflow.com/questions/57591432/gpg-signing-failed-inappropriate-ioctl-for-device-on-macos-with-maven

```bash
GPG_TTY=$(tty)
export GPG_TTY
```
to my ~/.bash_profile file. Now it is working.

## Херня

https://medium.com/mobileaction-tech/publish-private-java-library-to-github-packages-2fbee03deed2
https://www.jetbrains.com/help/idea/add-a-gradle-library-to-the-maven-repository.html#publish_remote

https://central.sonatype.com/publishing/namespaces

https://habr.com/ru/articles/649999/

https://central.sonatype.org/publish/requirements/gpg/#deploying-to-ossrh-with-gradle-introduction

```
ossrhUsername=
ossrhPassword=
```

Could not PUT 'https://s01.oss.sonatype.org/service/local/staging/deploy/maven2/'. Received status code 401 from server: Unauthorized

## gpg

name = 'Anton Belyaev'
email = ''
password = 5ss6h@


```bash
gpg --list-keys --keyid-format short
```

Отправка
```bash
gpg --keyserver keys.openpgp.org --send-keys 
```
Проверка
```bash
gpg --keyserver keys.openpgp.org --search-keys 
```

#java #library #maven