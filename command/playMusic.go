package command

import (
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) PlayMusic(session disgord.Session, evt *disgord.MessageCreate) {
	if bot.voice == nil {
		return
	}

	if !bot.isMentioned(evt) {
		return
	}

	content := evt.Message.Content

	start := strings.Index(content, "<")
	end := strings.Index(content, ">")
	keyword := content[:start] + content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		if e := bot.sendErrorMessage(evt.Ctx, session, evt.Message.ChannelID, err); e != nil {
			errr.Printf("%s\n", e)
		}
		return
	}

	bot.voice.Stop()
	err = bot.voice.Play(y)
	if err != nil {
		errr.Printf("%s\n", err)
		msg := "あかーん"
		if err := bot.sendMessage(evt.Ctx, session, evt.Message.ChannelID, &msg, nil); err != nil {
			errr.Printf("%s\n", err)
			return
		}
	}

	msg := "吟じます！\n" + y.Title
	if err := bot.sendMessage(evt.Ctx, session, evt.Message.ChannelID, &msg, nil); err != nil {
		errr.Printf("%s\n", err)
		return
	}
}
