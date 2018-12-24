package command

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/config"
	"github.com/makitune/discob/errr"
)

const dem = "Something bad happened"

type Bot struct {
	config          config.Config
	loginChans      map[string]chan struct{}
	voiceConnection *discordgo.VoiceConnection
}

func New(cfg config.Config) (bot *Bot) {
	return &Bot{
		config:     cfg,
		loginChans: make(map[string]chan struct{}),
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (bot *Bot) sendErrorMessage(s *discordgo.Session, c *discordgo.Channel, err error) {
	if err != nil {
		errr.Printf("%s\n", err)
	}

	msg := bot.config.Command.ErrorMessage
	if len(msg) == 0 {
		msg = dem
	}
	_, err = s.ChannelMessageSend(c.ID, msg)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (bot *Bot) sendImage(s *discordgo.Session, c *discordgo.Channel, keyword string) {
	me, err := search.SearchImage(keyword, bot.config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	_, err = s.ChannelMessageSendEmbed(c.ID, me)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (bot *Bot) foodPornMessage() (string, error) {
	return any(bot.config.Command.FoodPorn.Messages)
}

func (bot *Bot) headsUpMessage() (string, error) {
	return any(bot.config.Command.HeadsUp.Messages)
}

func (bot *Bot) welcomeKeyword() (string, error) {
	return any(bot.config.Command.Welcome.Keywords)
}

func (bot *Bot) welcomeMessage() (string, error) {
	return any(bot.config.Command.Welcome.Messages)
}

func (bot *Bot) joinVoiceChannelMessage() (string, error) {
	return any(bot.config.Command.JoinVoiceChannel.Messages)
}

func (bot *Bot) defectVoiceChannelMessage() (string, error) {
	return any(bot.config.Command.DefectVoiceChannel.Messages)
}

func any(target []string) (string, error) {
	max := len(target)
	if max == 0 {
		return "", errors.New("Not configuration")
	}
	return target[rand.Intn(max)], nil
}
