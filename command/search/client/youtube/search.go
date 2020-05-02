package youtube

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/makitune/discob/config"
)

type Search struct {
	scheme string
	host   string
	path   string
	query  url.Values
}

func NewSearch(cfg config.Search) *Search {
	q := url.Values{}
	q.Add("key", cfg.Key)
	q.Add("type", "video")
	q.Add("part", "snippet")
	q.Add("maxResults", "1")

	return &Search{
		scheme: "https",
		host:   "www.googleapis.com",
		path:   "youtube/v3/search",
		query:  q,
	}
}

func (s *Search) URL(q string) url.URL {
	s.query.Add("q", q)
	return url.URL{
		Scheme:   s.scheme,
		Host:     s.host,
		Path:     s.path,
		RawQuery: s.query.Encode(),
	}
}

func (s *Search) Parse(resp *http.Response) (*YoutubeSearchResponse, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := new(YoutubeSearchResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	if len(r.Items) == 0 {
		return nil, errors.New("No results found in Youtube")
	}

	return r, nil
}

// Youtube Data API(Search: list) Response
type YoutubeSearchResponse struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
		} `json:"snippet"`
	} `json:"items"`
}
