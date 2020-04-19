package command

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) DiskJockey(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if bot.voice == nil {
		postMusic(s, c, y)
	} else {
		bot.playMusic(s, c, y)
	}
}

func postMusic(s *discordgo.Session, c *discordgo.Channel, y *model.Youtube) {
	msg := strings.Join([]string{y.Title, y.Description, y.UrlString()}, "\n")
	sendMessage(s, c, msg)
}

func (bot *Bot) playMusic(s *discordgo.Session, c *discordgo.Channel, y *model.Youtube) {
	err := search.DownloadMusic(y, bot.config.Search)
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

	postMusic(s, c, y)

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
