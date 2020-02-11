package command

import (
	"errors"
	"math/rand"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

const dwm = "Welcome to "

func (bot *Bot) Welcome(session disgord.Session, evt *disgord.PresenceUpdate) {
	lc := bot.loginChans[evt.User.ID]

	if evt.Status == "online" && lc == nil {
		bot.loginChans[evt.User.ID] = make(chan struct{})
		bot.welcome(session, evt)
		go bot.headsup(session, evt)
	}

	if evt.Status == "offline" && lc != nil {
		lc <- struct{}{}
		delete(bot.loginChans, evt.User.ID)
	}
}

func (bot *Bot) welcome(session disgord.Session, evt *disgord.PresenceUpdate) {
	g, err := session.GetGuild(evt.Ctx, evt.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	ch, err := topTextChannel(g)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	wm, err := bot.welcomeMessage()
	if err != nil {
		wm = dwm + ch.Name
	}

	wk, err := bot.welcomeKeyword()
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}
	msg := evt.User.Mention() + "\t" + wm

	imgURL, err := search.SearchImage(wk, bot.config.Search)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	err = bot.sendMessage(evt.Ctx, session, ch.ID, &msg, imgURL)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}
}

func (bot *Bot) headsup(session disgord.Session, evt *disgord.PresenceUpdate) {
	g, err := session.GetGuild(evt.Ctx, evt.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	ch, err := topTextChannel(g)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	t := time.NewTicker(1 * time.Hour)
	u := evt.User
	lc := bot.loginChans[u.ID]

	for {
		select {
		case <-t.C:
			msg := u.Mention()
			if m, err := bot.headsUpMessage(); err == nil {
				msg = msg + "\t" + m
			}

			imgURL, err := search.SearchImage(fpkws[rand.Intn(len(fpkws))], bot.config.Search)
			if err != nil {
				errr.Printf("%s\n", err)
				return
			}

			err = bot.sendMessage(evt.Ctx, session, ch.ID, &msg, imgURL)
			if err != nil {
				errr.Printf("%s\n", err)
				return
			}
		case <-lc:
			return
		}
	}
}

func topTextChannel(g *disgord.Guild) (*disgord.Channel, error) {
	for _, c := range g.Channels {
		if c.Type == disgord.ChannelTypeGuildText {
			return c, nil
		}
	}

	return nil, errors.New("Text Channel not found")
}
