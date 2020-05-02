package search

import (
	"bytes"
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"

	"github.com/makitune/discob/command/model"
	"github.com/makitune/discob/config"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
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

func SearchGameReleaseSchedule(from time.Time, to time.Time) (*string, error) {
	urls, err := releasePages(from, to)
	if err != nil {
		return nil, err
	}

	docs := []*goquery.Document{}
	for _, u := range urls {
		d, err := getDoc(u)
		if err != nil {
			return nil, err
		}
		docs = append(docs, d)
	}

	rs := []model.Release{}
	for _, d := range docs {
		grs, err := parseGameReleases(d, from, to)
		if err != nil {
			return nil, err
		}
		rs = append(rs, grs...)
	}

	sch := model.NewGameReleaseSchedule(from, to, rs)
	str := discordMarkdown(&sch)
	return &str, err
}

func releasePages(start time.Time, end time.Time) ([]string, error) {
	if !end.After(start) {
		return []string{}, errors.New("start time is later than end time")
	}

	u, err := url.Parse("https://kakaku.com/game/release/")
	p := u.Path
	if err != nil {
		return []string{}, err
	}

	v := url.Values{}
	v.Add("Date", start.Format("200601"))
	u.Path = path.Join(p, v.Encode())
	urls := []string{u.String()}

	if !(start.YearDay() < end.YearDay() && start.Month() < end.Month()) {
		return urls, nil
	}

	v.Set("Date", end.Format("200601"))
	u.Path = path.Join(p, v.Encode())
	urls = append(urls, u.String())
	return urls, nil
}

func getDoc(u string) (*goquery.Document, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)

	det := chardet.NewTextDetector()
	detRslt, err := det.DetectBest(buf)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(buf)
	r, err := charset.NewReaderLabel(detRslt.Charset, br)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromReader(r)
}

func parseGameReleases(doc *goquery.Document, from time.Time, to time.Time) ([]model.Release, error) {
	grs := []model.Release{}
	cd := time.Time{}
	var err error
	doc.Find("#titleSche > tbody > tr").EachWithBreak(func(_ int, tr *goquery.Selection) bool {
		td := tr.Children()
		if td.HasClass("releaseLine") {
			return true
		}

		title := tr.Find("td.gameTitle").Text()
		if title == html.UnescapeString("&nbsp;") || len(title) == 0 {
			return true
		}

		d := strings.Split(tr.Find("td.weekly").Text(), "（")[0]

		if len(d) != 0 {
			t, pe := time.Parse("2006年1月2日", d)
			if pe != nil {
				err = pe
				return false
			}
			cd = t
		}

		if from.YearDay() <= cd.YearDay() {
			gr := model.NewGameRelease(title, cd, tr.Find("td.gameProduct").Text(), tr.Find("td.gamePrice").Text())
			grs = append(grs, &gr)
		}

		return cd.YearDay() < to.YearDay()
	})

	if err != nil {
		return nil, err
	}

	return grs, nil
}

func discordMarkdown(s model.ReleaseSchedule) string {
	str := "__**" + s.Title() + "**__"

	cd := time.Time{}
	for _, g := range s.Releases() {
		if cd.YearDay() != g.Date().YearDay() {
			cd = g.Date()
			str = str + "\n\n" + "__*" + cd.Format("2006年01月02日") + "*__"
		}

		str = str + "\n・" + g.OnelineWithoutDate(" / ")
	}
	return str
}
