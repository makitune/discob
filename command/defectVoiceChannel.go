package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	defectTrigger = "あでゅー"
)

func (bot *Bot) DefectVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voiceConnection == nil {
		return
	}

	if !strings.Contains(m.Content, defectTrigger) {
		return
	}
}
