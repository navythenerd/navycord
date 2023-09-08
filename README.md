# NavyCord
This is just another Discord Bot. I use this Bot for my personal Discord Server for fun and testing purposes.

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
        "dbname": "navycord",
        "user": "navycord",
        "password": "WFPZbVhdDQAtazQc78Rj"
    },
    "web": {
        "inviteService": true,
        "inviteHandle": "/let-me-in",
        "address": "",
        "port": 8000
    }
}
```