package command

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/errr"
)

var (
	joinTrigger        = "かもーん"
	defaultJoinMessage = "Here we go"
)

func (bot *Bot) JoinVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.Contains(m.Content, joinTrigger) {
		return
	}

	if bot.Voice != nil {
		return
	}

	if m.Author.Username == bot.Config.Discord.UserName || m.Author.Bot {
		return
	}

	c, err := s.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	vChan, err := findVoiceChannel(s, c.GuildID)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	connection, err := s.ChannelVoiceJoin(vChan.GuildID, vChan.ID, false, false)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	bot.Voice = &model.Voice{
		Connection: connection,
	}

	msg, err := bot.joinVoiceChannelMessage()
	if err != nil {
		msg = defaultJoinMessage
	}
	sendMessage(s, c, msg)
}

func findVoiceChannel(s *discordgo.Session, guildID string) (*discordgo.Channel, error) {
	st, err := s.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	for _, c := range st {
		if c.Type == discordgo.ChannelTypeGuildVoice {
			return c, nil
		}
	}

	return nil, errors.New("VoiceChannel not found")
}
