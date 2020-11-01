package model

import (
	"testing"
	"time"
)

func TestGameRelease(t *testing.T) {
	title := "ゾンビサバイバル コロニービルダー They Are Billions [PS4]"
	date := time.Date(2020, time.August, 20, 0, 0, 0, 0, time.Now().UTC().Location())
	product := "スパイク・チュンソフト"
	price := "¥3,800"
	gr := NewGameRelease(title, date, product, price)

	t.Run("Title", func(t *testing.T) {
		t.Parallel()

		if gr.Title() != title {
			t.Errorf("GameRelease.Title() == %v", gr.Title())
		}
	})

	t.Run("Date", func(t *testing.T) {
		t.Parallel()

		if gr.Date() != date {
			t.Errorf("GameRelease.Date() == %v", gr.Date())
		}
	})

	t.Run("Price", func(t *testing.T) {
		t.Parallel()

		if gr.Price() != price {
			t.Errorf("GameRelease.Price() == %v", gr.Price())
		}
	})

	t.Run("OnelineWithoutDate", func(t *testing.T) {
		t.Parallel()

		exp := gr.OnelineWithoutDate(",")
		act := title + "," + product + "," + price
		if exp != act {
			t.Errorf("GameRelease.OnelineWithoutDate() == %v", exp)
		}
	})
}
