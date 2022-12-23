package command

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

var greetings = []string{
	"おはこんハロチャオ",
	"ナンジャモ",
	"Iono",
}

func (bot *Bot) Paldea(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "おはこんハロチャオ" {
		return
	}

	user := m.Author
	if user.Username == bot.Config.Discord.UserName || user.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	bot.sendImage(s, c, greetings[rand.Intn(len(greetings))])
}
