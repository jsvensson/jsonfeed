// Package jsonfeed is a basic implementation of the JSON Feed specification, https://jsonfeed.org/.
package jsonfeed

import (
	"encoding/json"
	"io"
	"time"
)

// Feed is a JSON Feed containing the relevant feed information and a list of items.
// Struct fields with the "omitempty" JSON property are not required for the format.
type Feed struct {
	Version     string  `json:"version"`
	Title       string  `json:"title"`
	HomePageURL string  `json:"home_page_url,omitempty"`
	FeedURL     string  `json:"feed_url,omitempty"`
	Description string  `json:"description,omitempty"`
	UserComment string  `json:"user_comment,omitempty"`
	NextURL     string  `json:"next_url,omitempty"`
	Icon        string  `json:"icon,omitempty"`
	FavIcon     string  `json:"favicon,omitempty"`
	Author      *Author `json:"author,omitempty"`
	Expired     bool    `json:"expired,omitempty"`
	Hubs        []Hub   `json:"hubs,omitempty"`
	Items       []Item  `json:"items"`
}

// Item describes an entry in a JSON feed.
type Item struct {
	ID            string       `json:"id"`
	URL           string       `json:"url,omitempty"`
	ExternalURL   string       `json:"external_url,omitempty"`
	Title         string       `json:"title,omitempty"`
	ContentHTML   string       `json:"content_html,omitempty"`
	ContentText   string       `json:"content_text,omitempty"`
	Summary       string       `json:"summary,omitempty"`
	Image         string       `json:"image,omitempty"`
	BannerImage   string       `json:"banner_image,omitempty"`
	DatePublished time.Time    `json:"date_published,omitempty"`
	DateModified  time.Time    `json:"date_modified,omitempty"`
	Author        *Author      `json:"author,omitempty"`
	Tags          []string     `json:"tags,omitempty"`
	Attachments   []Attachment `json:"attachments,omitempty"`
}

// Attachment is an optional attachment to an Item.
type Attachment struct {
	URL      string `json:"url"`
	MimeType string `json:"mime_type"`
	Title    string `json:"title,omitempty"`
	Size     int    `json:"size_in_bytes,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

// Author is the author of a JSON feed, or an individual item in a JSON feed.
type Author struct {
	Name   string `json:"name,omitempty"`
	URL    string `json:"url,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// Hub describes an endpoint that can be used to subscribe to real-time notifications from the publisher of this feed.
// See the official specification for more information.
type Hub struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Parse accepts a raw JSON Feed via a reader and parses it into a Feed struct.
func Parse(reader io.Reader) (Feed, error) {
	var feed Feed
	err := json.NewDecoder(reader).Decode(&feed)
	if err != nil {
		return Feed{}, err
	}

	return feed, nil
}
