package command

import (
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/config"
	"github.com/makitune/discob/errr"
)

const dem = "Something bad happened"

type Bot struct {
	AnnounceChans chan struct{}
	Config        config.Config
	LoginChans    map[string]chan struct{}
	Voice         *model.Voice
	Repository    model.Youtuber
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

	msg := bot.Config.Command.ErrorMessage
	if len(msg) == 0 {
		msg = dem
	}
	_, err = s.ChannelMessageSend(c.ID, msg)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (bot *Bot) sendImage(s *discordgo.Session, c *discordgo.Channel, keyword string) {
	me, err := search.SearchImage(keyword, bot.Config.Search)
	if err != nil {
		bot.sendErrorMessage(s, c, err)
		return
	}

	_, err = s.ChannelMessageSendEmbed(c.ID, me)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (bot *Bot) isMentioned(m *discordgo.MessageCreate) bool {
	if len(m.Mentions) == 0 {
		return false
	}

	for _, mu := range m.Mentions {
		if mu.Username == bot.Config.Discord.UserName {
			return true
		}
	}

	return false
}

func (bot *Bot) foodPornMessage() (string, error) {
	return any(bot.Config.Command.FoodPorn.Messages)
}

func (bot *Bot) headsUpMessage() (string, error) {
	return any(bot.Config.Command.HeadsUp.Messages)
}

func (bot *Bot) welcomeKeyword() (string, error) {
	return any(bot.Config.Command.Welcome.Keywords)
}

func (bot *Bot) welcomeMessage() (string, error) {
	return any(bot.Config.Command.Welcome.Messages)
}

func (bot *Bot) joinVoiceChannelMessage() (string, error) {
	return any(bot.Config.Command.JoinVoiceChannel.Messages)
}

func (bot *Bot) defectVoiceChannelMessage() (string, error) {
	return any(bot.Config.Command.LeaveVoiceChannel.Messages)
}

func any(target []string) (string, error) {
	max := len(target)
	if max == 0 {
		return "", errors.New("Not configuration")
	}
	return target[rand.Intn(max)], nil
}
