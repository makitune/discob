package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command"
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

	var cfg command.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := discordgo.New()
	if err != nil {
		log.Fatalln(err)
	}

	bot.Token = cfg.Discord.Token
	bot.AddHandler(cfg.Alola)
	bot.AddHandler(cfg.FoodPorn)
	bot.AddHandler(cfg.HeadsUp)
	bot.AddHandler(cfg.Welcome)

	lock := make(chan error)
	err = bot.Open()
	if err != nil {
		log.Fatalln(err)
	}

	panic(<-lock)
}
