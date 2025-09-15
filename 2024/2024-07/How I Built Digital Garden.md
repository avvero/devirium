I built this project using the following components:

- [Personal public repository](https://github.com/avvero/devirium) with notes in Obsidian style (Zettelkasten)
- Branch of script to manage consistency of notes, notes relations, tags, etc. (Quartz4)
- Telegram channel for sharing updates
- [Telegram Bot](https://github.com/avvero/devirium-bot) to post updates in the channel
- [Quartz4](https://quartz.jzhao.xyz/) for static site generation
- GitHub Actions for bot auto-deploy
- GitHub Actions for static site generation on push to the note repository
- GitHub Actions for posting to the Telegram channel on push to the note repository

The design schema is provided below:

```mermaid
sequenceDiagram

Actor me as Me
participant rep as Note repository
participant send_note as Send Note Action
participant bot as Telegram Bot
participant tg as Telegram
participant deploy as Deploy Site Action
participant pages as GitHub Pages

me ->>+ rep: Update notes
rep -->>+ send_note: on push
loop for each note
    send_note ->>+ send_note: Resolve markdown links  
    send_note ->>+ bot: POST /git/webhook
    bot ->>+ bot: Compose message
    bot ->>+ tg: /sendMessage
    tg ->>- bot: ok
    bot ->>- send_note: ok
end

rep -->>+ deploy: on push
deploy ->>+ deploy: Clone Quartz repository<br>Clone content repository<br>Clear Quartz content directory
deploy ->>+ deploy: npx quartz build
deploy ->>+ pages: Deploy
pages ->>- deploy: ok
```

#digital_garden #ignore