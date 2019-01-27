package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot *Bot) Wikipedia(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasSuffix(m.Content, "ってしってる？") {
		return
	}
}
