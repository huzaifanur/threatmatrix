package main

type TwitterData struct {
	Data          Tweet          `json:"data"`
	Includes      Includes       `json:"includes"`
	MatchingRules []MatchingRule `json:"matching_rules"`
}

type Includes struct {
	Media  []Media `json:"media"`
	Users  []User  `json:"users"`
	Tweets []Tweet `json:"tweets"`
}

type Media struct {
	Height        int            `json:"height"`
	MediaKey      string         `json:"media_key"`
	PublicMetrics map[string]int `json:"public_metrics"`
	Type          string         `json:"type"`
	URL           string         `json:"url"`
	Width         int            `json:"width"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Attachment struct {
	MediaKeys          []string `json:"media_keys"`
	MediaSourceTweetID []string `json:"media_source_tweet_id"`
}

type EditControls struct {
	EditsRemaining int    `json:"edits_remaining"`
	IsEditEligible bool   `json:"is_edit_eligible"`
	EditableUntil  string `json:"editable_until"`
}

type Hashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Mention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	MediaKey    string `json:"media_key"`
}

type Entities struct {
	Hashtags []Hashtag `json:"hashtags"`
	Mentions []Mention `json:"mentions"`
	URLs     []URL     `json:"urls"`
}

type Tweet struct {
	Attachments         Attachment             `json:"attachments"`
	AuthorID            string                 `json:"author_id"`
	ConversationID      string                 `json:"conversation_id"`
	CreatedAt           string                 `json:"created_at"`
	EditControls        EditControls           `json:"edit_controls"`
	EditHistoryTweetIDs []string               `json:"edit_history_tweet_ids"`
	Entities            Entities               `json:"entities"`
	Geo                 map[string]interface{} `json:"geo"`
	ID                  string                 `json:"id"`
	Lang                string                 `json:"lang"`
	PossiblySensitive   bool                   `json:"possibly_sensitive"`
	PublicMetrics       map[string]int         `json:"public_metrics"`
	ReferencedTweets    []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"referenced_tweets"`
	ReplySettings string `json:"reply_settings"`
	Text          string `json:"text"`
}

type MatchingRule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}
