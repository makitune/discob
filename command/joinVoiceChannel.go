package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	jointrigger = "かも〜ん"
)

func (bot *Bot) JoinVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.Contains(m.Content, jointrigger) {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}
}
