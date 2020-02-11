package command

import (
	"strings"

	"github.com/andersfylling/disgord"
)

var (
	stopTrigger = "うるさいですね"
)

func (bot *Bot) StopMusic(session disgord.Session, evt *disgord.MessageCreate) {
	if !strings.Contains(evt.Message.Content, stopTrigger) {
		return
	}

	bot.voice.Stop()
}
