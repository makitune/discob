package command

import (
	"strings"

	"github.com/makitune/discob/command/search"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) Wikipedia(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasSuffix(m.Content, "ってしってる？") {
		return
	}

	user := m.Author
	if user.Username == bot.Config.Discord.UserName || user.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	i := strings.Index(m.Content, "ってしってる？")
	keyword := m.Content[:i]
	urlString, err := search.SearchWikipediaURL(keyword)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	sendMessage(s, c, "ほれっ\n"+urlString)
}
