# Discob

Discob is a bot for Discord using the [DiscordGo](https://github.com/bwmarrin/discordgo) library.

## Requirements

- Google Custom Search API

## Usage

```console
discob -path config.json
```

## Configuration

Create `config.json` file and specifies it at execute using this template:

```json
{
    "discord": {
        "username": "Bot Name",
        "token": "Bot Token"
    },
    "cse": {
        "id": "Search engine ID",
        "key": "API key"
    },
    "command": {
        "errormessage": "something bad happened",
        "foodporn": {
            "messages": [
                "cheer up"
            ]
        },
        "headsup": {
			"messages": [
				"time to eat"
			]
		},
        "welcome": {
            "keywords": [
                "welcome"
            ],
            "messages": [
                "welcome back"
            ]
        }
    }
}
```
