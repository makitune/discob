package command

import (
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) DiskJockey(session disgord.Session, evt *disgord.MessageCreate) {
	if !bot.isMentioned(evt) {
		return
	}

	if bot.voice != nil {
		return
	}

	event := evt.Message
	start := strings.Index(event.Content, "<")
	end := strings.Index(event.Content, ">")
	keyword := event.Content[:start] + event.Content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		if e := bot.sendErrorMessage(evt.Ctx, session, event.ChannelID, err); e != nil {
			errr.Printf("%s\n", err)
		}
		return
	}

	msg := strings.Join([]string{y.Title, y.Description, y.UrlString()}, "\n")
	bot.sendMessage(evt.Ctx, session, event.ChannelID, &msg, nil)
}
