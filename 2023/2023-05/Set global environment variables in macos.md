# Set global environment variables in macos

Нашел рецепт тут - https://emmanuelbernard.com/blog/2012/05/09/setting-global-variables-intellij/

```
launchctl setenv TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE ${HOME}/.colima/docker.sock
launchctl setenv DOCKER_HOST unix:///${HOME}/.colima/docker.sock
launchctl setenv TESTCONTAINERS_RYUK_DISABLED true
```

#env #mac
#draft