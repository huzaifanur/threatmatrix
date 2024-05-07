package main

import (
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/unicode/norm"
)

const dummyTweet string = "asdw"

func main() {

	var tweet Tweet

	jsonData := []byte(dummyTweet)
	err := json.Unmarshal(jsonData, &tweet)
	if err != nil {
		log.Fatal(err)
	}

	// 1- get lang
	lang := tweet.Lang

	// replace newlines with spaces
	text := strings.ReplaceAll(tweet.Text, "\n", " ")
	// normalize
	text = normalizeText(text)

	encoding := "utf-8"
	encoded := false

	//
	if found := slices.Contains(IsoLanguages, lang); found {
		encoding = "ISO-8859-1"
		text = encodeDecodeISO8859_1(text)
		encoded = true
	}

}

/*
Normalization ensures that strings with different Unicode
representations but identical visual appearances are treated
as equal.

NFKD stands for "Normalization Form Compatibility Decomposition."
It's one of the Unicode normalization forms defined
by the Unicode Consortium.
*/
func normalizeText(text string) string {
	// Create a Normalizer with NFKD normalization
	n := norm.NFKD

	// Normalize the text
	normalizedText := n.String(text)

	return normalizedText
}

func encodeDecodeISO8859_1(str string) string {
	// Encode to ISO8859-1
	encoded := make([]byte, 0)
	for _, r := range str {
		if e, ok := charmap.ISO8859_1.EncodeRune(r); ok {
			encoded = append(encoded, e)
		}
		// Skip the rune if it cannot be encoded.
	}

	// Decode from ISO8859-1
	var decoded string
	decoder := charmap.ISO8859_1.NewDecoder()
	decodedBytes, _ := decoder.Bytes(encoded)
	// Ignoring error as we are skipping non-decodable bytes.
	decoded = string(decodedBytes)

	return decoded
}
