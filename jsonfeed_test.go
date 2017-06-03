package jsonfeed_test

import (
	"os"
	"testing"

	"github.com/jsvensson/jsonfeed"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type JsonFeedTestSuite struct {
	suite.Suite
	feed jsonfeed.Feed
}

func TestJsonFeedSuite(t *testing.T) {
	suite.Run(t, new(JsonFeedTestSuite))
}

func (t *JsonFeedTestSuite) SetupTest() {
	input, err := os.Open("testdata/testfeed.json")
	assert.NoError(t.T(), err, "unable to open test feed")

	t.feed, err = jsonfeed.Parse(input)
	assert.NoError(t.T(), err, "unable to parse test feed")
}

func (t *JsonFeedTestSuite) TestHasTopLevelFields() {
	assert.Equal(t.T(), "My Example Feed", t.feed.Title)
	assert.Equal(t.T(), "https://jsonfeed.org/version/1", t.feed.Version)
	assert.Equal(t.T(), "https://example.org/", t.feed.HomePageURL)
	assert.Equal(t.T(), "https://example.org/feed.json", t.feed.FeedURL)
	assert.Equal(t.T(), 2, len(t.feed.Items))
}
