package command

import (
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/errr"
)

var (
	leaveTrigger        = "あでゅー"
	defaultLeaveMessage = "See ya"
)

func (bot *Bot) LeaveVoiceChannel(session disgord.Session, evt *disgord.MessageCreate) {
	if !strings.Contains(evt.Message.Content, leaveTrigger) {
		return
	}

	if bot.voice == nil {
		return
	}

	bot.voice.Stop()

	if err := bot.voice.Connection.Close(); err != nil {
		bot.voice = nil
		if e := bot.sendErrorMessage(evt.Ctx, session, evt.Message.ChannelID, err); e != nil {
			errr.Printf("%s\n", e)
		}
		return
	}

	bot.voice = nil

	msg, err := bot.leaveVoiceChannelMessage()
	if err != nil {
		msg = defaultLeaveMessage
	}

	if err := bot.sendMessage(evt.Ctx, session, evt.Message.ChannelID, &msg, nil); err != nil {
		errr.Printf("%s\n", err)
		return
	}
}
