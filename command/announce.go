package command

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) Announce(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	if bot.announceChans != nil {
		return
	}

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

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	n := time.Now().In(jst)
	weekend := time.Date(n.Year(), n.Month(), n.Day()+(7-int(n.Weekday())), 0, 0, 0, 0, n.Location())

	bot.announceChans = make(chan struct{})
	time.AfterFunc(weekend.Sub(n), func() {
		go announce(time.Hour*24*7, bot.announceChans, func() {
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			t := time.Now().In(jst)
			msg, err := search.SearchGameReleaseSchedule(t, t.AddDate(0, 0, 6))
			if err != nil {
				bot.sendErrorMessage(s, c, err)
			}
			sendMessage(s, c, *msg)
		})
	})
}

func announce(d time.Duration, stopC <-chan struct{}, handler func()) {
	handler()

	t := time.NewTicker(d)
	for {
		select {
		case <-t.C:
			handler()
		case <-stopC:
			return
		}
	}
}
