package model

import "net/url"

type Youtube struct {
	Title       string
	Description string
	VideoID     string
}

func (y *Youtube) UrlString() string {
	query := url.Values{}
	query.Add("v", y.VideoID)
	u := url.URL{
		Scheme:   "https",
		Host:     "www.youtube.com",
		Path:     "/watch",
		RawQuery: query.Encode(),
	}
	return u.String()
}
