package gofeed

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/mmcdole/gofeed/atom"
	"github.com/mmcdole/gofeed/json"
	"github.com/mmcdole/gofeed/rss"
)

// ErrFeedTypeNotDetected is returned when the detection system can not figure
// out the Feed format
var ErrFeedTypeNotDetected = errors.New("Failed to detect feed type")

// HTTPError represents an HTTP error returned by a server.
type HTTPError struct {
	StatusCode int
	Status     string
}

// Error returns the string representation of the HTTP error.
func (err HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

// Parser is a universal feed parser that detects
// a given feed type, parsers it, and translates it
// to the universal feed type.
type Parser struct {
	AtomTranslator Translator
	RSSTranslator  Translator
	JSONTranslator Translator
	UserAgent      string
	AuthConfig     *Auth
	Client         *http.Client
	rp             *rss.Parser
	ap             *atom.Parser
	jp             *json.Parser
}

// Auth is a structure allowing to
// use the BasicAuth during the HTTP request
// It must be instantiated with your new Parser
type Auth struct {
	Username string
	Password string
}

// NewParser creates a universal feed parser.
func NewParser() *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses a RSS or Atom or JSON feed into
// the universal gofeed.Feed.  It takes an
// io.Reader which should return the xml/json content.
func (f *Parser) Parse(feed io.Reader) (*Feed, error) {
	_ = "STUB: not implemented"
	// Wrap the feed io.Reader in a io.TeeReader
	// so we can capture all the bytes read by the
	// DetectFeedType function and construct a new
	// reader with those bytes intact for when we
	// attempt to parse the feeds.
	return nil, nil
}

// Glue the read bytes from the detect function
// back into a new reader

// ParseURL fetches the contents of a given url and
// attempts to parse the response into the universal feed type.
func (f *Parser) ParseURL(feedURL string) (feed *Feed, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseURLWithContext fetches contents of a given url and
// attempts to parse the response into the universal feed type.
// You can instantiate the Auth structure with your Username and Password
// to use the BasicAuth during the HTTP call.
// It will be automatically added to the header of the request
// Request could be canceled or timeout via given context
func (f *Parser) ParseURLWithContext(feedURL string, ctx context.Context) (feed *Feed, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseString parses a feed XML string and into the
// universal feed type.
func (f *Parser) ParseString(feed string) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Parser) parseAtomFeed(feed io.Reader) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Parser) parseRSSFeed(feed io.Reader) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Parser) parseJSONFeed(feed io.Reader) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Parser) atomTrans() Translator { _ = "STUB: not implemented"; return *new(Translator) }

func (f *Parser) rssTrans() Translator { _ = "STUB: not implemented"; return *new(Translator) }

func (f *Parser) jsonTrans() Translator { _ = "STUB: not implemented"; return *new(Translator) }

func (f *Parser) httpClient() *http.Client { _ = "STUB: not implemented"; return nil }
