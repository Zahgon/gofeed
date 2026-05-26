package shared

import (
	"net/url"

	xpp "github.com/mmcdole/goxpp"
)

var (
	// HTML attributes which contain URIs
	// https://pythonhosted.org/feedparser/resolving-relative-links.html
	// To catch every possible URI attribute is non-trivial:
	// https://stackoverflow.com/questions/2725156/complete-list-of-html-tag-attributes-which-have-a-url-value
	htmlURIAttrs = map[string]bool{
		"action":     true,
		"background": true,
		"cite":       true,
		"codebase":   true,
		"data":       true,
		"href":       true,
		"poster":     true,
		"profile":    true,
		"scheme":     true,
		"src":        true,
		"uri":        true,
		"usemap":     true,
	}

	// List of xml attributes that contain URIs to be resolved relative to
	// xml:base
	// From the Atom spec https://tools.ietf.org/html/rfc4287
	uriAttrs = map[string]bool{
		"href":   true,
		"scheme": true,
		"src":    true,
		"uri":    true,
	}
)

// XMLBase.NextTag iterates through the tokens until it reaches a StartTag or
// EndTag. It resolves urls in tag attributes relative to the current xml:base.
//
// NextTag is similar to goxpp's NextTag method except it wont throw an error
// if the next immediate token isnt a Start/EndTag.  Instead, it will continue
// to consume tokens until it hits a Start/EndTag or EndDocument.
func NextTag(p *xpp.XMLPullParser) (event xpp.XMLEventType, err error) {
	_ = "STUB: not implemented"
	return *new(xpp.XMLEventType), nil
}

// resolve relative URI attributes according to xml:base
func resolveAttrs(p *xpp.XMLPullParser) error { _ = "STUB: not implemented"; return nil }

// Continue processing even if URL resolution fails (e.g., for non-HTTP URIs like at://)

// resolve u relative to b
func XmlBaseResolveUrl(b *url.URL, u string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// There's no reason someone would use a path in xml:base if they
// didn't mean for it to be a directory

// Transforms html by resolving any relative URIs in attributes
// if an error occurs during parsing or serialization, then the original string
// is returned along with the error.
func ResolveHTML(base *url.URL, relHTML string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// recursively traverse HTML resolving any relative URIs in attributes

// html.Render() always writes a complete html5 document, so strip the html
// and body tags
