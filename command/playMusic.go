package command

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) PlayMusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if bot.voice == nil {
		return
	}

	if !bot.isMentioned(m) {
		return
	}

	if m.Author.Username == bot.config.Discord.UserName || m.Author.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	start := strings.Index(m.Content, "<")
	end := strings.Index(m.Content, ">")
	keyword := m.Content[:start] + m.Content[end+1:]
	y, err := search.SearchYoutube(keyword, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	err = search.DownloadMusic(y, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	if y.FilePath == nil {
		return
	}

	if bot.voice.Playing() {
		if err := bot.voice.Stop(); err != nil {
			errr.Printf("%s\n", err)
			return
		}
	}

	msg := "吟じます！\n" + y.Title
	sendMessage(s, c, msg)

	bot.voice.Youtube = y
	if err = bot.voice.Play(); err != nil {
		errr.Printf("%s\n", err)
	}

	if y.FilePath == nil {
		return
	}

	if _, err = os.Stat(*y.FilePath); !os.IsNotExist(err) {
		if err := os.Remove(*y.FilePath); err != nil {
			errr.Printf("%s\n", err)
		}
	}
}