package command

import (
	"math/rand"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
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

func (bot *Bot) Alola(session disgord.Session, evt *disgord.MessageCreate) {
	event := evt.Message
	if event.Content != "あろーら" {
		return
	}

	imageURL, err := search.SearchImage(pokemons[rand.Intn(len(pokemons))], bot.config.Search)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}

	err = bot.sendMessage(evt.Ctx, session, event.ChannelID, nil, imageURL)
	if err != nil {
		errr.Printf("%s\n", err)
		return
	}
}
