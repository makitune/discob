package command

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

const dwm = "Welcome to "

func (bot *Bot) Welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	lc := bot.LoginChans[p.User.ID]
	if p.Status == discordgo.StatusOnline && lc == nil {
		bot.LoginChans[p.User.ID] = make(chan struct{})
		bot.welcome(s, p)
		bot.headsup(s, p)
	}

	if p.Status == discordgo.StatusOffline && lc != nil {
		lc <- struct{}{}
		delete(bot.LoginChans, p.User.ID)
	}
}

func (bot *Bot) welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	g, err := s.Guild(p.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	c, err := topTextChannel(g)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	wm, err := bot.welcomeMessage()
	if err != nil {
		wm = dwm + g.Name
	}

	msg := p.User.Mention() + "\t" + wm
	sendMessage(s, c, msg)

	wk, err := bot.welcomeKeyword()
	if err != nil {
		return
	}
	bot.sendImage(s, c, wk)
}

func (bot *Bot) headsup(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	g, err := s.Guild(p.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	c, err := topTextChannel(g)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	t := time.NewTicker(1 * time.Hour)
	u := p.User
	lc := bot.LoginChans[u.ID]

	for {
		select {
		case <-t.C:
			msg := u.Mention()
			if m, err := bot.headsUpMessage(); err == nil {
				msg = msg + "\t" + m
			}

			sendMessage(s, c, msg)

			max := len(fpkws)
			bot.sendImage(s, c, fpkws[rand.Intn(max)])
		case <-lc:
			return
		}
	}
}

func topTextChannel(guild *discordgo.Guild) (*discordgo.Channel, error) {
	for _, c := range guild.Channels {
		if c.Type == discordgo.ChannelTypeGuildText && c.Position == 0 {
			return c, nil
		}
	}

	return nil, errors.New("Text Channel not found")
}
