package shared

import (
	"errors"
	"regexp"

	xpp "github.com/mmcdole/goxpp"
)

var (
	emailNameRgx = regexp.MustCompile(`^([^@]+@[^\s]+)\s+\(([^@]+)\)$`)
	nameEmailRgx = regexp.MustCompile(`^([^@]+)\s+\(([^@]+@[^)]+)\)$`)
	nameOnlyRgx  = regexp.MustCompile(`^([^@()]+)$`)
	emailOnlyRgx = regexp.MustCompile(`^([^@()]+@[^@()]+)$`)

	TruncatedEntity         = errors.New("truncated entity")
	InvalidNumericReference = errors.New("invalid numeric reference")
)

const CDATA_START = "<![CDATA["
const CDATA_END = "]]>"

// FindRoot iterates through the tokens of an xml document until
// it encounters its first StartTag event.  It returns an error
// if it reaches EndDocument before finding a tag.
func FindRoot(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	_ = "STUB: not implemented"
	return *new(xpp.XMLEventType), nil
}

// ParseText is a helper function for parsing the text
// from the current element of the XMLPullParser.
// This function can handle parsing naked XML text from
// an element.
func ParseText(p *xpp.XMLPullParser) (string, error) { _ = "STUB: not implemented"; return "", nil }

// StripCDATA removes CDATA tags from the string
// content outside of CDATA tags is passed via DecodeEntities
func StripCDATA(str string) string { _ = "STUB: not implemented"; return "" }

// DecodeEntities decodes escaped XML entities
// in a string and returns the unescaped string
func DecodeEntities(str string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Find the next entity

// If there is only the '&' left here

// Find the end of the entity

// it's not an entitiy. just a plain old '&' possibly with extra bytes

// Check if there is a space somewhere within the 'entitiy'.
// If there is then skip the whole thing since it's not a real entity.

// Skip the entity

// ParseNameAddress parses name/email strings commonly
// found in RSS feeds of the format "Example Name (example@site.com)"
// and other variations of this format.
func ParseNameAddress(nameAddressText string) (name string, address string) {
	_ = "STUB: not implemented"
	return "", ""
}

func indexAt(str, substr string, start int) int { _ = "STUB: not implemented"; return 0 }
