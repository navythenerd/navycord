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
    "twitch": {

    }
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

## Twitch Bot

The Twitch bot implements a basic command system, which can be configured through the `commands.json` file. 

### Default commands

`!dc` or `!discord` is implemented as a default command and uses the `Discord` integration to create a valid invite link and send it in twitch chat.

### commands.json

```
{
    "commands": [
        {
            "trigger": "!ping",
            "response": "pong",
            "permissions": ["broadcaster"]
        },
        {
            "trigger": "!link",
            "response": "Here you can find a link: https://some.link",
            "permissions": []
        }
    ],
    "aliases": [
        {
            "alias": "!foo",
            "trigger": "!link"
        }
    ],
    "timers": [
        {
            "interval": 10,
            "response": "Bar"
        },
        {
            "interval": 20,
            "response": "Foo"
        }
    ]
}
```

### Permissions

The Twitch chat bot supports a basic permission system based upon the Twitch Badge system. Each higher permission inherits the permissions of the lower Badge. The permissions are given by the following badges in ascending order:

 - no badge (everyone can execute command)
 - subscriber (subscriber and higher can execute command)
 - vip (vip and higher can execute command)
 - moderator (moderator and broadcaster can execute command)
 - broadcaster (only broadcaster can execute command)