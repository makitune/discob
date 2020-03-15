package search

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

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

func DownloadMusic(y *model.Youtube, cfg config.Search) error {
	cmd, err := exec.LookPath("youtube-dl")
	if err != nil {
		return err
	}

	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		return err
	}

	dir, err := outputDir(cfg)
	if err != nil {
		return err
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			return err
		}
	}

	filename := y.VideoID + ".m4a"
	path := filepath.Join(dir, filename)
	options := []string{
		"-f",
		"bestaudio[ext=m4a]",
		"-o",
		path,
	}

	args := append(options, y.UrlString())
	err = exec.Command(cmd, args...).Run()
	if err != nil {
		return err
	}

	y.FilePath = &path
	return nil
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

func SearchWikipediaURL(keyword string) (string, error) {
	q := url.Values{}
	q.Add("q", keyword)
	u := url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "search",
		RawQuery: q.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str := string(data)
	start := strings.Index(str, "https://ja.wikipedia.org/wiki/%25")
	if start == -1 {
		return "", errors.New(keyword + " not found")
	}
	end := strings.Index(str[start:], "&")
	title, err := humanize(str[start+30 : start+end])
	if err != nil {
		return "nil", err
	}
	return "https://ja.wikipedia.org/wiki/" + title, nil
}

func humanize(keyword string) (string, error) {
	s := strings.Replace(keyword, "%25", "%", -1)
	return url.QueryUnescape(s)
}
