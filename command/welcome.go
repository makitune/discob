package command

import (
	"discord-bot/errr"

	"github.com/bwmarrin/discordgo"
)

const dwm = "Welcome to "

func (cfg *Config) Welcome(s *discordgo.Session, p *discordgo.PresenceUpdate) {
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
			wm, err := cfg.welcomeMessage()
			if err != nil {
				wm = dwm + g.Name
			}

			msg := p.User.Mention() + "\t" + wm
			sendMessage(s, c, msg)

			wk, err := cfg.WelcomeKeyword()
			if err != nil {
				return
			}
			cfg.sendImage(s, c, wk)
		}
	}
}
