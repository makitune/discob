package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) PlayMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voice == nil {
		return
	}

	if !bot.isMentioned(m) {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	start := strings.Index(m.Content, "<")
	end := strings.Index(m.Content, ">")
	keyword := m.Content[:start] + m.Content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	err = search.DownloadMusic(y, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	if len(y.FilePath) == 0 {
		return
	}

	if bot.voice.Playing() {
		if err := bot.voice.Stop(); err != nil {
			errr.Printf("%s\n", err)
			return
		}
	}

	bot.voice.Youtube = y
	bot.voice.Play()
}
