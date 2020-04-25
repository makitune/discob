package model

type Messanger interface {
	String() string
}

type Youtuber interface {
	Item(keyword string) (*Music, error)
	PlayListItem() (*Music, error)
	Download(m *Music) (*Music, error)
}
