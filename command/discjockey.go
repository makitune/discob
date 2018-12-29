package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) DiskJockey(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !bot.isMentioned(m) {
		return
	}

	if bot.voiceConnection != nil {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	start := strings.Index(m.Content, "<")
	end := strings.Index(m.Content, ">")
	keyword := m.Content[:start] + m.Content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	msg := strings.Join([]string{y.Title, y.Description, y.UrlString()}, "\n")
	sendMessage(s, c, msg)
}
