package command

import (
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) Wikipedia(session disgord.Session, evt *disgord.MessageCreate) {
	content := evt.Message.Content
	if !strings.HasSuffix(content, "ってしってる？") {
		return
	}

	i := strings.Index(content, "ってしってる？")
	keyword := content[:i]
	u, err := search.SearchWikipediaURL(keyword)
	if err != nil {
		if e := bot.sendErrorMessage(evt.Ctx, session, evt.Message.ChannelID, err); e != nil {
			errr.Printf("%s\n", err)
		}
		return
	}

	msg := "ほれっ " + u
	err = bot.sendMessage(evt.Ctx, session, evt.Message.ChannelID, &msg, nil)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}
}
