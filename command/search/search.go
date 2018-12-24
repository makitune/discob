package search

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/config"
)

var (
	defaultOutputDir = "/opt/discob"
)

func SearchImage(keyword string, cfg config.Search) (*discordgo.MessageEmbed, error) {
	query := url.Values{}
	num := 10
	query.Add("key", cfg.Key)
	query.Add("cx", cfg.Cx)
	query.Add("searchType", "image")
	query.Add("num", strconv.Itoa(num))
	query.Add("q", keyword)

	u := url.URL{
		Scheme:   "https",
		Host:     "www.googleapis.com",
		Path:     "/customsearch/v1",
		RawQuery: query.Encode(),
	}
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := new(customSearchResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	num = rand.Intn(num)
	return &discordgo.MessageEmbed{
		URL:  resp.Items[num].Link,
		Type: "image",
		Image: &discordgo.MessageEmbedImage{
			URL: resp.Items[num].Link,
		},
	}, nil
}

func SearchYoutube(keyword string, cfg config.Search) (*model.Youtube, error) {
	query := url.Values{}
	query.Add("key", cfg.Key)
	query.Add("type", "video")
	query.Add("part", "snippet")
	query.Add("maxResults", "1")
	query.Add("q", keyword)

	u := url.URL{
		Scheme:   "https",
		Host:     "www.googleapis.com",
		Path:     "/youtube/v3/search",
		RawQuery: query.Encode(),
	}
	res, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := new(youtubeSearchResponse)
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.New("No results found in Youtube")
	}

	return newYoutube(resp), nil
}

func outputDir(cfg config.Search) (string, error) {
	var dir string
	if len(cfg.OutputDir) == 0 {
		dir = defaultOutputDir
	} else {
		dir = cfg.OutputDir
	}

	return filepath.Abs(dir)
}
