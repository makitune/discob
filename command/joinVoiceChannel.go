package command

import (
	"errors"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/errr"
)

var (
	joinTrigger        = "かもーん"
	defaultJoinMessage = "Here we go"
)

func (bot *Bot) JoinVoiceChannel(session disgord.Session, evt *disgord.MessageCreate) {
	if !strings.Contains(evt.Message.Content, joinTrigger) {
		return
	}

	if bot.voice != nil {
		return
	}

	vChan, err := findVoiceChannel(session, evt)
	if err != nil {
		bot.sendErrorMessage(evt.Ctx, session, evt.Message.ChannelID, err)
		return
	}

	connection, err := session.VoiceConnect(vChan.GuildID, vChan.ID)
	if err != nil {
		if e := bot.sendErrorMessage(evt.Ctx, session, evt.Message.ChannelID, err); e != nil {
			errr.Printf("%s\n", e)
		}
		return
	}

	bot.voice = model.New(connection)

	msg, err := bot.joinVoiceChannelMessage()
	if err != nil {
		msg = defaultJoinMessage
	}
	if err := bot.sendMessage(evt.Ctx, session, evt.Message.ChannelID, &msg, nil); err != nil {
		errr.Printf("%s\n", err)
		return
	}

}

func findVoiceChannel(s disgord.Session, evt *disgord.MessageCreate) (*disgord.Channel, error) {
	chs, err := s.GetGuildChannels(evt.Ctx, evt.Message.GuildID)
	if err != nil {
		return nil, err
	}

	for _, ch := range chs {
		if ch.Type == disgord.ChannelTypeGuildVoice {
			return ch, nil
		}
	}

	return nil, errors.New("VoiceChannel not found")
}
