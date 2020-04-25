package youtube

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/makitune/discob/config"
)

type PlaylistItems struct {
	scheme string
	host   string
	path   string
	query  url.Values
}

func NewPlaylistItems(cfg config.Search) *PlaylistItems {
	q := url.Values{}
	q.Add("key", cfg.Key)
	q.Add("part", "id")
	q.Add("playlistId", "RDCLAK5uy_m1h6RaRmM8e_3k7ec4ZVJzfo2pXdLrY_k")
	q.Add("maxResults", "50")

	return &PlaylistItems{
		scheme: "https",
		host:   "www.googleapis.com",
		path:   "youtube/v3/playlistItems",
		query:  q,
	}
}

func (pli *PlaylistItems) URL() url.URL {
	return url.URL{
		Scheme:   pli.scheme,
		Host:     pli.host,
		Path:     pli.path,
		RawQuery: pli.query.Encode(),
	}
}

func (pli *PlaylistItems) Parse(resp *http.Response) (*YoutubePlaylistItemsResponse, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := new(YoutubePlaylistItemsResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	if len(r.Items) == 0 {
		return nil, errors.New("No results found in Youtube")
	}

	return r, nil
}

// Youtube Data API(PlaylistItems: list) Response
type YoutubePlaylistItemsResponse struct {
	Kind     string   `json:"kind"`
	Etag     string   `json:"etag"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []Item   `json:"items"`
}

type Item struct {
	Kind           ItemKind       `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type ContentDetails struct {
	VideoID          string `json:"videoId"`
	VideoPublishedAt string `json:"videoPublishedAt"`
}

type Snippet struct {
	PublishedAt  string       `json:"publishedAt"`
	ChannelID    ChannelID    `json:"channelId"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Thumbnails   Thumbnails   `json:"thumbnails"`
	ChannelTitle ChannelTitle `json:"channelTitle"`
	PlaylistID   PlaylistID   `json:"playlistId"`
	Position     int64        `json:"position"`
	ResourceID   ResourceID   `json:"resourceId"`
}

type ResourceID struct {
	Kind    ResourceIDKind `json:"kind"`
	VideoID string         `json:"videoId"`
}

type Thumbnails struct {
	Default  Default  `json:"default"`
	Medium   Default  `json:"medium"`
	High     Default  `json:"high"`
	Standard *Default `json:"standard,omitempty"`
	Maxres   *Default `json:"maxres,omitempty"`
}

type Default struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}

type ItemKind string

const (
	YoutubePlaylistItem ItemKind = "youtube#playlistItem"
)

type ChannelID string

const (
	UCBR860B28Hp2BmDPdntcQ ChannelID = "UCBR8-60-B28hp2BmDPdntcQ"
)

type ChannelTitle string

const (
	YouTube ChannelTitle = "YouTube"
)

type PlaylistID string

const (
	RDCLAK5UyM1H6RaRmM8E3K7Ec4ZVJzfo2PXdLrYK PlaylistID = "RDCLAK5uy_m1h6RaRmM8e_3k7ec4ZVJzfo2pXdLrY_k"
)

type ResourceIDKind string

const (
	YoutubeVideo ResourceIDKind = "youtube#video"
)
