package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"threatmatrix/clean"
	"threatmatrix/threatflag"
)

var s string = `{
  "data": {
    "author_id": "1783242700789067776",
    "created_at": "2024-05-13T20:58:58.000Z",
    "edit_history_tweet_ids": ["1790124601021706634"],
    "entities": {
      "mentions": [
        {
          "start": 0,
          "end": 15,
          "username": "AltcoinDailyio",
          "id": "958118843636854784"
        },
        {
          "start": 175,
          "end": 185,
          "username": "GME_Erc20",
          "id": "1786036195056812032"
        }
      ]
    },
    "geo": {},
    "id": "1790124601021706634",
    "in_reply_to_user_id": "958118843636854784",
    "lang": "en",
    "referenced_tweets": [
      { "type": "replied_to", "id": "1790111369687413183" }
    ],
    "text": "@AltcoinDailyio \"From meme stock to market legend, $GME's story is one for nigga  nigger black monkey the history books! Who else is eagerly awaiting the next chapter? Let's write it together! $GME\" X: @GME_Erc20 https://t.co/bQaL5J82Um"
  },
  "includes": {
    "users": [
      {
        "description": "",
        "id": "1783242700789067776",
        "name": "crypto bullish",
        "username": "CryptoB64586"
      },
      {
        "description": "🎥 Follow our YouTube channel for DAILY news & opinion videos! Brothers Aaron & Austin. #Crypto commentators. #Bitcoin, #Ethereum, #NFTs, & #altcoins! 🚀",
        "id": "958118843636854784",
        "location": "Not Financial Advice",
        "name": "Altcoin Daily",
        "username": "AltcoinDailyio"
      }
    ],
    "tweets": [
      {
        "author_id": "1783242700789067776",
        "created_at": "2024-05-13T20:58:58.000Z",
        "edit_history_tweet_ids": ["1790124601021706634"],
        "entities": {
          "mentions": [
            {
              "start": 0,
              "end": 15,
              "username": "AltcoinDailyio",
              "id": "958118843636854784"
            },
            {
              "start": 175,
              "end": 185,
              "username": "GME_Erc20",
              "id": "1786036195056812032"
            }
          ]
        },
        "geo": {},
        "id": "1790124601021706634",
        "in_reply_to_user_id": "958118843636854784",
        "lang": "en",
        "referenced_tweets": [
          { "type": "replied_to", "id": "1790111369687413183" }
        ],
        "text": "@AltcoinDailyio \"From meme stock to market legend, $GME's story is one for the history books! Who else is eagerly awaiting the next chapter? Let's write it together! $GME\" X: @GME_Erc20 https://t.co/bQaL5J82Um"
      },
      {
        "author_id": "958118843636854784",
        "created_at": "2024-05-13T20:06:23.000Z",
        "edit_history_tweet_ids": ["1790111369687413183"],
        "entities": {},
        "geo": {},
        "id": "1790111369687413183",
        "lang": "en",
        "referenced_tweets": [
          { "type": "quoted", "id": "1790081979939172690" }
        ],
        "text": "The $DOT community would never... https://t.co/mr335tlvoW"
      }
    ]
  },
  "errors": [
    {
      "parameter": "entities.mentions.username",
      "resource_id": "GME_Erc20",
      "value": "GME_Erc20",
      "detail": "User has been suspended: [GME_Erc20].",
      "title": "Forbidden",
      "resource_type": "user",
      "type": "https://api.twitter.com/2/problems/resource-not-found"
    }
  ],
  "matching_rules": [
    { "id": "1785264155941072897", "tag": "funny things" },
    { "id": "1785264155941072898", "tag": "" }
  ]
}
`

