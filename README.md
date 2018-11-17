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
        "foodporn": {
            "trigger": [
                "tired"
            ],
            "keywords": [
                "food porn"
            ],
            "messages": [
                "cheer up"
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
