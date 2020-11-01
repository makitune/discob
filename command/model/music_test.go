package model

import (
	"testing"
)

func TestMusic(t *testing.T) {
	title := "米津玄師 MV「感電」"
	description := `TBS金曜ドラマ「MIU404」主題歌

米津玄師 5th Album「STRAY SHEEP」
2020.8.5 release

01. カムパネルラ
02. Flamingo
03. 感電
04. PLACEBO ＋ 野田洋次郎
05. パプリカ
06. 馬と鹿
07. 優しい人
08. Lemon
09. まちがいさがし
10. ひまわり
11. 迷える羊
12. Décolleté
13. TEENAGE RIOT
14. 海の幽霊
15. カナリヤ

▶︎ Blu-ray / DVD （アートブック盤収録）

[ LIVE VIDEO ]
米津玄師 2019 TOUR / 脊椎がオパールになる頃

01. Flamingo / 02. LOSER / 03. 砂の惑星 / 04. 飛燕 / 05. かいじゅうのマーチ / 06. アイネクライネ/ 07. 春雷 / 08. Moonlight / 09. fogbound / 10. amen / 11. Paper Flower / 12. Undercover / 13. 爱丽丝 / 14. ピースサイン /15. TEENAGE RIOT / 16. Nighthawks / 17. orion / 18. Lemon / EN1. ごめんね/ EN2. クランベリーとパンケーキ / EN3. 灰色と青

[ MUSIC VIDEO ]
01. Lemon / 02. Flamingo / 03. TEENAGE RIOT / 04. 海の幽霊 / 05. パプリカ / 06. 馬と鹿

▶︎ Tie up 
Lemon / TBS系金曜ドラマ「アンナチュラル」主題歌
Flamingo / ソニーワイヤレスヘッドホンCM
TEENAGE RIOT / ギャツビーCM
海の幽霊 / 映画「海獣の子供」主題歌
馬と鹿 / TBS系日曜劇場「ノーサイド・ゲーム」主題歌
感電 / TBS系金曜ドラマ「MIU404」主題歌

▶︎ 商品形態 
おまもり盤（初回限定）：CD＋ボックス＋キーホルダー　￥4,500+税
アートブック盤（初回限定）：CD+Blu-ray+アートブック　￥6,800+税
アートブック盤（初回限定）：CD+DVD+アートブック　￥6,800+税
通常盤：CD only　￥3,000＋税

________________________________________________________________
HP     http://reissuerecords.net
Twitter  　https://twitter.com/hachi_08
Staff Twitter      https://twitter.com/reissuerecords
OFFICIAL CHANNEL     http://www.youtube.com/user/08yakari
`
	videoID := "UFQEttrn6CQ"
	m := &Music{
		Title:       title,
		Description: description,
		VideoID:     videoID,
	}

	t.Run("URL", func(t *testing.T) {
		t.Parallel()

		if m.URL() != "https://youtu.be/UFQEttrn6CQ" {
			t.Errorf("m.URL() == %v", m.URL())
		}
	})

	t.Run("Message", func(t *testing.T) {
		t.Parallel()

		act := title + "\n" + description + "\n" + m.URL()
		if m.Message() != act {
			t.Errorf("m.Message() == %v", m.Message())
		}
	})

	t.Run("FileName", func(t *testing.T) {
		t.Parallel()

		act := videoID + ".m4a"
		if m.FileName() != act {
			t.Errorf("m.FileName() == %v", m.FileName())
		}
	})
}
