package command

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

var pokemons = []string{
	"アローラロコン",
	"アローラニャース",
	"アローラライチュウ",
	"アローラサンドパン",
	"アローラガラガラ",
	"アローラコラッタ",
	"アローラダグトリオ",
	"アローラベトベトン",
	"アローラゴローン",
	"アローラナッシー",
}

func (bot *Bot) Alola(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "あろーら" {
		return
	}

	user := m.Author
	if user.Username == bot.config.Discord.UserName || user.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	bot.sendImage(s, c, pokemons[rand.Intn(len(pokemons))])
}
