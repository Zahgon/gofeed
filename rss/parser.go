package rss

import (
	"io"

	xpp "github.com/mmcdole/goxpp"
)

// Parser is a RSS Parser
type Parser struct{}

// Parse parses an xml feed into an rss.Feed
func (rp *Parser) Parse(feed io.Reader) (*Feed, error) { _ = "STUB: not implemented"; return nil, nil }

func (rp *Parser) parseRoot(p *xpp.XMLPullParser) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Items found in feed root

// Skip any extensions found in the feed root.

func (rp *Parser) parseChannel(p *xpp.XMLPullParser) (rss *Feed, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip element as it isn't an extension and not
// part of the spec

func (rp *Parser) parseItem(p *xpp.XMLPullParser) (item *Item, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseLink(p *xpp.XMLPullParser) (url string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rp *Parser) parseSource(p *xpp.XMLPullParser) (source *Source, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseEnclosure(p *xpp.XMLPullParser) (enclosure *Enclosure, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore any enclosure tag

func (rp *Parser) parseImage(p *xpp.XMLPullParser) (image *Image, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseGUID(p *xpp.XMLPullParser) (guid *GUID, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseCategory(p *xpp.XMLPullParser) (cat *Category, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseTextInput(p *xpp.XMLPullParser) (*TextInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseSkipHours(p *xpp.XMLPullParser) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseSkipDays(p *xpp.XMLPullParser) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseCloud(p *xpp.XMLPullParser) (*Cloud, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rp *Parser) parseVersion(p *xpp.XMLPullParser) (ver string) {
	_ = "STUB: not implemented"
	return ""
}
