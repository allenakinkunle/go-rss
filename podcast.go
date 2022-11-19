package rss

import (
	"encoding/xml"
	"errors"
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
	"time"
)

var ErrParsingOPMLFile = errors.New("error parsing invalid OPML file")

type PodcastRSSParser struct {
	timeout   time.Duration
	userAgent string
}

func NewPodcastRSSParser(timeout time.Duration, userAgent string) *PodcastRSSParser {
	return &PodcastRSSParser{
		timeout:   timeout,
		userAgent: userAgent,
	}
}

func (p *PodcastRSSParser) ParseOPML(opmlContent []byte) ([]outline, error) {
	var opml opml
	var parsedOutlines []outline

	if err := xml.Unmarshal(opmlContent, &opml); err != nil {
		return nil, ErrParsingOPMLFile
	}

	for _, outline := range opml.Body.Outlines {
		if outline.Outlines != nil {
			for _, o := range outline.Outlines {
				parsedOutlines = append(parsedOutlines, o)
			}
		} else {
			parsedOutlines = append(parsedOutlines, outline)
		}
	}
	return parsedOutlines, nil
}

func (p *PodcastRSSParser) Client(xmlUrl string) (*http.Response, error) {
	client := http.Client{
		Timeout: p.timeout,
	}

	request, err := http.NewRequest(http.MethodGet, xmlUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", p.userAgent)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastRSSParser) ParsePodcastRSSFeed(rd io.Reader) (*gofeed.Feed, error) {
	parser := gofeed.NewParser()
	feed, err := parser.Parse(rd)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
