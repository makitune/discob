package model

import (
	"path"
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
	return path.Join("https://youtu.be", m.VideoID)
}

func (m *Music) FileName() string {
	return m.VideoID + ".m4a"
}
