package command

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

const dwm = "Welcome to "

func (bot *Bot) Welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	lc := bot.loginChans[p.User.ID]
	if p.Status == discordgo.StatusOnline && lc == nil {
		lc = make(chan struct{})
		bot.welcome(s, p)
		bot.headsup(s, p)
	}

	if p.Status == discordgo.StatusOffline && lc != nil {
		delete(bot.loginChans, p.User.ID)
	}
}

func (bot *Bot) welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	g, err := s.Guild(p.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	for _, c := range g.Channels {
		if c.Type == discordgo.ChannelTypeGuildText && c.Position == 0 {
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
			break
		}
	}
}

func (bot *Bot) headsup(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	g, err := s.Guild(p.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	var c *discordgo.Channel
	for _, ch := range g.Channels {
		if ch.Type == discordgo.ChannelTypeGuildText && ch.Position == 0 {
			c = ch
			break
		}
	}

	if c == nil {
		return
	}

	t := time.NewTicker(1 * time.Hour)
	u := p.User
	for {
		<-t.C
		for _, p := range g.Presences {
			id := p.User.ID
			if id == u.ID {
				if p.Status != discordgo.StatusOnline {
					t.Stop()
					return
				}

				msg := u.Mention()
				if m, err := bot.headsUpMessage(); err == nil {
					msg = msg + "\t" + m
				}

				sendMessage(s, c, msg)

				max := len(fpkws)
				bot.sendImage(s, c, fpkws[rand.Intn(max)])
			}
		}
	}
}
