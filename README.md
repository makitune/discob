# Discob

Discob is a private bot for Discord using the [Disgord](https://github.com/andersfylling/disgord) library.

## Requirements

- Google APIs
  - Custom Search
  - Youtube Data
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
    "foodporn": {
      "messages": ["cheer up"]
    },
    "headsup": {
      "messages": ["time to eat"]
    },
    "joinVoiceChannel": {
      "messages": ["Here we go"]
    },
    "leaveVoiceChannel": {
      "messages": ["see ya"]
    },
    "welcome": {
      "keywords": ["welcome"],
      "messages": ["welcome back"]
    }
  }
}
```
