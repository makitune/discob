package model

import (
	"time"
)

type Release interface {
	Title() string
	Date() time.Time
	Price() string

	OnelineWithoutDate(sep string) string
}

type ReleaseSchedule interface {
	Title() string
	Releases() []Release
}

type GameReleaseSchedule struct {
	start time.Time
	end   time.Time
	rs    []Release
}

func NewGameReleaseSchedule(start time.Time, end time.Time, rs []Release) GameReleaseSchedule {
	return GameReleaseSchedule{start, end, rs}
}

func (grs *GameReleaseSchedule) Title() string {
	f := "2006/01/02"
	return grs.start.Format(f) + "〜" + grs.end.Format(f) + "の新作ゲームソフト"
}

func (grs *GameReleaseSchedule) Releases() []Release {
	return grs.rs
}
