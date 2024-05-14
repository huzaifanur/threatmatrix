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
    "data":{
       "author_id":"1519138805135011844",
       "created_at":"2024-05-13T20:59:01.000Z",
       "edit_history_tweet_ids":[
          "1790124613176828298"
       ],
       "entities":{
          "mentions":[
             {
                "start":3,
                "end":14,
                "username":"CollinRugg",
                "id":"890061634181373952"
             }
          ]
       },
       "geo":{
          
       },
       "id":"1790124613176828298",
       "lang":"en",
       "referenced_tweets":[
          {
             "type":"retweeted",
             "id":"1790054346409009158"
          }
       ],
       "text":"RT @CollinRugg: NEW: GameStop stock explodes as 'Roaring Kitty' returns to X in his first nigga nigga nigga post since bastard June 18, 2021.\n\nShort sellers have nigger  nigga suf…"
    },
    "includes":{
       "users":[
          {
             "description":"God, Family, America, Patriot, Trump2024! Football, Trucks, Camping, Stocks",
             "id":"1519138805135011844",
             "name":"Marcelo Punches",
             "username":"CarlosDanger805"
          },
          {
             "description":"Co-Owner of Trending Politics | Investor | American 🇺🇸",
             "id":"890061634181373952",
             "location":"United States",
             "name":"Collin Rugg",
             "username":"CollinRugg"
          }
       ],
       "tweets":[
          {
             "author_id":"1519138805135011844",
             "created_at":"2024-05-13T20:59:01.000Z",
             "edit_history_tweet_ids":[
                "1790124613176828298"
             ],
             "entities":{
                "mentions":[
                   {
                      "start":3,
                      "end":14,
                      "username":"CollinRugg",
                      "id":"890061634181373952"
                   }
                ]
             },
             "geo":{
                
             },
             "id":"1790124613176828298",
             "lang":"en",
             "referenced_tweets":[
                {
                   "type":"retweeted",
                   "id":"1790054346409009158"
                }
             ],
             "text":"RT @CollinRugg: NEW: GameStop stock explodes as 'Roaring Kitty' returns to X in his first post since June 18, 2021.\n\nShort sellers have suf…"
          },
          {
             "author_id":"890061634181373952",
             "created_at":"2024-05-13T16:19:48.000Z",
             "edit_history_tweet_ids":[
                "1790054346409009158"
             ],
             "entities":{
                
             },
             "geo":{
                
             },
             "id":"1790054346409009158",
             "lang":"en",
             "text":"NEW: GameStop stock explodes as 'Roaring Kitty' niggers returns to X in his first post since June 18, 2021.\n\nShort sellers have suffered a mark-to-market loss of ~$1 billion as GME is up over 70% today. (CNBC)\n\n'Roaring Kitty' went viral in 2021 for kickstarting the \"meme stock frenzy\"… https://t.co/MKhk1O8iCy"
          }
       ]
    },
    "matching_rules":[
       {
          "id":"1785264155941072897",
          "tag":"funny things"
       }
    ]
 }`

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
	// fmt.Printf("%s", cleanText)
	tokens := TokenizeTweet(cleanText)

	ngrams := generateNGrams(tokens, 10)

	// for i, v := range ngrams {
	// 	fmt.Println(i, v)
	// }

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
	fmt.Println(categories)

	threat := func() int {
		sum := 0
		for _, v := range categories {
			sum += v
		}
		return sum
	}()

	fmt.Println(threat)

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
