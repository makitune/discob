package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot *Bot) Wikipedia(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasSuffix(m.Content, "ってしってる？") {
		return
	}

	user := m.Author
	if user.Username == bot.config.Discord.UserName || user.Bot {
		return
	}
}
