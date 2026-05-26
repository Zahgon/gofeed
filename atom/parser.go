package atom

import (
	"io"

	xpp "github.com/mmcdole/goxpp"
)

var (
	// Atom elements which contain URIs
	// https://tools.ietf.org/html/rfc4287
	atomUriElements = map[string]bool{
		"icon": true,
		"id":   true,
		"logo": true,
		"uri":  true,
		"url":  true, // atom 0.3
	}
)

// Parser is an Atom Parser
type Parser struct{}

// Parse parses an xml feed into an atom.Feed
func (ap *Parser) Parse(feed io.Reader) (*Feed, error) { _ = "STUB: not implemented"; return nil, nil }

func (ap *Parser) parseRoot(p *xpp.XMLPullParser) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseEntry(p *xpp.XMLPullParser) (*Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseSource(p *xpp.XMLPullParser) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseContent(p *xpp.XMLPullParser) (*Content, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parsePerson(name string, p *xpp.XMLPullParser) (*Person, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseLink(p *xpp.XMLPullParser) (*Link, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseCategory(p *xpp.XMLPullParser) (*Category, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ap *Parser) parseGenerator(p *xpp.XMLPullParser) (*Generator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Atom 1.0
// Atom 0.3

func (ap *Parser) parseAtomText(p *xpp.XMLPullParser) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// get current base URL before it is clobbered by DecodeElement

// decode non-CDATA contents depending on type

// resolve relative URIs in URI-containing elements according to xml:base

func (ap *Parser) parseLanguage(p *xpp.XMLPullParser) string { _ = "STUB: not implemented"; return "" }

func (ap *Parser) parseVersion(p *xpp.XMLPullParser) string { _ = "STUB: not implemented"; return "" }

func (ap *Parser) stripWrappingDiv(content string) (result string) {
	_ = "STUB: not implemented"
	return ""
}
