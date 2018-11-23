package command

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
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

func (bot *Bot) FoodPorn(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, trg := range fptrg {
		if strings.Contains(m.Content, trg) {
			user := m.Author
			if user.Username == bot.config.Discord.UserName || user.Bot {
				return
			}

			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				errr.Printf("%s\n", err)
				return
			}

			msg, err := bot.foodPornMessage()
			if err != nil {
				msg = dfm
			}
			sendMessage(s, c, msg)

			max := len(fpkws)
			bot.sendImage(s, c, fpkws[rand.Intn(max)])
			return
		}
	}
}
