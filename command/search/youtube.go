package search

import (
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/command/search/client/youtube"
	"github.com/makitune/discob/config"
)

type Youtube struct {
	cfg config.Search
}

func NewYoutube(cfg config.Search) *Youtube {
	rand.Seed(time.Now().Unix())
	return &Youtube{cfg: cfg}
}

func (y *Youtube) Item(keyword string) (*model.Music, error) {
	req := youtube.NewSearch(y.cfg)
	u := req.URL(keyword)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	r, err := req.Parse(resp)
	if err != nil {
		return nil, err
	}

	item := &r.Items[0]
	return &model.Music{
		Title:       item.Snippet.Title,
		Description: item.Snippet.Description,
		VideoID:     item.ID.VideoID,
	}, nil
}

func (y *Youtube) RecommendedItem() (*model.Music, error) {
	req := youtube.NewPlaylistItems(y.cfg)
	u := req.URL()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	r, err := req.Parse(resp)
	if err != nil {
		return nil, err
	}

	index := rand.Intn(len(r.Items))
	item := &r.Items[index]
	return &model.Music{
		Title:       item.Snippet.Title,
		Description: item.Snippet.Description,
		VideoID:     item.ContentDetails.VideoID,
	}, nil
}

func (y *Youtube) Download(m *model.Music) (*model.Music, error) {
	cmd, err := exec.LookPath("youtube-dl")
	if err != nil {
		return nil, err
	}

	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		return nil, err
	}

	dir, err := outputDir(y.cfg)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			return nil, err
		}
	}

	path := filepath.Join(dir, m.FileName())
	options := []string{
		"-f",
		"m4a",
		"-o",
		path,
		"--no-cache-dir",
	}

	args := append(options, m.URL())
	c := exec.Command(cmd, args...)
	err = c.Run()
	if err != nil {
		return nil, err
	}

	m.FilePath = &path
	return m, nil
}
