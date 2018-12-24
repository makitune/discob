package command

import (
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) PlayMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voiceConnection == nil {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	if !bot.isMentioned(m) {
		return
	}

	start := strings.Index(m.Content, "<")
	end := strings.Index(m.Content, ">")
	keyword := m.Content[:start] + m.Content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	err = search.DownloadMusic(y, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	if len(y.FilePath) == 0 {
		return
	}

	if bot.stopChan != nil {
		*bot.stopChan <- true
		bot.stopChan = nil
	}
	bot.stopChan = new(chan bool)
	dgvoice.PlayAudioFile(bot.voiceConnection, y.FilePath, *bot.stopChan)
}