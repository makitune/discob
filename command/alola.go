package command

import (
	"discord-bot/errr"
	"math/rand"

	"github.com/bwmarrin/discordgo"
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

func (cfg *Config) Alola(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "あろーら" {
		return
	}

	user := m.Author
	if user.Username == cfg.Discord.UserName || user.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	cfg.sendImage(s, c, pokemons[rand.Intn(len(pokemons))])
}
