# Discob

Discob is a private bot for Discord using the [DiscordGo](https://github.com/bwmarrin/discordgo) library.

## Requirements

- Google APIs
  - Custom Search
  - Youtube Data
- youtube-dl
- ffmpeg

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
    "leaveVoiceChannel": {
      "messages": ["see ya"]
    },
    "foodporn": {
      "messages": ["cheer up"]
    },
    "headsup": {
      "messages": ["time to eat"]
    },
    "joinVoiceChannel": {
      "messages": ["Here we go"]
    },
    "welcome": {
      "keywords": ["welcome"],
      "messages": ["welcome back"]
    }
  }
}
```
