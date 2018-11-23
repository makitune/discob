package command

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) HeadsUp(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	if string(p.Status) != "online" {
		return
	}

	go func() {
		u := p.User
		g, err := s.Guild(p.GuildID)
		if err != nil {
			errr.Printf("%s\n", err)
			return
		}

		var c *discordgo.Channel
		for _, ch := range g.Channels {
			if ch.Type == 0 && ch.Position == 0 {
				c = ch
			}
		}

		if c == nil {
			return
		}

		t := time.NewTicker(1 * time.Hour)
		for {
			<-t.C
			for _, p := range g.Presences {
				id := p.User.ID
				if id == u.ID {
					if string(p.Status) != "online" {
						t.Stop()
						return
					}

					msg := u.Mention() + "、あなた疲れてるのよ\n"
					sendMessage(s, c, msg)

					max := len(fpkws)
					bot.sendImage(s, c, fpkws[rand.Intn(max)])
				}
			}
		}
	}()
}
