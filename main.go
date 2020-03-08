package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command"
	"github.com/makitune/discob/config"
)

var (
	path = flag.String("path", "config.json", "Config file path")
)

func main() {
	flag.Parse()
	fi, err := os.Stat(*path)
	if err != nil {
		log.Fatalln(err)
	}

	if fi.IsDir() {
		log.Fatalf("%s is directory", fi.Name())
	}

	f, err := os.Open(*path)
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

	bot := command.New(cfg)
	s, err := discordgo.New()
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
	s.AddHandler(bot.PlayMusic)
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
