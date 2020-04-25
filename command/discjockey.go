package command

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/errr"
)

func (bot *Bot) DiskJockey(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !bot.isMentioned(m) {
		return
	}

	if m.Author.Username == bot.Config.Discord.UserName || m.Author.Bot {
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
	mic, err := bot.Repository.Item(keyword)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	if bot.Voice == nil {
		sendMessage(s, c, mic.Message())
		return
	}

	err = bot.playMusic(s, c, mic)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}
}

func (bot *Bot) playMusic(s *discordgo.Session, c *discordgo.Channel, m *model.Music) error {
	music, err := bot.Repository.Download(m)
	if err != nil {
		return err
	}

	if bot.Voice.Playing() {
		if err := bot.Voice.Stop(); err != nil {
			return err
		}
	}

	sendMessage(s, c, music.Message())

	if err = bot.Voice.Play(music); err != nil {
		return err
	}

	if music.FilePath == nil {
		return nil
	}

	if _, err = os.Stat(*music.FilePath); !os.IsNotExist(err) {
		if err := os.Remove(*music.FilePath); err != nil {
			return err
		}
	}

	return nil
}
