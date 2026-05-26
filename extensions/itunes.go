package ext

// ITunesFeedExtension is a set of extension
// fields for RSS feeds.
type ITunesFeedExtension struct {
	Author     string            `json:"author,omitempty"`
	Block      string            `json:"block,omitempty"`
	Categories []*ITunesCategory `json:"categories,omitempty"`
	Explicit   string            `json:"explicit,omitempty"`
	Keywords   string            `json:"keywords,omitempty"`
	Owner      *ITunesOwner      `json:"owner,omitempty"`
	Subtitle   string            `json:"subtitle,omitempty"`
	Summary    string            `json:"summary,omitempty"`
	Image      string            `json:"image,omitempty"`
	Complete   string            `json:"complete,omitempty"`
	NewFeedURL string            `json:"newFeedUrl,omitempty"`
	Type       string            `json:"type,omitempty"`
}

// ITunesItemExtension is a set of extension
// fields for RSS items.
type ITunesItemExtension struct {
	Author            string `json:"author,omitempty"`
	Block             string `json:"block,omitempty"`
	Duration          string `json:"duration,omitempty"`
	Explicit          string `json:"explicit,omitempty"`
	Keywords          string `json:"keywords,omitempty"`
	Subtitle          string `json:"subtitle,omitempty"`
	Summary           string `json:"summary,omitempty"`
	Image             string `json:"image,omitempty"`
	IsClosedCaptioned string `json:"isClosedCaptioned,omitempty"`
	Episode           string `json:"episode,omitempty"`
	Season            string `json:"season,omitempty"`
	Order             string `json:"order,omitempty"`
	EpisodeType       string `json:"episodeType,omitempty"`
}

// ITunesCategory is a category element for itunes feeds.
type ITunesCategory struct {
	Text        string          `json:"text,omitempty"`
	Subcategory *ITunesCategory `json:"subcategory,omitempty"`
}

// ITunesOwner is the owner of a particular itunes feed.
type ITunesOwner struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// NewITunesFeedExtension creates an ITunesFeedExtension given an
// extension map for the "itunes" key.
func NewITunesFeedExtension(extensions map[string][]Extension) *ITunesFeedExtension {
	_ = "STUB: not implemented"
	return nil
}

// NewITunesItemExtension creates an ITunesItemExtension given an
// extension map for the "itunes" key.
func NewITunesItemExtension(extensions map[string][]Extension) *ITunesItemExtension {
	_ = "STUB: not implemented"
	return nil
}

func parseImage(extensions map[string][]Extension) (image string) {
	_ = "STUB: not implemented"
	return ""
}

func parseOwner(extensions map[string][]Extension) (owner *ITunesOwner) {
	_ = "STUB: not implemented"
	return nil
}

func parseCategories(extensions map[string][]Extension) (categories []*ITunesCategory) {
	_ = "STUB: not implemented"
	return nil
}
