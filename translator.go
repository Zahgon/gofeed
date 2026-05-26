package gofeed

import (
	"time"

	"github.com/mmcdole/gofeed/atom"
	ext "github.com/mmcdole/gofeed/extensions"
	"github.com/mmcdole/gofeed/json"
	"github.com/mmcdole/gofeed/rss"
)

// Translator converts a particular feed (atom.Feed or rss.Feed of json.Feed)
// into the generic Feed struct
type Translator interface {
	Translate(feed interface{}) (*Feed, error)
}

// DefaultRSSTranslator converts an rss.Feed struct
// into the generic Feed struct.
//
// This default implementation defines a set of
// mapping rules between rss.Feed -> Feed
// for each of the fields in Feed.
type DefaultRSSTranslator struct{}

// Translate converts an RSS feed into the universal
// feed type.
func (t *DefaultRSSTranslator) Translate(feed interface{}) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DefaultRSSTranslator) translateFeedItem(rssItem *rss.Item) (item *Item) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedTitle(rss *rss.Feed) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedDescription(rss *rss.Feed) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedLink(rss *rss.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedFeedLink(rss *rss.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedLinks(rss *rss.Feed) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedUpdated(rss *rss.Feed) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedUpdatedParsed(rss *rss.Feed) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedPublished(rss *rss.Feed) (published string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedPublishedParsed(rss *rss.Feed) (published *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedAuthor(rss *rss.Feed) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedAuthors(rss *rss.Feed) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedLanguage(rss *rss.Feed) (language string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedImage(rss *rss.Feed) *Image {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedCopyright(rss *rss.Feed) (rights string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedGenerator(rss *rss.Feed) (generator string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateFeedCategories(rss *rss.Feed) (categories []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateFeedItems(rss *rss.Feed) (items []*Item) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemTitle(rssItem *rss.Item) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemDescription(rssItem *rss.Item) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemContent(rssItem *rss.Item) (content string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemLink(rssItem *rss.Item) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemLinks(rssItem *rss.Item) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemUpdated(rssItem *rss.Item) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemUpdatedParsed(rssItem *rss.Item) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemPublished(rssItem *rss.Item) (pubDate string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemPublishedParsed(rssItem *rss.Item) (pubDate *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemAuthor(rssItem *rss.Item) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemAuthors(rssItem *rss.Item) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemGUID(rssItem *rss.Item) (guid string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultRSSTranslator) translateItemImage(rssItem *rss.Item) *Image {
	_ = "STUB: not implemented"
	return nil
}

func firstImageFromHtmlDocument(document string) *Image { _ = "STUB: not implemented"; return nil }

func (t *DefaultRSSTranslator) translateItemCategories(rssItem *rss.Item) (categories []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) translateItemEnclosures(rssItem *rss.Item) (enclosures []*Enclosure) {
	_ = "STUB: not implemented"
	return nil
}

// Accumulate the enclosures

func (t *DefaultRSSTranslator) extensionsForKeys(keys []string, extensions ext.Extensions) (matches []map[string][]ext.Extension) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultRSSTranslator) firstEntry(entries []string) (value string) {
	_ = "STUB: not implemented"
	return ""
}

// DefaultAtomTranslator converts an atom.Feed struct
// into the generic Feed struct.
//
// This default implementation defines a set of
// mapping rules between atom.Feed -> Feed
// for each of the fields in Feed.
type DefaultAtomTranslator struct{}

// Translate converts an Atom feed into the universal
// feed type.
func (t *DefaultAtomTranslator) Translate(feed interface{}) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *DefaultAtomTranslator) translateFeedItem(entry *atom.Entry) (item *Item) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedTitle(atom *atom.Feed) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedDescription(atom *atom.Feed) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedLink(atom *atom.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedFeedLink(atom *atom.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedLinks(atom *atom.Feed) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedUpdated(atom *atom.Feed) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedUpdatedParsed(atom *atom.Feed) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedAuthor(atom *atom.Feed) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedAuthors(atom *atom.Feed) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedLanguage(atom *atom.Feed) (language string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedImage(atom *atom.Feed) (image *Image) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedCopyright(atom *atom.Feed) (rights string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedGenerator(atom *atom.Feed) (generator string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateFeedCategories(atom *atom.Feed) (categories []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateFeedItems(atom *atom.Feed) (items []*Item) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemTitle(entry *atom.Entry) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemDescription(entry *atom.Entry) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemContent(entry *atom.Entry) (content string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemLink(entry *atom.Entry) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemLinks(entry *atom.Entry) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemUpdated(entry *atom.Entry) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemUpdatedParsed(entry *atom.Entry) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemPublished(entry *atom.Entry) (published string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemPublishedParsed(entry *atom.Entry) (published *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemAuthor(entry *atom.Entry) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemAuthors(entry *atom.Entry) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemGUID(entry *atom.Entry) (guid string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultAtomTranslator) translateItemImage(entry *atom.Entry) (image *Image) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemCategories(entry *atom.Entry) (categories []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) translateItemEnclosures(entry *atom.Entry) (enclosures []*Enclosure) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) firstLinkWithType(linkType string, links []*atom.Link) *atom.Link {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultAtomTranslator) firstPerson(persons []*atom.Person) (person *atom.Person) {
	_ = "STUB: not implemented"
	return nil
}

// DefaultJSONTranslator converts an json.Feed struct
// into the generic Feed struct.
//
// This default implementation defines a set of
// mapping rules between json.Feed -> Feed
// for each of the fields in Feed.
type DefaultJSONTranslator struct{}

// Translate converts an JSON feed into the universal
// feed type.
func (t *DefaultJSONTranslator) Translate(feed interface{}) (*Feed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO UserComment is missing in global Feed
// TODO NextURL is missing in global Feed
// TODO Favicon is missing in global Feed
// TODO Exipred is missing in global Feed
// TODO Hubs is not supported in json.Feed
// TODO Extensions is not supported in json.Feed

func (t *DefaultJSONTranslator) translateFeedItem(jsonItem *json.Item) (item *Item) {
	_ = "STUB: not implemented"
	return nil
}

// TODO ExternalURL is missing in global Feed
// TODO BannerImage is missing in global Feed

func (t *DefaultJSONTranslator) translateFeedTitle(json *json.Feed) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedDescription(json *json.Feed) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedLink(json *json.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedFeedLink(json *json.Feed) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedLinks(json *json.Feed) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateFeedUpdated(json *json.Feed) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedUpdatedParsed(json *json.Feed) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateFeedPublished(json *json.Feed) (published string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedPublishedParsed(json *json.Feed) (published *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateFeedAuthor(json *json.Feed) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

// Author.URL is missing in global feed
// Author.Avatar is missing in global feed

func (t *DefaultJSONTranslator) translateFeedAuthors(json *json.Feed) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

// Author.URL is missing in global feed
// Author.Avatar is missing in global feed

func (t *DefaultJSONTranslator) translateFeedLanguage(json *json.Feed) (language string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateFeedImage(json *json.Feed) (image *Image) {
	_ = "STUB: not implemented"
	// Using the Icon rather than the image
	// icon (optional, string) is the URL of an image for the feed suitable to be used in a timeline. It should be square and relatively large — such as 512 x 512
	return nil
}

func (t *DefaultJSONTranslator) translateFeedItems(json *json.Feed) (items []*Item) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemTitle(jsonItem *json.Item) (title string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemDescription(jsonItem *json.Item) (desc string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemContent(jsonItem *json.Item) (content string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemLink(jsonItem *json.Item) (link string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemLinks(jsonItem *json.Item) (links []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemUpdated(jsonItem *json.Item) (updated string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemUpdatedParsed(jsonItem *json.Item) (updated *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemPublished(jsonItem *json.Item) (pubDate string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemPublishedParsed(jsonItem *json.Item) (pubDate *time.Time) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemAuthor(jsonItem *json.Item) (author *Person) {
	_ = "STUB: not implemented"
	return nil
}

// Author.URL is missing in global feed
// Author.Avatar is missing in global feed

func (t *DefaultJSONTranslator) translateItemAuthors(jsonItem *json.Item) (authors []*Person) {
	_ = "STUB: not implemented"
	return nil
}

// Author.URL is missing in global feed
// Author.Avatar is missing in global feed

func (t *DefaultJSONTranslator) translateItemGUID(jsonItem *json.Item) (guid string) {
	_ = "STUB: not implemented"
	return ""
}

func (t *DefaultJSONTranslator) translateItemImage(jsonItem *json.Item) (image *Image) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemCategories(jsonItem *json.Item) (categories []string) {
	_ = "STUB: not implemented"
	return nil
}

func (t *DefaultJSONTranslator) translateItemEnclosures(jsonItem *json.Item) (enclosures []*Enclosure) {
	_ = "STUB: not implemented"
	return nil
}

// Title is not defined in global enclosure
// SizeInBytes is not defined in global enclosure
