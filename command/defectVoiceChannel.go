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
	if !strings.Contains(m.Content, defectTrigger) {
		return
	}

	if bot.voice == nil {
		return
	}

	c, err := s.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	if bot.voice.Playing() {
		bot.voice.Stop()
	}

	err = bot.voice.Connection.Disconnect()
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	bot.voice = nil

	msg, err := bot.defectVoiceChannelMessage()
	if err != nil {
		msg = defaultDefectMessage
	}
	sendMessage(s, c, msg)
}
