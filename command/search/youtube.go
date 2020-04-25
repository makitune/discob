package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/command/search/client/youtube"
	"github.com/makitune/discob/config"
)

type Youtube struct {
	cfg config.Search
}

func NewYoutube(cfg config.Search) *Youtube {
	return &Youtube{cfg: cfg}
}

func (y *Youtube) Item(keyword string) (*model.Music, error) {
	req := youtube.NewSearch(y.cfg)
	u := req.URL(keyword)
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := new(youtube.YoutubeSearchResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	if len(r.Items) == 0 {
		return nil, errors.New("No results found in Youtube")
	}

	item := &r.Items[0]
	return &model.Music{
		Title:       item.Snippet.Title,
		Description: item.Snippet.Description,
		VideoID:     item.ID.VideoID,
	}, nil
}

func (y *Youtube) PlayListItem() (*model.Music, error) {
	return nil, nil
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
	}

	args := append(options, m.URL())
	c := exec.Command(cmd, args...)
	err = c.Run()
	if err != nil {
		return nil, err
	}

	m.FilePath = &path
	fmt.Println(path)
	fmt.Println(m)
	return m, nil
}
