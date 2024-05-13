package main

import (
	"encoding/json"
	"fmt"
	"strings"
	// "encoding/json"
	// "fmt"
	// "slices"
	// "strings"
	// e "threatmatrix/encoding"
	// "threatmatrix/threatflag"
	// t "threatmatrix/tokenizer"
)

var s string = `{"data":{"attachments":{"media_keys":["7_1787811000881713153"],"media_source_tweet_id":["1787811070830055877"]},"author_id":"190740340","conversation_id":"1788459752172744762","created_at":"2024-05-09T06:43:27.000Z","edit_controls":{"edits_remaining":5,"is_edit_eligible":false,"editable_until":"2024-05-09T07:43:27.000Z"},"edit_history_tweet_ids":["1788459752172744762"],"entities":{"annotations":[{"start":75,"end":85,"probability":0.7074,"type":"Person","normalized_text":"cascahexene"}],"mentions":[{"start":3,"end":17,"username":"Rainmaker1973","id":"177101260"}],"urls":[{"start":88,"end":111,"url":"https://t.co/quqx909Rpf","expanded_url":"https://twitter.com/i/status/1787811070830055877/video/1","display_url":"pic.twitter.com/quqx909Rpf","media_key":"7_1787811000881713153"}]},"geo":{},"id":"1788459752172744762","lang":"en","possibly_sensitive":false,"public_metrics":{"retweet_count":72,"reply_count":0,"like_count":0,"quote_count":0,"bookmark_count":0,"impression_count":0},"referenced_tweets":[{"type":"retweeted","id":"1788446299613806934"}],"reply_settings":"everyone","text":"RT @Rainmaker1973: The distinctive appearance of the Asian leopard cat\n\n[😊 😍🙂☺️😑😶🫥   cascahexene]\nhttps://t.co/quqx909Rpf"},"includes":{"media":[{"duration_ms":13117,"height":1920,"media_key":"7_1787811000881713153","preview_image_url":"https://pbs.twimg.com/ext_tw_video_thumb/1787811000881713153/pu/img/r8d_GX5stssC8g8m.jpg","public_metrics":{"view_count":37536},"type":"video","variants":[{"content_type":"application/x-mpegURL","url":"https://video.twimg.com/ext_tw_video/1787811000881713153/pu/pl/BXttK0Ejr6POpZqs.m3u8?tag=14&container=cmaf"},{"bit_rate":2176000,"content_type":"video/mp4","url":"https://video.twimg.com/ext_tw_video/1787811000881713153/pu/vid/avc1/720x1280/jI71-ko0ROJHsCJg.mp4?tag=14"},{"bit_rate":632000,"content_type":"video/mp4","url":"https://video.twimg.com/ext_tw_video/1787811000881713153/pu/vid/avc1/320x568/sfGI0zDJxji3TGNL.mp4?tag=14"},{"bit_rate":10368000,"content_type":"video/mp4","url":"https://video.twimg.com/ext_tw_video/1787811000881713153/pu/vid/avc1/1080x1920/BSqI1Jio0In3uK2K.mp4?tag=14"},{"bit_rate":950000,"content_type":"video/mp4","url":"https://video.twimg.com/ext_tw_video/1787811000881713153/pu/vid/avc1/480x852/XCZDlVbPNuBfnliV.mp4?tag=14"}],"width":1080}],"users":[{"id":"190740340","name":"Nikita, ne tuda!","username":"RbICb"},{"id":"177101260","name":"Massimo","username":"Rainmaker1973"}],"tweets":[{"attachments":{"media_keys":["7_1787811000881713153"],"media_source_tweet_id":["1787811070830055877"]},"author_id":"190740340","conversation_id":"1788459752172744762","created_at":"2024-05-09T06:43:27.000Z","edit_controls":{"edits_remaining":5,"is_edit_eligible":false,"editable_until":"2024-05-09T07:43:27.000Z"},"edit_history_tweet_ids":["1788459752172744762"],"entities":{"annotations":[{"start":75,"end":85,"probability":0.7074,"type":"Person","normalized_text":"cascahexene"}],"mentions":[{"start":3,"end":17,"username":"Rainmaker1973","id":"177101260"}],"urls":[{"start":88,"end":111,"url":"https://t.co/quqx909Rpf","expanded_url":"https://twitter.com/i/status/1787811070830055877/video/1","display_url":"pic.twitter.com/quqx909Rpf","media_key":"7_1787811000881713153"}]},"geo":{},"id":"1788459752172744762","lang":"en","possibly_sensitive":false,"public_metrics":{"retweet_count":72,"reply_count":0,"like_count":0,"quote_count":0,"bookmark_count":0,"impression_count":0},"referenced_tweets":[{"type":"retweeted","id":"1788446299613806934"}],"reply_settings":"everyone","text":"RT @Rainmaker1973: The distinctive appearance of the Asian leopard cat\n\n[📹 cascahexene]\nhttps://t.co/quqx909Rpf"},{"attachments":{"media_keys":["7_1787811000881713153"],"media_source_tweet_id":["1787811070830055877"]},"author_id":"177101260","conversation_id":"1788446299613806934","created_at":"2024-05-09T05:50:00.000Z","edit_controls":{"edits_remaining":5,"is_edit_eligible":true,"editable_until":"2024-05-09T06:50:00.000Z"},"edit_history_tweet_ids":["1788446299613806934"],"entities":{"annotations":[{"start":56,"end":66,"probability":0.6268,"type":"Person","normalized_text":"cascahexene"}],"urls":[{"start":69,"end":92,"url":"https://t.co/quqx909Rpf","expanded_url":"https://twitter.com/i/status/1787811070830055877/video/1","display_url":"pic.twitter.com/quqx909Rpf","media_key":"7_1787811000881713153"}]},"geo":{},"id":"1788446299613806934","lang":"en","possibly_sensitive":false,"public_metrics":{"retweet_count":72,"reply_count":17,"like_count":639,"quote_count":5,"bookmark_count":67,"impression_count":23725},"reply_settings":"everyone","text":"The distinctive appearance of the Asian leopard cat\n\n[📹 cascahexene]\nhttps://t.co/quqx909Rpf"}]},"matching_rules":[{"id":"1785264155941072899","tag":"cats with media"},{"id":"1785264155941072896","tag":"happy cats with media"}]}`

