package command

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

const (
	dfk = "food porn"
	dfm = "Take it easy"
)

func (cfg *Config) FoodPorn(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(cfg.Command.FoodPorn.Trigger) == 0 {
		errr.Printf("No configuration")
		return
	}

	for _, trg := range cfg.Command.FoodPorn.Trigger {
		if strings.Contains(m.Content, trg) {
			user := m.Author
			if user.Username == cfg.Discord.UserName || user.Bot {
				return
			}

			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				errr.Printf("%s\n", err)
				return
			}

			msg, err := cfg.foodPornMessage()
			if err != nil {
				msg = dfm
			}
			sendMessage(s, c, msg)

			kw, err := cfg.foodPornKeyword()
			if err != nil {
				kw = dfk
			}
			cfg.sendImage(s, c, kw)
			return
		}
	}
}

func (cfg *Config) HeadsUp(s *discordgo.Session, p *discordgo.PresenceUpdate) {
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

					kw, err := cfg.foodPornKeyword()
					if err != nil {
						kw = dfk
					}
					cfg.sendImage(s, c, kw)
				}
			}
		}
	}()
}
