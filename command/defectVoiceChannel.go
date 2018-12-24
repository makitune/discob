package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

var (
	defectTrigger        = "あでゅー"
	defaultDefectMessage = "See ya"
)

func (bot *Bot) DefectVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voiceConnection == nil {
		return
	}

	if !strings.Contains(m.Content, defectTrigger) {
		return
	}

	c, err := s.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	err = bot.voiceConnection.Disconnect()
	if err != nil {
		errr.Printf("%s\n", err)
		bot.sendErrorMessage(s, c, err)
		return
	}

	bot.voiceConnection = nil

	msg, err := bot.defectVoiceChannelMessage()
	if err != nil {
		msg = defaultDefectMessage
	}
	sendMessage(s, c, msg)
}
