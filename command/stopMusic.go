package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

var (
	stopTrigger = "うるさいですね"
)

func (bot *Bot) StopMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.Contains(m.Content, stopTrigger) {
		return
	}

	if !bot.voice.Playing() {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	if err := bot.voice.Stop(); err != nil {
		errr.Printf("%s\n", err)
		return
	}
}
