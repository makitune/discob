package command

import (
	"errors"
	"strings"

	"github.com/makitune/discob/errr"

	"github.com/bwmarrin/discordgo"
)

var (
	joinTrigger        = "かも〜ん"
	defaultJoinMessage = "Here we go"
)

func (bot *Bot) JoinVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voiceConnection != nil {
		return
	}

	if !strings.Contains(m.Content, joinTrigger) {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	c, err := s.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	vChan, err := findVoiceChannel(s, c.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		bot.sendErrorMessage(s, c, err)
		return
	}

	connection, err := s.ChannelVoiceJoin(vChan.GuildID, vChan.ID, false, false)
	if err != nil {
		errr.Printf("%s\n", err)
		bot.sendErrorMessage(s, c, err)
		return
	}

	bot.voiceConnection = connection

	msg, err := bot.joinVoiceChannelMessage()
	if err != nil {
		msg = defaultJoinMessage
	}
	sendMessage(s, c, msg)
}

func findVoiceChannel(s *discordgo.Session, guildID string) (*discordgo.Channel, error) {
	g, err := s.Guild(guildID)
	if err != nil {
		return nil, err
	}

	for _, c := range g.Channels {
		if c.Type == discordgo.ChannelTypeGuildVoice {
			return c, nil
		}
	}

	return nil, errors.New("VoiceChannel not found")
}
