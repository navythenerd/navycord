# NerdGuardian
This is a general purpose bot for streamers. This bot combines twitch, discord and web facilities in one monolithic bot application.
The bot is geared towards advanced users who are willing to host their own bot. This repository also builds a good foundation to build your own custom bot application.

## Configuration

In the current state the main configuration of this bot is done through `config.json`. A `config.json` could look like the following and needs to be updated with your server's ids.

### config.json

```
{
    "discord": {
        "appId": "123456789",
        "token": "someVeryLongBotToken",
        "guildId": "123456789",
        "logChannelId": "123456789",
        "inviteChannelId": "123456789",
        "rulesMessageId": "123456789",
        "verifiedRoleId": "123456789",
        "rulesChannelId": "123456789",
        "agreeRulesEmoteReaction": "✅",
        "rules": "./assets/rules.md"
    },
    "storage": {
        "host": "db",
        "port": 5432,
        "dbname": "nerdguardian",
        "user": "nerdguardian",
        "password": "somePassword"
    },
    "web": {
        "inviteService": true,
        "inviteHandle": "/let-me-in",
        "address": "",
        "port": 8000
    }
}
```