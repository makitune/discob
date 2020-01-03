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

### Commands

Usage on Discord.

- alola  
  Send a pokemon image known as an Alolan Form when a message contains "あろーら"  (meaning "Alola!").
- discjockey  
  Send a YouTube URL from the keyword following "@bot" when mentioning the bot is NOT joining a voice chat.
- foodporn  
  Send a foodporn message in config.json and a foodporn image when a message contains "疲" or "つかれ" (meaning "Tiredness").
- headsup  
  Send a headsup message in config.json every hour when someone logs in.
- joinVoiceChannel  
  Join top voice channel and send a joinVoiceChannel message in config.json when a message contains "かもーん" (meaning "Come on").
- leaveVoiceChannel  
  Leave voice channel and send a leaveVoiceChannel message in config.json when a message contains "あでゅー" (meaning "Goodbye").
- playMusic  
  Play a music from the keyword following "@bot" when mentioning the bot is joining a voice chat.
- stopMusic  
  Stop a playing music when a message contains "うるさいですね" (meaning "Noisy").
- welcome  
  Send a welcome message and an image from keyword in config.json when someone logs in.
- wikipedia  
  Send a Wikipedia URL from the keyword when a message has the suffix "ってしってる？" (meaning "Do you know?").

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
