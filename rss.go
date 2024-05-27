package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	rssFeed := RSSFeed{}

	resp, err := httpClient.Get(url)
	if err != nil {
		return rssFeed, err
	}

	defer resp.Body.Close()

	dat, er := io.ReadAll(resp.Body)
	if er != nil {
		return rssFeed, er
	}

	errr := xml.Unmarshal(dat, &rssFeed)
	if errr != nil {
		return rssFeed, errr
	}

	return rssFeed, nil
}
