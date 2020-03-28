package model

import (
	"strings"
	"time"
)

type GameRelease struct {
	title   string
	date    time.Time
	product string
	price   string
}

func NewGameRelease(title string, date time.Time, product string, price string) GameRelease {
	return GameRelease{title, date, product, price}
}

func (g *GameRelease) Title() string {
	return g.title
}

func (g *GameRelease) Date() time.Time {
	return g.date
}

func (g *GameRelease) Price() string {
	return g.price
}

func (g *GameRelease) OnelineWithoutDate(sep string) string {
	return strings.Join([]string{g.title, g.product, g.price}, sep)
}