func main() {

	var err error
	jsonData := []byte(s)

	var twitterData TweetData

	err = json.Unmarshal(jsonData, &twitterData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	cleanText, cleanErr := clean.NormalizeAndDemojize(twitterData.Data.Text)
	if cleanErr != nil {
		fmt.Println("SAS")
	}

	tdata := twitterData.Data
	authorId := tdata.AuthorID
	tweetId := tdata.ID
	createdAt := tdata.CreatedAt
	inReplyTo := tdata.InReplyToUserID
	lang := tdata.Lang
	text := cleanText
	rules := twitterData.Rules

	includes := twitterData.Includes
	name, handleName, description, location := func() (string, string, string, string) {
		for _, v := range includes.Users {
			if authorId == v.ID {

				return v.Name, v.Username, v.Description, v.Location
			}
		}

		return "", "", "", ""
	}()

	tokens := TokenizeTweet(cleanText)

	ngrams := generateNGrams(tokens, 10)

	data, err := threatflag.LoadWordlists("en")
	if err != nil {
		fmt.Println(err)
	}
	categories := make(map[string]int)
	for _, v := range data {
		for _, mapRow := range v {
			for _, word := range ngrams {
				if mapRow["local term"] == word {

					// fmt.Println(mapRow["local term"], word)
					if count, exists := categories[mapRow["category"]]; exists {
						// If the category exists, increment the count by 1
						categories[mapRow["category"]] = count + 1
					} else {
						// If the category does not exist, add it with a count of 1
						categories[mapRow["category"]] = 1
					}
				}
			}
		}
	}
	// fmt.Println(categories)

	threat := func() int {
		sum := 0
		for _, v := range categories {
			sum += v
		}
		return sum
	}()

	jsonBytes, err := json.Marshal(categories)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	jsonString := string(jsonBytes)

	var pub = PublishTweetRecord{
		TweetID:         tweetId,
		TweetLink:       "",
		CreatedAt:       createdAt,
		Text:            text,
		InReplyTo:       inReplyTo,
		Lang:            lang,
		Location:        location,
		UserID:          authorId,
		UserName:        name,
		HandleName:      handleName,
		UserDescription: description,
		Threat:          threat,
		Categories:      jsonString,
		Rules:           rules,
	}
	fmt.Println(pub)
	jsonsData, earr := json.MarshalIndent(pub, "", "  ")
	if earr != nil {
		fmt.Println("Error marshaling to JSON:", earr)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonsData))
}

// generateNGrams generates n-grams for a given slice of tokens and n value.
func generateNGrams(tokens []string, n int) []string {
	var ngramsList [][]string
	for i := 1; i <= n; i++ {
		for j := 0; j < len(tokens)-i+1; j++ {
			ngramsList = append(ngramsList, tokens[j:j+i])
		}
	}

	var allgrams []string
	for _, ngrams := range ngramsList {
		allgrams = append(allgrams, strings.Join(ngrams, " "))
	}

	return allgrams
}

// TokenizeTweet breaks down the tweet into tokens considering various Twitter-specific elements.
func TokenizeTweet(tweet string) []string {
	// Regular expressions for different tweet elements
	regexps := map[string]*regexp.Regexp{
		"url":         regexp.MustCompile(`(https?://\S+)`),
		"hashtag":     regexp.MustCompile(`(\#\w+)`),
		"mention":     regexp.MustCompile(`(\@\w+)`),
		"punctuation": regexp.MustCompile(`[\.,\?!;:\(\)]`),
	}

	// Function to replace and split the tweet into tokens
	replaceAndSplit := func(r *regexp.Regexp, str, repl string) []string {
		return strings.Fields(r.ReplaceAllString(str, repl))
	}

	// Slice to hold the tokens
	var tokens []string

	// Tokenize URLs
	for _, match := range regexps["url"].FindAllString(tweet, -1) {
		tokens = append(tokens, match)
		tweet = strings.Replace(tweet, match, " ", -1)
	}

	// Tokenize hashtags
	for _, match := range regexps["hashtag"].FindAllString(tweet, -1) {
		tokens = append(tokens, match)
		tweet = strings.Replace(tweet, match, " ", -1)
	}

	// Tokenize mentions
	for _, match := range regexps["mention"].FindAllString(tweet, -1) {
		tokens = append(tokens, match)
		tweet = strings.Replace(tweet, match, " ", -1)
	}

	// Tokenize punctuation
	tweet = regexps["punctuation"].ReplaceAllString(tweet, " $0 ")

	// Tokenize the rest of the tweet
	tokens = append(tokens, replaceAndSplit(regexp.MustCompile(`\s+`), tweet, " ")...)

	return tokens
}
