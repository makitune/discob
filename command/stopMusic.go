package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	stopTrigger = "うるさいですね"
)

func (bot *Bot) StopMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.Contains(m.Content, stopTrigger) {
		return
	}

	if bot.stopChan == nil {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	close(bot.stopChan)
	bot.stopChan = nil
}
