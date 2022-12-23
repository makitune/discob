package command

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
)

var pokemons = []string{
	"コラッタ",
	"ラッタ",
	"ライチュウ",
	"サンド",
	"サンドパン",
	"ロコン",
	"キュウコン",
	"ディグダ",
	"ダグトリオ",
	"ニャース",
	"ペルシアン",
	"イシツブテ",
	"ゴローン",
	"ゴローニャ",
	"ベトベター",
	"ベトベトン",
	"ナッシー",
	"ガラガラ",
}

func (bot *Bot) Alola(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != "あろーら" {
		return
	}

	user := m.Author
	if user.Username == bot.Config.Discord.UserName || user.Bot {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	bot.sendImage(s, c, "アローラ"+pokemons[rand.Intn(len(pokemons))])
}
