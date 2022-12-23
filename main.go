package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/config"
)

var (
	cfgOption = flag.String("config", "config.json", "Config file path")
)

func main() {
	flag.Parse()
	fi, err := os.Stat(*cfgOption)
	if err != nil {
		log.Fatalln(err)
	}

	if fi.IsDir() {
		log.Fatalf("%s is directory", fi.Name())
	}

	f, err := os.Open(*cfgOption)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	var cfg config.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	bot := dependencyInject(cfg)
	s, err := discordgo.New(cfg.Discord.Token)
	if err != nil {
		log.Fatalln(err)
	}

	s.Token = cfg.Discord.Token
	s.AddHandler(bot.Announce)
	s.AddHandler(bot.Alola)
	s.AddHandler(bot.DiskJockey)
	s.AddHandler(bot.FoodPorn)
	s.AddHandler(bot.JoinVoiceChannel)
	s.AddHandler(bot.LeaveVoiceChannel)
	s.AddHandler(bot.Paldea)
	s.AddHandler(bot.StopMusic)
	s.AddHandler(bot.Welcome)
	s.AddHandler(bot.Wikipedia)

	lock := make(chan error)
	err = s.Open()
	if err != nil {
		log.Fatalln(err)
	}

	panic(<-lock)
}

func dependencyInject(cfg config.Config) *command.Bot {
	return &command.Bot{
		Config:     cfg,
		LoginChans: make(map[string]chan struct{}),
		Repository: search.NewYoutube(cfg.Search),
	}
}
