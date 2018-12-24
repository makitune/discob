package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) DiskJockey(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	if !bot.isMentioned(m) {
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

	if bot.voiceConnection == nil {
		msg := strings.Join([]string{y.Title, y.Description, y.UrlString()}, "\n")
		sendMessage(s, c, msg)
	} else {
		err := search.DownloadMusic(y, bot.config.Search)
		if err != nil {
			bot.sendErrorMessage(s, c, err)
			return
		}

		if len(y.FilePath) == 0 {
			return
		}
	}
}

func (bot *Bot) isMentioned(m *discordgo.MessageCreate) bool {
	if len(m.Mentions) == 0 {
		return false
	}

	for _, mu := range m.Mentions {
		if mu.Username == bot.config.Discord.UserName {
			return true
		}
	}

	return false
}
