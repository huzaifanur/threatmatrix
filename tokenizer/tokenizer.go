package tokenizer

import (
	"regexp"
	"strings"
)

type TweetTokenizer struct {
	// Regex patterns for Twitter-specific features
	HandlesRegex      *regexp.Regexp
	HashtagsRegex     *regexp.Regexp
	URLsRegex         *regexp.Regexp
	EmoticonsRegex    *regexp.Regexp
	HTMLTagsRegex     *regexp.Regexp
	ArrowsRegex       *regexp.Regexp
	EntitiesRegex     *regexp.Regexp
	NumsRegex         *regexp.Regexp
	PunctuationsRegex *regexp.Regexp
}

// NewTweetTokenizer initializes a new TweetTokenizer with compiled regex patterns
func NewTweetTokenizer() *TweetTokenizer {
	return &TweetTokenizer{
		HandlesRegex:      regexp.MustCompile(`(?i)(@\w+)`),
		HashtagsRegex:     regexp.MustCompile(`(?i)(#\w+)`),
		URLsRegex:         regexp.MustCompile(`(?i)(http[s]?://\S+)`),
		EmoticonsRegex:    regexp.MustCompile(`([:;]-?[)DpP])`),
		HTMLTagsRegex:     regexp.MustCompile(`(</?\w+>)`),
		ArrowsRegex:       regexp.MustCompile(`(<-|->)`),
		EntitiesRegex:     regexp.MustCompile(`(&\w+;)`),
		NumsRegex:         regexp.MustCompile(`(\d+)`),
		PunctuationsRegex: regexp.MustCompile(`([.,!?])`),
	}
}

// Tokenize splits the tweet into tokens based on the regex patterns
func (tt *TweetTokenizer) Tokenize(tweet string) []string {
	// Define a regex pattern that matches any of the special elements
	allRegex := regexp.MustCompile(`(?i)(@\w+|#\w+|http[s]?://\S+|[:;]-?[)DpP]|</?\w+>|<-|->|&\w+;|\d+|[.,!?])`)

	// Find all matches for the combined pattern
	allMatches := allRegex.FindAllString(tweet, -1)

	// Split the tweet into words, preserving the order
	words := strings.Fields(tweet)

	// Initialize an empty slice to hold the tokens
	var tokens []string

	// Create a map to keep track of matches
	matchMap := make(map[string]bool)
	for _, match := range allMatches {
		matchMap[match] = true
	}

	// Iterate over the words and match them with the special elements
	for _, word := range words {
		if matchMap[word] {
			// If the word matches any special element, add it to the tokens
			tokens = append(tokens, word)
		} else {
			// Otherwise, add the whole word as a token
			tokens = append(tokens, word)
		}
	}

	return tokens
}
