package command

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/config"
	"github.com/makitune/discob/errr"
)

const dem = "Something bad happened"

type Bot struct {
	config     config.Config
	loginChans map[disgord.Snowflake]chan struct{}
	voice      *model.Voice
}

func New(cfg config.Config) (bot *Bot) {
	return &Bot{
		config:     cfg,
		loginChans: make(map[disgord.Snowflake]chan struct{}),
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

func (bot *Bot) sendMessage(ctx context.Context, s disgord.Session, channelID disgord.Snowflake, msg *string, imgURL *string) error {

	params, err := newCreateMessageParams(msg, imgURL)
	if err != nil {
		return bot.sendErrorMessage(ctx, s, channelID, err)
	}
	_, err = s.CreateMessage(ctx, channelID, params)
	return err
}

func newCreateMessageParams(msg *string, imgURL *string) (*disgord.CreateMessageParams, error) {
	if msg == nil && imgURL == nil {
		return nil, errors.New("Sending message not set")
	}

	if msg == nil {
		return &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Image: &disgord.EmbedImage{URL: *imgURL},
			},
		}, nil
	}
	if imgURL == nil {
		return &disgord.CreateMessageParams{Content: *msg}, nil
	}

	return &disgord.CreateMessageParams{
		Content: *msg,
		Embed: &disgord.Embed{
			Image: &disgord.EmbedImage{URL: *imgURL},
		},
	}, nil
}

func (bot *Bot) sendErrorMessage(ctx context.Context, s disgord.Session, channelID disgord.Snowflake, err error) error {
	if err != nil {
		errr.Printf("%s\n", err)
	}

	msg := bot.config.Command.ErrorMessage
	if len(msg) == 0 {
		msg = dem
	}
	_, sendErr := s.CreateMessage(ctx, channelID, &disgord.CreateMessageParams{Content: msg})
	return sendErr
}

func (bot *Bot) isMentioned(evt *disgord.MessageCreate) bool {
	if len(evt.Message.Mentions) == 0 {
		return false
	}

	if evt.Message.MentionEveryone {
		return true
	}

	for _, u := range evt.Message.Mentions {
		if u.Username == bot.config.Discord.UserName {
			return true
		}
	}

	return false
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
