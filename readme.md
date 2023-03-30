# Fajar Bot

Fajar BOT v1.0.0  
Update Breaking Changes, check `!fb command` (no longer use dots for commands)

> Welcome to join in and feel free to contribute

<!-- > Source [link](https://dev.to/aurelievache/learning-go-by-examples-part-4-create-a-bot-for-discord-in-go-43cf) <!-->

Click this [link](https://discord.com/oauth2/authorize?client_id=1010842038532583456&permissions=8&scope=bot) to add a bot to your server.

## Todo

- [x] Jokes Bapak-Bapak
- [ ] Waktu sholat
- [ ] Al-Qur'an
- [ ] Search Engine
- [ ] Translate
- [ ] Migrate from Heroku to fly.io

## List Commands

- `!fb jokes` = Get one random joke.
- `!fb env` = Check environment.
- `!fb sholat` = COMING SOON!!.
- `!fb search` = Search Engine use google.com.
- `!fb translate-langlist` = List supported language.
- `!fb translate-codelang <language>` = Examine the language code.
- `!fb translate` = COMING SOON!!.
- `!fb ping` = test ping.
- `!fb pong` = test ping.
- `!fb intro` = About this bot.
- `!fb contribute` = Link Repository.

## How to check if bot token is valid
```bash
curl https://discordapp.com/api/v7/gateway/bot -H "Authorization: Bot <token>"
```
