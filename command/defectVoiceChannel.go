package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
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

	err = bot.voiceConnection.Disconnect()
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	bot.voiceConnection = nil
}
