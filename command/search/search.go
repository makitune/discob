package search

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Key string `json:"key"`
	Cx  string `json:"id"`
}

func SearchImage(keyword string, cfg Config) (*discordgo.MessageEmbed, error) {
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

	resp := new(Response)
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
