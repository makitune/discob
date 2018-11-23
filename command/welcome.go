package command

import (
	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

const dwm = "Welcome to "

func (bot *Bot) Welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
	if string(p.Status) != "online" {
		return
	}

	g, err := s.Guild(p.GuildID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	for _, c := range g.Channels {
		if c.Type == 0 && c.Position == 0 {
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
	}
}
