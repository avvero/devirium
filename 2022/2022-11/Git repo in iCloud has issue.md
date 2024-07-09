# Git repo in iCloud has issue

```bash
➜  zet git:(master) git pull
fatal: bad object refs/remotes/origin/master 2
error: https://github.com/avvero/zet.git did not send all necessary objects
```

Нашел решение тут - https://stackoverflow.com/questions/72515916/git-fatal-bad-object-refs-heads-2-master
```
I had a similar problem with a " 2" suffix being added to a filename within the .git directory. The git repository is in a directory synced by iCloud Drive, so presumably iCloud in its infinite wisdom added the suffix during a sync operation.
```

Мне помогло это
```bash
mv .git/refs/heads/master\ 2 .git/refs/heads/master
mv .git/refs/remotes/origin/master\ 2 .git/refs/remotes/origin/master
```

#git #icloud #issue #fix #error
#draft