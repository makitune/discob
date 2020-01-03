package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
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
	client := disgord.New(disgord.Config{
		BotToken: cfg.Discord.Token,
		Logger:   disgord.DefaultLogger(false),
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	filter, err := std.NewMsgFilter(context.Background(), client)
	if err != nil {
		log.Println(err)
	}
	client.On(disgord.EvtMessageCreate,
		filter.NotByBot,
		bot.Alola,
		bot.LeaveVoiceChannel,
		bot.DiskJockey,
		bot.FoodPorn,
		bot.JoinVoiceChannel,
		bot.PlayMusic,
		bot.StopMusic,
		bot.Wikipedia)
	client.On(disgord.EvtPresenceUpdate,
		bot.Welcome)
}