func main() {

	var err error
	jsonData := []byte(s)

	// Define a variable of the type of your struct
	var twitterData TwitterData

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &twitterData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// var text string
	// // Now you can use 'twitterData' which is populated with the JSON data
	// // fmt.Println(twitterData.Data.Lang)
	// text, err = e.EncodeAndNormalizeAndDemojize(twitterData.Data.Text, twitterData.Data.Lang)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling JSON:", err)
	// 	return
	// }

	// tokenizer := t.NewTweetTokenizer()
	// tokens := tokenizer.Tokenize((text))
	// wordlists, _ := threatflag.LoadWordlists("en")
	// // # FIND MAX NGRAM SIZE FROM BAD WORDS

	// // fmt.Println((wordlists[0]["local term"]))
	// var n int
	// for _, value := range wordlists {
	// 	localTerms := value["local term"]

	// 	for _, v := range localTerms {
	// 		tokens = tokenizer.Tokenize(v)
	// 		if len(tokens) > n {

	// 			n = len(tokens)
	// 		}
	// 	}
	// }

	// ngrams := generateNGrams(tokens, n)
	// for _, ngram := range ngrams {
	// 	println(ngram)
	// }

	// // New code starts here
	// categories := make(map[string]int)
	// for _, wordlist := range wordlists {
	// 	for _, category := range wordlist["simple category"] {
	// 		categories[strings.ToLower(strings.TrimSpace(category))] = 0
	// 	}
	// }

	// for _, wordlist := range wordlists {
	// 	for _, localTerm := range wordlist["local term"] {
	// 		if slices.Contains(ngrams, strings.ToLower(localTerm)) {
	// 			category := strings.ToLower(strings.TrimSpace(wordlist["simple category"][0])) // Assuming the first entry is the category
	// 			categories[category]++
	// 		}
	// 	}
	// }

	// threat := 0
	// for _, count := range categories {
	// 	threat += count
	// }

	// fmt.Println("Threat level:", threat)
	// fmt.Println("Categories count:")
	// for category, count := range categories {
	// 	fmt.Printf("%s: %d\n", category, count)
	// }
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
