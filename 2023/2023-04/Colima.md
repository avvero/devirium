Fix issue https://github.com/testcontainers/testcontainers-java/issues/5034
but some time stops working

During my tests, I couldn't make ryuk work with colima out of the box. I managed to do it by starting the VM with an assigned IP and then overriding TESTCONTAINERS_HOST_OVERRIDE with it.
From that point onwards, together with what is described in the documentation, all my project's tests worked flawlessly.

1. Install Colima: brew install colima
2. Install docker: brew install docker
3. Start Colima with an assigned IP address: colima start --network-address --cpu 1 --memory 1
4. Colima automatically creates and sets a docker context named colima, therefore docker commands on the command line work out of the box
5. To use with testcontainers, make sure the following environment variables are set on your session:
```
export TESTCONTAINERS_HOST_OVERRIDE=$(colima ls -j | jq -r '.address')
export TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE=/var/run/docker.sock
export DOCKER_HOST=unix://$HOME/.colima/default/docker.sock
```

global config
```
launchctl setenv TESTCONTAINERS_HOST_OVERRIDE $(colima ls -j | jq -r '.address')
launchctl setenv TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE /var/run/docker.sock
launchctl setenv DOCKER_HOST unix://${HOME}/.colima/docker.sock
```

colima start --cpu 1 --memory 2 --network-address oraxe

## delete

brew uninstall colima docker docker-compose qemu
brew autoremove
brew cleanup
sudo rm -rf /opt/colima
rm -rf .docker .lima .colima

#colima #docker #mac #test_containers