package rss

import (
	"encoding/xml"
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
)

type IPodcastRSSParser interface {
	ParseOPML(opmlContent []byte) ([]outline, error)
	Client(xmlUrl string) (*http.Response, error)
	ParsePodcastRSSFeed(rd io.Reader) (*gofeed.Feed, error)
}

type opml struct {
	XMLName xml.Name `xml:"opml"`
	Head    head     `xml:"head"`
	Body    body     `xml:"body"`
}

type head struct {
	Title string `xml:"title"`
}

type body struct {
	Outlines []outline `xml:"outline"`
}

type outline struct {
	Outlines []outline `xml:"outline"`
	Name     string    `xml:"text,attr"`
	FeedUrl  string    `xml:"xmlUrl,attr"`
}
