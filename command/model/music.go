package model

import (
	"net/url"
	"strings"
)

type Music struct {
	Title       string
	Description string
	VideoID     string
	FilePath    *string
}

func (m *Music) Message() string {
	return strings.Join([]string{m.Title, m.Description, m.URL()}, "\n")
}

func (m *Music) URL() string {
	query := url.Values{}
	query.Add("v", m.VideoID)
	u := url.URL{
		Scheme:   "https",
		Host:     "www.youtube.com",
		Path:     "/watch",
		RawQuery: query.Encode(),
	}
	return u.String()
}

func (m *Music) FileName() string {
	return m.VideoID + ".m4a"
}
