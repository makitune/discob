package command

import (
	"discord-bot/command/search"
	"discord-bot/errr"
	"errors"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Discord struct {
		UserName string `json:"username"`
		Token    string `json:"token"`
	} `json:"discord"`
	Search  search.Config `json:"cse"`
	Command struct {
		FoodPorrn BotCommand `json:"foodPorrn"`
		Welcome   BotCommand `json:"welcome"`
	} `json:"command"`
}

type BotCommand struct {
	Trg  []string `json:"trigger"`
	Kws  []string `json:"keywords"`
	Msgs []string `json:"messages"`
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

func sendErrorMessage(s *discordgo.Session, c *discordgo.Channel, err error) {
	if err != nil {
		errr.Printf("%s\n", err)
	}

	msg := "反応がない。。ただの屍のようだ・・・"
	_, err = s.ChannelMessageSend(c.ID, msg)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (cfg *Config) sendImage(s *discordgo.Session, c *discordgo.Channel, keyword string) {
	me, err := search.SearchImage(keyword, cfg.Search)
	if err != nil {
		sendErrorMessage(s, c, err)
		return
	}

	_, err = s.ChannelMessageSendEmbed(c.ID, me)
	if err != nil {
		errr.Printf("%s\n", err)
	}
}

func (cfg *Config) foodPorrnKeyword() (string, error) {
	return any(cfg.Command.FoodPorrn.Kws)
}

func (cfg *Config) foodPorrnMessage() (string, error) {
	return any(cfg.Command.FoodPorrn.Msgs)
}

func (cfg *Config) WelcomeKeyword() (string, error) {
	return any(cfg.Command.Welcome.Kws)
}

func (cfg *Config) welcomeMessage() (string, error) {
	return any(cfg.Command.Welcome.Msgs)
}

func any(target []string) (string, error) {
	max := len(target)
	if max == 0 {
		return "", errors.New("Not configuration")
	}
	return target[rand.Intn(max)], nil
}
