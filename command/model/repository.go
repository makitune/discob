package model

type Messanger interface {
	String() string
}

type Youtuber interface {
	Item(keyword string) (*Music, error)
	RecommendedItem() (*Music, error)
	Download(m *Music) (*Music, error)
}
