package main

type Tweet struct {
	ID                  string              `json:"id"`
	Text                string              `json:"text"`
	CreatedAt           string              `json:"created_at"`
	AuthorID            string              `json:"author_id"`
	EditHistoryTweetIDs []string            `json:"edit_history_tweet_ids"`
	EditControls        EditControls        `json:"edit_controls"`
	ConversationID      string              `json:"conversation_id"`
	NoteTweet           *NoteTweet          `json:"note_tweet,omitempty"`
	InReplyToUserID     string              `json:"in_reply_to_user_id"`
	ReferencedTweets    []ReferencedTweet   `json:"referenced_tweets"`
	Attachments         Attachments         `json:"attachments"`
	Geo                 Geo                 `json:"geo"`
	ContextAnnotations  []ContextAnnotation `json:"context_annotations"`
	Entities            Entities            `json:"entities"`
	Withheld            *Withheld           `json:"withheld,omitempty"`
	PublicMetrics       PublicMetrics       `json:"public_metrics"`
	OrganicMetrics      OrganicMetrics      `json:"organic_metrics"`
	PromotedMetrics     PromotedMetrics     `json:"promoted_metrics"`
	PossiblySensitive   bool                `json:"possibly_sensitive"`
	Lang                string              `json:"lang"`
	Source              string              `json:"source"`
	ReplySettings       string              `json:"reply_settings"`
}

type EditControls struct {
	IsEditEligible bool   `json:"is_edit_eligible"`
	EditableUntil  string `json:"editable_until"`
	EditsRemaining int    `json:"edits_remaining"`
}

type NoteTweet struct {
	Text     string   `json:"text"`
	Entities Entities `json:"entities"`
}

type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Attachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIDs   []string `json:"poll_ids"`
}

type Geo struct {
	Coordinates *Coordinates `json:"coordinates,omitempty"`
	PlaceID     string       `json:"place_id"`
}

type Coordinates struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type ContextAnnotation struct {
	Domain Domain `json:"domain"`
	Entity Entity `json:"entity"`
}

type Domain struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Entity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Entities struct {
	Annotations []Annotation `json:"annotations"`
	URLs        []URL        `json:"urls"`
	Hashtags    []Hashtag    `json:"hashtags"`
	Mentions    []Mention    `json:"mentions"`
	Cashtags    []Cashtag    `json:"cashtags"`
}

type Annotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float64 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type URL struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
	UnwoundURL  string `json:"unwound_url"`
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
}

type Cashtag struct {
	Start int    `json:"start"`
	End   int    `json:"end"`
	Tag   string `json:"tag"`
}

type Withheld struct {
	Copyright    bool     `json:"copyright"`
	CountryCodes []string `json:"country_codes"`
	Scope        string   `json:"scope"`
}

type PublicMetrics struct {
	RetweetCount    int `json:"retweet_count"`
	ReplyCount      int `json:"reply_count"`
	LikeCount       int `json:"like_count"`
	QuoteCount      int `json:"quote_count"`
	ImpressionCount int `json:"impression_count"`
	BookmarkCount   int `json:"bookmark_count"`
}

type OrganicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
}

type PromotedMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
}
