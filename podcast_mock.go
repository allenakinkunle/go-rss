package rss

import (
	"bytes"
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
)

type mockPodcastRSSParser struct{}

func NewMockPodcastRSSParser() mockPodcastRSSParser {
	return mockPodcastRSSParser{}
}

func (mockPodcastRSSParser) ParseOPML(opmlContent []byte) ([]outline, error) {
	outlines := []outline{
		{Name: "Test1", FeedUrl: "https://test1.com"},
		{Name: "Test2", FeedUrl: "https://test2.com"},
	}
	return outlines, nil
}

func (mockPodcastRSSParser) Client(xmlUrl string) (*http.Response, error) {
	rssBody := `
	<?xml version="1.0" encoding="utf-8"?>
	<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:atom="http://www.w3.org/2005/Atom">
  	<channel>
    	<title>Test Feed</title>
    	<description>Test TrimmedDescription</description>
		<itunes:explicit>no</itunes:explicit>
         <image>
          <url>https://test.com</url>
        </image>
		<itunes:owner>
          <itunes:name>Test</itunes:name>
          <itunes:email>test@test.com</itunes:email>
        </itunes:owner>
        <itunes:author>Test Author</itunes:author>
    	<item>
		  <guid isPermaLink="false">guid</guid>
		  <title>Test Item</title>
		  <description>Test TrimmedDescription</description>
		  <pubDate>Thu, 01 Jan 1970 00:00:00 +0000</pubDate>
		  <enclosure url="https://test.com.mp3" length="1234" type="audio/mpeg"/>
    	</item>
  	</channel>
	</rss>`

	resp := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(rssBody)),
	}
	return &resp, nil
}

func (mockPodcastRSSParser) ParsePodcastRSSFeed(rd io.Reader) (*gofeed.Feed, error) {
	return NewPodcastRSSParser(0, "").ParsePodcastRSSFeed(rd)
}
