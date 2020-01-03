package command

import (
	"math/rand"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/command/search"
	"github.com/makitune/discob/errr"
)

const (
	dfm = "Take it easy"
)

var (
	fptrg = []string{
		"疲",
		"つかれ",
	}
	fpkws = []string{
		"飯テロ",
		"飯テロ 肉",
		"飯テロ 刺身",
		"飯テロ ステーキ",
		"飯テロ ラーメン",
		"飯テロ ジャンクフード",
		"飯テロ スイーツ",
	}
)

func (bot *Bot) FoodPorn(session disgord.Session, evt *disgord.MessageCreate) {
	event := evt.Message
	keyword := fpkws[rand.Intn(len(fpkws))]

	for _, trg := range fptrg {
		if !strings.Contains(event.Content, trg) {
			continue
		}

		msg, err := bot.foodPornMessage()
		if err != nil {
			msg = dfm
		}
		imgURL, err := search.SearchImage(keyword, bot.config.Search)
		if err != nil {
			errr.Printf("%s\n", err)
			return
		}

		err = bot.sendMessage(evt.Ctx, session, event.ChannelID, &msg, imgURL)
		if err != nil {
			errr.Printf("%s\n", err)
			return
		}
	}
}
