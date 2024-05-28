package main

import (
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

type PublishTweetRecord struct {
	TweetID         string    // ID of the tweet
	TweetLink       string    // Link to the tweet
	CreatedAt       time.Time // Timestamp with timezone
	Text            string    // Content of the tweet
	InReplyTo       string    // ID of the tweet it replies to
	Lang            string    // Language of the tweet
	Location        string    // Location from where the tweet was posted
	UserID          string    // ID of the user who tweeted
	UserName        string    // Name of the user
	HandleName      string    // Secondary name of the user
	UserDescription string    // Description of the user
	Threat          int       // Threat level
	Categories      string    // Categories of the tweet
	Rules           []Rule
}
