package main

import (
	"encoding/json"
	"time"
)

// Main structure for the JSON payload
type TweetData struct {
	Data     Data     `json:"data"`
	Includes Includes `json:"includes"`
	Rules    []Rule   `json:"matching_rules"`
}

// Data structure within the JSON payload
type Data struct {
	AuthorID            string            `json:"author_id"`
	CreatedAt           time.Time         `json:"created_at"`
	EditHistoryTweetIDs []string          `json:"edit_history_tweet_ids"`
	Entities            Entities          `json:"entities"`
	Geo                 Geo               `json:"geo"`
	ID                  string            `json:"id"`
	InReplyToUserID     string            `json:"in_reply_to_user_id"`
	Lang                string            `json:"lang"`
	ReferencedTweets    []ReferencedTweet `json:"referenced_tweets"`
	Text                string            `json:"text"`
}

type Geo struct {
	// Define the fields for the Geo object if any exist
}

// Entities structure within the Data
type Entities struct {
	Mentions []Mention `json:"mentions"`
}

// Mention structure within the Entities
type Mention struct {
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Username string `json:"username"`
	ID       string `json:"id"`
}

// ReferencedTweet structure within the Data
type ReferencedTweet struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Includes structure within the JSON payload
type Includes struct {
	Users  []User  `json:"users"`
	Tweets []Tweet `json:"tweets"`
}

// User structure within the Includes
type User struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Location    string `json:"location"`
	Name        string `json:"name"`
	Username    string `json:"username"`
}

// Tweet structure within the Includes
type Tweet struct {
	AuthorID            string            `json:"author_id"`
	CreatedAt           time.Time         `json:"created_at"`
	EditHistoryTweetIDs []string          `json:"edit_history_tweet_ids"`
	Entities            Entities          `json:"entities"`
	Geo                 Geo               `json:"geo"`
	ID                  string            `json:"id"`
	InReplyToUserID     string            `json:"in_reply_to_user_id"`
	Lang                string            `json:"lang"`
	ReferencedTweets    []ReferencedTweet `json:"referenced_tweets"`
	Text                string            `json:"text"`
}

// Rule structure for the matching rules
type Rule struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type TweetRecord struct {
	ID              int             `json:"id"`               // Primary Key
	Source          string          `json:"source"`           // Source of the tweet
	Platform        string          `json:"platform"`         // Platform of the tweet
	TweetID         string          `json:"tweet_id"`         // ID of the tweet
	TweetLink       string          `json:"tweet_link"`       // Link to the tweet
	DTime           time.Time       `json:"dtime"`            // Timestamp with timezone
	Text            string          `json:"text"`             // Content of the tweet
	InReplyTo       string          `json:"in_reply_to"`      // ID of the tweet it replies to
	Lang            string          `json:"lang"`             // Language of the tweet
	Location        string          `json:"location"`         // Location from where the tweet was posted
	UserID          string          `json:"user_id"`          // ID of the user who tweeted
	UserName        string          `json:"user_name"`        // Name of the user
	UserName2       string          `json:"user_name_2"`      // Secondary name of the user
	UserDescription string          `json:"user_description"` // Description of the user
	Threat          int             `json:"threat"`           // Threat level
	Categories      json.RawMessage `json:"categories"`       // Categories of the tweet
	ReviewedH1      bool            `json:"reviewed_h1"`      // Reviewed by human 1
	ReviewedH2      bool            `json:"reviewed_h2"`      // Reviewed by human 2
	CheckoutH1      time.Time       `json:"checkout_h1"`      // Checkout time by human 1
	CheckoutH2      time.Time       `json:"checkout_h2"`      // Checkout time by human 2
	InvestigateH2   bool            `json:"investigate_h2"`   // Investigation flag for human 2
}
