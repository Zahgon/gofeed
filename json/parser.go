package json

import (
	"io"
)

// Parser is an JSON Feed Parser
type Parser struct{}

// Parse parses an json feed into an json.Feed
func (ap *Parser) Parse(feed io.Reader) (*Feed, error) { _ = "STUB: not implemented"; return nil, nil }
