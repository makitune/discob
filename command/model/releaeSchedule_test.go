package model

import (
	"testing"
	"time"
)

func TestGameReleaseSchedule(t *testing.T) {
	now := time.Date(2020, time.July, 12, 0, 0, 0, 0, time.Now().Location())
	rs := []Release{
		&ReleaseMock{"DEATH STRANDING [WIN]"},
		&ReleaseMock{"Ghost of Tsushima [PS4]"},
		&ReleaseMock{"・ペーパーマリオ オリガミキング [Nintendo Switch]"},
	}
	grs := NewGameReleaseSchedule(now, now.AddDate(0, 0, 6), rs)

	t.Run("Title", func(t *testing.T) {
		t.Parallel()

		if grs.Title() != "2020/07/12〜2020/07/18の新作ゲームソフト" {
			t.Errorf("GameReleaseSchedule.Title() == %v", grs.Title())
		}
	})

	t.Run("name string", func(t *testing.T) {
		t.Parallel()

		for i, r := range grs.Releases() {
			if r != rs[i] {
				t.Errorf("GameReleaseSchedule.Releases() == %v", grs.Releases())
				return
			}
		}
	})
}

type ReleaseMock struct {
	title string
}

func (rm *ReleaseMock) Title() string {
	return rm.title
}

func (rm *ReleaseMock) Date() time.Time {
	return time.Time{}
}

func (rm *ReleaseMock) Price() string {
	return ""
}

func (rm *ReleaseMock) OnelineWithoutDate(sep string) string {
	return ""
}
