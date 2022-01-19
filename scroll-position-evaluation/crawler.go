package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// TODO: 色んな種類のソーシャルブックマークサービスに対応する.

const (
	HatenaBookmarkHotEntryURL = "https://b.hatena.ne.jp/hotentry/all"
)

func crawl() ([]string, error) {
	var urls []string
	doc, err := get(HatenaBookmarkHotEntryURL)
	if err != nil {
		return nil, err
	}

	doc.Find(".entrylist-contents-title").Each(func(i int, s *goquery.Selection) {
		url, exsits := s.Find("a").Attr("href")
		if exsits {
			urls = append(urls, url)
		}
	})
	return urls, nil
}

func get(url string) (doc *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	return
}
